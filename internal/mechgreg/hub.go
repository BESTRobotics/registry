package mechgreg

import (
	"github.com/BESTRobotics/registry/internal/models"
)

// NewHub creates a new hub from the provided hub model.  If the
// hub already exists an error will be returned instead.
func (mg *MechanicalGreg) NewHub(h models.Hub) (int, error) {
	return 0, nil
}

// GetHub returns a hub based on the ID in the query.
func (mg *MechanicalGreg) GetHub(id int) (models.Hub, error) {
	return models.Hub{}, nil
}

// GetHubs returns pages of hubs, with a facility to find hubs that
// are inactive.
func (mg *MechanicalGreg) GetHubs(includeInactive bool) ([]models.Hub, error) {
	return nil, nil
}

// ModHub modifies an existing hub to match the state provided.
func (mg *MechanicalGreg) ModHub(h models.Hub) error {
	return nil
}

// DeactivateHub sets the InactiveSince value on hubs that are no
// longer in operation.  This allows us to keep the hubs in the system
// rather than deleting them, which would both unnecessarily
// complicate the DB structure and would imply that once gone a hub
// won't ever come back.
func (mg *MechanicalGreg) DeactivateHub(id int) error {
	return nil
}

// ActivateHub brings a hub back from an inactive state.
func (mg *MechanicalGreg) ActivateHub(id int) error {
	return nil
}
