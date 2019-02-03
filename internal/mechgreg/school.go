package mechgreg

import (
	"github.com/asdine/storm"

	"github.com/BESTRobotics/registry/internal/models"
)

// NewSchool creates a new school resource on the server.  Schools are
// not deletable, since we'd like to keep a record of all schools that
// have ever participated.
func (mg *MechanicalGreg) NewSchool(s models.School) (int, error) {
	switch mg.s.Save(&s) {
	case nil:
		return s.ID, nil
	case storm.ErrAlreadyExists:
		return 0, ErrResourceExists
	default:
		return 0, ErrInternal
	}
}

// GetSchool returns a single school via ID.
func (mg *MechanicalGreg) GetSchool(id int) (models.School, error) {
	var school models.School

	switch mg.s.One("ID", id, &school) {
	case nil:
		return school, nil
	case storm.ErrNotFound:
		return models.School{}, ErrNoSuchResource
	default:
		return models.School{}, ErrInternal
	}
}

// GetSchools returns all known schools.
func (mg *MechanicalGreg) GetSchools() ([]models.School, error) {
	var out []models.School
	err := mg.s.All(&out)
	return out, err
}

// ModSchool can be used to update a school that already exists.
func (mg *MechanicalGreg) ModSchool(s models.School) error {
	switch mg.s.Update(&s) {
	case nil:
		return nil
	case storm.ErrNotFound:
		return ErrNoSuchResource
	default:
		return ErrInternal
	}
}
