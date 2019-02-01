package mechgreg

import (
	"log"
	"time"

	"github.com/BESTRobotics/registry/internal/models"
)

// NewHub creates a new hub from the provided hub model.  If the
// hub already exists an error will be returned instead.
func (mg *MechanicalGreg) NewHub(h models.Hub) (int, error) {
	var th models.Hub
	sh := models.Hub{Name: h.Name}
	if err := mg.rb.DB.Where(&sh).First(&th).Error; err == nil {
		// This worked, so the hub already existed
		return 0, ErrResourceExists
	}

	if err := mg.rb.DB.Create(&h).Error; err != nil {
		log.Println(err)
		return 0, ErrInternal
	}
	return h.ID, nil
}

// GetHub returns a hub based on the ID in the query.
func (mg *MechanicalGreg) GetHub(id int) (models.Hub, error) {
	var hub models.Hub

	if err := mg.rb.DB.First(&hub, id).Error; err != nil {
		log.Println(err)
		return models.Hub{}, ErrNoSuchResource
	}
	return hub, nil
}

// GetHubs returns pages of hubs, with a facility to find hubs that
// are inactive.
func (mg *MechanicalGreg) GetHubs(includeInactive bool) ([]models.Hub, error) {
	var out []models.Hub
	var err error

	if includeInactive {
		err = mg.rb.DB.Find(&out).Error
	} else {
		err = mg.rb.DB.Not(&models.Hub{InactiveSince: time.Time{}}).Find(&out).Error
	}

	if err != nil {
		log.Println(err)
		return nil, ErrInternal
	}
	return out, nil	
}

// ModHub modifies an existing hub to match the state provided.
func (mg *MechanicalGreg) ModHub(h models.Hub) error {
	var th models.Hub
	if err := mg.rb.DB.First(&th, h.ID).Error; err != nil {
		return ErrNoSuchResource
	}

	if err := mg.rb.DB.Model(&h).Updates(h).Error; err != nil {
		return ErrInternal
	}
	return nil
}

// DeactivateHub sets the InactiveSince value on hubs that are no
// longer in operation.  This allows us to keep the hubs in the system
// rather than deleting them, which would both unnecessarily
// complicate the DB structure and would imply that once gone a hub
// won't ever come back.
func (mg *MechanicalGreg) DeactivateHub(id int) error {
	var th models.Hub
	if err := mg.rb.DB.First(&th, id).Error; err != nil {
		return ErrNoSuchResource
	}

	if err := mg.rb.DB.Model(&models.Hub{}).Updates(models.Hub{ID: id, InactiveSince: time.Now()}).Error; err != nil {
		log.Println(err)
		return ErrInternal
	}
	return nil
}

// ActivateHub brings a hub back from an inactive state.
func (mg *MechanicalGreg) ActivateHub(id int) error {
	var th models.Hub
	if err := mg.rb.DB.First(&th, id).Error; err != nil {
		return ErrNoSuchResource
	}

	if err := mg.rb.DB.Model(&models.Hub{}).Updates(models.Hub{ID: id, InactiveSince: time.Time{}}).Error; err != nil {
		log.Println(err)
		return ErrInternal
	}
	return nil
}
