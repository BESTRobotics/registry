package mechgreg

import (
	"net/http"

	"github.com/asdine/storm"

	"github.com/BESTRobotics/registry/internal/models"
)

// NewUser creates a new user from the provided user model.  If the
// user already exists an error will be returned instead.
func (mg *MechanicalGreg) NewUser(u models.User) (int, error) {
	err := mg.s.Save(&u)
	switch err {
	case nil:
		return u.ID, nil
	case storm.ErrAlreadyExists:
		return 0, NewConstraintError("Username and email must be unique!", err, http.StatusConflict)
	default:
		return 0, NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}
}

// GetUser returns a user based on the ID in the query.
func (mg *MechanicalGreg) GetUser(uid int) (models.User, error) {
	var user models.User

	err := mg.s.One("ID", uid, &user)
	switch err {
	case nil:
		return user, nil
	case storm.ErrNotFound:
		return models.User{}, NewConstraintError("No such user exists with that ID", err, http.StatusNotFound)
	default:
		return models.User{}, NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}
}

// GetUserPage returns a slice of users starting with a minID and
// returning count number of users.
func (mg *MechanicalGreg) GetUserPage(page int, count int) ([]models.User, error) {
	var out []models.User

	err := mg.s.All(&out, storm.Limit(count), storm.Skip(page*count))
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModUser modifies an existing user to match the state provided.
func (mg *MechanicalGreg) ModUser(u models.User) error {
	err := mg.s.Update(&u)
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		return NewConstraintError("No such user exists with that ID", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}
}
