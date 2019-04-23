package mechgreg

import (
	"log"
	"net/http"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"

	"github.com/BESTRobotics/registry/internal/models"
)

// RegisterBRCHub checks that the requested hub and season exist, then
// adds a BRCHub to the registry's known BRCHubs.
func (mg *MechanicalGreg) RegisterBRCHub(hubID, seasonID int) (int, error) {
	// Make sure the hub and season exists, then make sure no
	// BRCHub exists for this combination, then save the BRCHub.

	if _, err := mg.GetHub(hubID); err != nil {
		return -1, err
	}

	if _, err := mg.GetSeason(seasonID); err != nil {
		return -1, err
	}

	var brchub models.BRCHub
	query := mg.s.Select(q.And(q.Eq("HubID", hubID), q.Eq("SeasonID", seasonID)))
	if err := query.First(&brchub); err != storm.ErrNotFound {
		return -1, NewConstraintError("BRCHub already exists", err, http.StatusPreconditionFailed)
	}

	newBRCHub := &models.BRCHub{
		HubID:    hubID,
		SeasonID: seasonID,
	}

	err := mg.s.Save(newBRCHub)
	switch err {
	case nil:
		return newBRCHub.ID, nil
	case storm.ErrAlreadyExists:
		return -1, NewConstraintError("A BRCHub with that ID already exists?", err, http.StatusConflict)
	default:
		log.Println(err)
		return -1, NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}
}

// GetBRCHub loads the hub and fills in the underlying hub and season.
func (mg *MechanicalGreg) GetBRCHub(hubID, seasonID int) (models.BRCHub, error) {
	// Fetch the BRCHub
	var brchub models.BRCHub
	query := mg.s.Select(q.And(q.Eq("HubID", hubID), q.Eq("SeasonID", seasonID)))
	err := query.First(&brchub)
	switch err {
	case nil:
		break
	case storm.ErrNotFound:
		return models.BRCHub{}, NewConstraintError("This hub did not participate in that season", err, http.StatusNotFound)
	default:
		return models.BRCHub{}, NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}

	if err := mg.populateBRCHub(&brchub); err != nil {
		return models.BRCHub{}, err
	}

	return brchub, nil
}

// GetBRCHubs returns all BRCHubs for a particular hub.  This data is
// not paginated.
func (mg *MechanicalGreg) GetBRCHubs(hubID int) ([]models.BRCHub, error) {
	var out []models.BRCHub

	err := mg.s.Find("HubID", hubID, &out)
	switch err {
	case nil:
		break
	case storm.ErrNotFound:
		return []models.BRCHub{}, nil
	default:
		return nil, NewInternalError("An unspecified internal error has occured", err, http.StatusInternalServerError)
	}

	// Fill in the nested structures.
	for i := range out {
		mg.populateBRCHub(&out[i])
	}

	return out, nil
}

// GetBRCHubsForSeason returns all hubs for a particular season.  This
// data is not paginated.
func (mg *MechanicalGreg) GetBRCHubsForSeason(seasonID int) ([]models.BRCHub, error) {
	var out []models.BRCHub

	err := mg.s.Find("SeasonID", seasonID, &out)
	switch err {
	case nil:
		break
	case storm.ErrNotFound:
		return []models.BRCHub{}, nil
	default:
		return nil, NewInternalError("An unspecified internal error has occured", err, http.StatusInternalServerError)
	}

	// Fill in the nested structures.
	for i := range out {
		mg.populateBRCHub(&out[i])
	}

	return out, nil
}

func (mg *MechanicalGreg) populateBRCHub(h *models.BRCHub) error {
	// Get the underlying hub
	hub, err := mg.GetHub(h.HubID)
	if err != nil {
		return err
	}

	// Get the underlying season
	season, err := mg.GetSeason(h.SeasonID)
	if err != nil {
		return err
	}

	// Insert the underlying components to the BRCHub.
	h.Hub = hub
	h.Season = season
	return nil
}

// This is just like the public one, but can set all fields.
func (mg *MechanicalGreg) updateBRCHub(hubID, seasonID int, update models.BRCHub) error {
	update.HubID = hubID
	update.SeasonID = seasonID

	err := mg.s.Update(&update)
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		return NewConstraintError("This hub did not participate in that season", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}
}

// UpdateBRCHub can update all fields except for events and approval
// status.
func (mg *MechanicalGreg) UpdateBRCHub(hubID, seasonID int, update models.BRCHub) error {
	// Fetch the BRCHub
	var brchub models.BRCHub
	query := mg.s.Select(q.And(q.Eq("HubID", hubID), q.Eq("SeasonID", seasonID)))
	err := query.First(&brchub)
	switch err {
	case nil:
		break
	case storm.ErrNotFound:
		return NewConstraintError("This hub did not participate in that season", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}

	// Update() needs to see an ID in the struct
	update.ID = brchub.ID
	return mg.updateBRCHub(hubID, seasonID, update)
}
