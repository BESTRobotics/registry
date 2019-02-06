package mechgreg

import (
	"log"

	"github.com/asdine/storm"

	"github.com/BESTRobotics/registry/internal/models"
)

// NewEvent creates a new event and returns its ID.
func (mg *MechanicalGreg) NewEvent(e models.Event) (int, error) {
	switch mg.s.Save(&e) {
	case nil:
		return e.ID, nil
	case storm.ErrAlreadyExists:
		return 0, ErrResourceExists
	default:
		return 0, ErrInternal
	}
}

// ModEvent updates an existing event.
func (mg *MechanicalGreg) ModEvent(e models.Event) error {
	switch mg.s.Update(&e) {
	case nil:
		return nil
	case storm.ErrNotFound:
		return ErrNoSuchResource
	default:
		return ErrInternal
	}
}

// GetEvent returnsa single event
func (mg *MechanicalGreg) GetEvent(id int) (models.Event, error) {
	var event models.Event
	switch mg.s.One("ID", id, &event) {
	case nil:
		break
	case storm.ErrNotFound:
		return models.Event{}, ErrNoSuchResource
	default:
		return models.Event{}, ErrInternal
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

// RegisterTeam takes a team and puts them in an event.  It does so
// implicitly and without confirmation.
func (mg *MechanicalGreg) RegisterTeam(eventID, rosterID int) (int, error) {
	if _, err := mg.GetEvent(eventID); err != nil {
		return err
	}

	if _, err := mg.GetRoster(rosterID); err != nil {
		return err
	}

	attendee := models.EventAttendee{
		Type: "TEAM",
		Event: models.Event{
			ID: eventID,
		},
		Team: models.TeamRoster{
			ID: rosterID,
		},
	}

	switch mg.s.Save(&attendee) {
	case nil:
		return attendee.ID, nil
	case storm.ErrAlreadyExists:
		return 0, ErrResourceExists
	default:
		return 0, ErrInternal
	}
}

// DeregisterTeam removes a team from an event.  This would be able to
// be shared with the user if we knew the attendee ID in advance, but
// we don't so this is a bit clunkier.
func (mg *MechanicalGreg) DeregisterTeam(eventID, teamID int) error {
	return nil
}
