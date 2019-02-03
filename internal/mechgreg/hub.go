package mechgreg

import (
	"log"
	"time"

	"github.com/asdine/storm"

	"github.com/BESTRobotics/registry/internal/models"
)

// NewHub creates a new hub from the provided hub model.  If the
// hub already exists an error will be returned instead.
func (mg *MechanicalGreg) NewHub(h models.Hub) (int, error) {
	h.Director = models.User{}
	h.Admins = nil

	switch mg.s.Save(&h) {
	case nil:
		return h.ID, nil
	case storm.ErrAlreadyExists:
		return 0, ErrResourceExists
	default:
		return 0, ErrInternal

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
		return models.Hub{}, ErrNoSuchResource
	default:
		return models.Hub{}, ErrInternal
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
		err = mg.s.Find("InactiveSince", time.Time{}, &tmp)
	}
	if err != nil {
		return nil, err
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

// ModHub modifies an existing hub to match the state provided.
func (mg *MechanicalGreg) ModHub(h models.Hub) error {
	// These fields require special handline to safely update.
	h.Director = models.User{}
	h.Admins = nil
	h.InactiveSince = time.Time{}

	// Run the update
	return mg.modHub(h)
}

// modHub is just like ModHub but doesn't clear certain fields.
func (mg *MechanicalGreg) modHub(h models.Hub) error {
	switch mg.s.Update(&h) {
	case nil:
		return nil
	case storm.ErrNotFound:
		return ErrNoSuchResource
	default:
		return ErrInternal
	}
}

// DeactivateHub sets the InactiveSince value on hubs that are no
// longer in operation.  This allows us to keep the hubs in the system
// rather than deleting them, which would both unnecessarily
// complicate the DB structure and would imply that once gone a hub
// won't ever come back.
func (mg *MechanicalGreg) DeactivateHub(id int) error {
	return mg.modHub(models.Hub{ID: id, InactiveSince: time.Now()})
}

// ActivateHub brings a hub back from an inactive state.
func (mg *MechanicalGreg) ActivateHub(id int) error {
	// Needs to use UpdateField in order to explicitely zero the
	// value.
	switch (mg.s.UpdateField(&models.Hub{ID: id}, "InactiveSince", time.Time{})) {
	case nil:
		return nil
	case storm.ErrNotFound:
		return ErrNoSuchResource
	default:
		return ErrInternal
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
	// value.
	switch (mg.s.UpdateField(&models.Hub{ID: hubID}, "Admins", admins)) {
	case nil:
		return nil
	case storm.ErrNotFound:
		return ErrNoSuchResource
	default:
		return ErrInternal
	}
}