package mechgreg

import (
	"github.com/asdine/storm"

	"github.com/BESTRobotics/registry/internal/models"
)

// NewUser creates a new user from the provided user model.  If the
// user already exists an error will be returned instead.
func (mg *MechanicalGreg) NewUser(u models.User) (int, error) {
	switch mg.s.Save(&u) {
	case nil:
		return u.ID, nil
	case storm.ErrAlreadyExists:
		return 0, ErrResourceExists
	default:
		return 0, ErrInternal
	}
}

// GetUser returns a user based on the ID in the query.
func (mg *MechanicalGreg) GetUser(uid int) (models.User, error) {
	var user models.User

	switch mg.s.One("ID", uid, &user) {
	case nil:
		return user, nil
	case storm.ErrNotFound:
		return models.User{}, ErrNoSuchResource
	default:
		return models.User{}, ErrInternal
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
	switch mg.s.Update(&u) {
	case nil:
		return nil
	case storm.ErrNotFound:
		return ErrNoSuchResource
	default:
		return ErrInternal
	}
}
