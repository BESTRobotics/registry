package mechgreg

import (
	"net/http"

	"github.com/asdine/storm"

	"github.com/BESTRobotics/registry/internal/models"
)

// NewSchool creates a new school resource on the server.  Schools are
// not deletable, since we'd like to keep a record of all schools that
// have ever participated.
func (mg *MechanicalGreg) NewSchool(s models.School) (int, error) {
	err := mg.s.Save(&s)
	switch err {
	case nil:
		return s.ID, nil
	case storm.ErrAlreadyExists:
		return 0, NewConstraintError("A school with that name already exists", err, http.StatusConflict)
	default:
		return 0, NewInternalError("An unspecified internal error has occured", err, http.StatusInternalServerError)
	}
}

// GetSchool returns a single school via ID.
func (mg *MechanicalGreg) GetSchool(id int) (models.School, error) {
	var school models.School

	err := mg.s.One("ID", id, &school)
	switch err {
	case nil:
		return school, nil
	case storm.ErrNotFound:
		return models.School{}, NewConstraintError("No school with that ID exists", err, http.StatusNotFound)
	default:
		return models.School{}, NewInternalError("An unspecified internal error has occured", err, http.StatusInternalServerError)
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
		return NewConstraintError("No school with that ID exists", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified internal error has occured", err, http.StatusInternalServerError)
	}
}
