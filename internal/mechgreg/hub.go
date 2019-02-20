package mechgreg

import (
	"log"
	"net/http"
	"time"

	"github.com/asdine/storm"

	"github.com/BESTRobotics/registry/internal/models"
)

// NewHub creates a new hub from the provided hub model.  If the
// hub already exists an error will be returned instead.
func (mg *MechanicalGreg) NewHub(h models.Hub) (int, error) {
	h.Director = models.User{}
	h.Admins = nil

	err := mg.s.Save(&h)
	switch err {
	case nil:
		return h.ID, nil
	case storm.ErrAlreadyExists:
		return 0, NewConstraintError("A hub with these specifications already exists", err, http.StatusConflict)
	default:
		return 0, NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)

	}
}

// GetHub returns a hub based on the ID in the query.
func (mg *MechanicalGreg) GetHub(id int) (models.Hub, error) {
	var hub models.Hub

	err := mg.s.One("ID", id, &hub)
	switch err {
	case nil:
		break
	case storm.ErrNotFound:
		return models.Hub{}, NewConstraintError("No hub exists for that ID", err, http.StatusNotFound)
	default:
		return models.Hub{}, NewInternalError("An unspecified failure has occured while loading the hub", err, http.StatusInternalServerError)
	}

	director, err := mg.GetUser(hub.Director.ID)
	if err != nil {
		log.Println("Couldn't load director", err)
	}

	var admins []models.User
	for i := range hub.Admins {
		admin, err := mg.GetUser(hub.Admins[i].ID)
		if err != nil {
			log.Println("Error loading admin", err)
			continue
		}
		admins = append(admins, admin)
	}
	hub.Admins = admins
	hub.Director = director
	return hub, nil
}

// GetHubs returns pages of hubs, with a facility to find hubs that
// are inactive.
func (mg *MechanicalGreg) GetHubs(includeInactive bool) ([]models.Hub, error) {
	var tmp []models.Hub
	var out []models.Hub
	var err error

	if includeInactive {
		err = mg.s.All(&tmp)
	} else {
		err = mg.s.Find("InactiveSince", models.DateTime{}, &tmp)
	}
	switch err {
	case nil:
		break
	case storm.ErrNotFound:
		// In this specific case, notfound actually means
		// there are no teams satisfying the query.
		return []models.Hub{}, nil
	default:
		return nil, NewInternalError("An unspecified internal error has occured", err, http.StatusInternalServerError)
	}

	// This looks rather innefficient, but remember that the
	// backing boltdb is memory mapped, and the alternative would
	// be to duplicate code from the GetHub function.
	for i := range tmp {
		h, err := mg.GetHub(tmp[i].ID)
		if err != nil {
			log.Println("Error loading hub:", err)
			continue
		}
		out = append(out, h)
	}

	return out, nil
}

// GetHubsForUser returns all the hubs that a use has power over in
// some way, shape, or form.
func (mg *MechanicalGreg) GetHubsForUser(userID int) ([]models.Hub, error) {
	involvements := make(map[int]models.Hub)

	// Query for all hubs that have ever been, then figure out if
	// any of them have this person.
	hubs, err := mg.GetHubs(true)
	if err != nil {
		return nil, err
	}

	// Iterate through the hubs and find any that have this user
	// as a director or admin.  This isn't N^2 even though it
	// looks like it!
	for i := range hubs {
		if hubs[i].Director.ID == userID {
			involvements[hubs[i].ID] = hubs[i]
		}

		for j := range hubs[i].Admins {
			if hubs[i].Admins[j].ID == userID {
				involvements[hubs[i].ID] = hubs[i]
			}
		}
	}

	// Downconvert to just a list
	var out []models.Hub
	for _, hub := range involvements {
		out = append(out, hub)
	}
	return out, nil
}

// ModHub modifies an existing hub to match the state provided.
func (mg *MechanicalGreg) ModHub(h models.Hub) error {
	// These fields require special handline to safely update.
	h.Director = models.User{}
	h.Admins = nil
	h.InactiveSince = models.DateTime{}

	// Run the update
	return mg.modHub(h)
}

// modHub is just like ModHub but doesn't clear certain fields.
func (mg *MechanicalGreg) modHub(h models.Hub) error {
	err := mg.s.Update(&h)
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		return NewConstraintError("No hub exists for that ID", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}
}

// DeactivateHub sets the InactiveSince value on hubs that are no
// longer in operation.  This allows us to keep the hubs in the system
// rather than deleting them, which would both unnecessarily
// complicate the DB structure and would imply that once gone a hub
// won't ever come back.
func (mg *MechanicalGreg) DeactivateHub(id int) error {
	return mg.modHub(models.Hub{ID: id, InactiveSince: models.DateTime(time.Now())})
}

// ActivateHub brings a hub back from an inactive state.
func (mg *MechanicalGreg) ActivateHub(id int) error {
	// Needs to use UpdateField in order to explicitely zero the
	// value.
	err := mg.s.UpdateField(&models.Hub{ID: id}, "InactiveSince", models.DateTime{})
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		return NewConstraintError("No hub exists for that ID", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}
}

// SetHubDirector sets the director for the hub in question.
func (mg *MechanicalGreg) SetHubDirector(hubID int, director models.User) error {
	uid := director.ID

	user, err := mg.GetUser(uid)
	if err != nil {
		return err
	}

	return mg.modHub(models.Hub{ID: hubID, Director: models.User{ID: user.ID}})
}

// GetHubDirector returns the director for a given hub.
func (mg *MechanicalGreg) GetHubDirector(id int) (models.User, error) {
	h, err := mg.GetHub(id)
	if err != nil {
		return models.User{}, err
	}
	return h.Director, nil
}

// AddHubAdmin adds an admin to the specified hub.
func (mg *MechanicalGreg) AddHubAdmin(hubID int, admin models.User) error {
	hub, err := mg.GetHub(hubID)
	if err != nil {
		return err
	}

	hub.Admins = patchUserSlice(hub.Admins, true, admin)

	return mg.modHub(hub)
}

// DelHubAdmin removes an administrator from the hub.
func (mg *MechanicalGreg) DelHubAdmin(hubID int, admin models.User) error {
	hub, err := mg.GetHub(hubID)
	if err != nil {
		return err
	}

	admins := patchUserSlice(hub.Admins, false, admin)

	// Needs to use UpdateField in order to explicitely zero the
	// value
	err = mg.s.UpdateField(&models.Hub{ID: hubID}, "Admins", admins)
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		return NewConstraintError("No hub exists for that ID", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}
}
