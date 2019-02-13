package mechgreg

import (
	"log"
	"net/http"

	"github.com/asdine/storm"

	"github.com/BESTRobotics/registry/internal/models"
)

// NewEvent creates a new event and returns its ID.
func (mg *MechanicalGreg) NewEvent(e models.Event) (int, error) {
	err := mg.s.Save(&e)
	switch err {
	case nil:
		return e.ID, nil
	case storm.ErrAlreadyExists:
		return 0, NewConstraintError("An event with these specifications already exists!", err, http.StatusConflict)
	default:
		return 0, NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}
}

// ModEvent updates an existing event.
func (mg *MechanicalGreg) ModEvent(e models.Event) error {
	err := mg.s.Update(&e)
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		return NewConstraintError("No event exists with that ID", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}
}

// GetEvent returnsa single event
func (mg *MechanicalGreg) GetEvent(id int) (models.Event, error) {
	var event models.Event
	err := mg.s.One("ID", id, &event)
	switch err {
	case nil:
		break
	case storm.ErrNotFound:
		return models.Event{}, NewConstraintError("No event exists with that ID", err, http.StatusNotFound)
	default:
		return models.Event{}, NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}

	hub, err := mg.GetHub(event.Hub.ID)
	if err != nil {
		return models.Event{}, err
	}

	event.Hub = hub
	return event, nil
}

// GetEvents returns all events.  This is of somewhat questionable
// utility since there's no way to archive events or otherwise
// suppress them from showing up here.
func (mg *MechanicalGreg) GetEvents() ([]models.Event, error) {
	var out []models.Event
	var tmp []models.Event

	if err := mg.s.All(&tmp); err != nil {
		log.Printf("Error loading events: %s", err)
		return nil, err
	}

	for i := range tmp {
		e, err := mg.GetEvent(tmp[i].ID)
		if err != nil {
			log.Println("Error loading event:", err)
			continue
		}
		out = append(out, e)
	}

	return out, nil
}
