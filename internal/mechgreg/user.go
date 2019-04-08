package mechgreg

import (
	"net/http"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"

	"github.com/BESTRobotics/registry/internal/models"
)

func init() {
	viper.SetDefault("auth.hashcost", 10)
}

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

// FillUserProfile is used to load and embed the profile in places
// where a fully populated user is needed.
func (mg *MechanicalGreg) FillUserProfile(u *models.User) error {
	p, err := mg.GetUserProfile(u.ID)
	if err != nil {
		return err
	}
	u.UserProfile = &p
	return nil
}

// GetUserProfile fetches just the profile information for a particular user.
func (mg *MechanicalGreg) GetUserProfile(uid int) (models.UserProfile, error) {
	var p models.UserProfile
	if err := mg.s.One("UserID", uid, &p); err != nil {
		return models.UserProfile{},
			NewInternalError("Profile could not be retrieved", err, http.StatusInternalServerError)
	}
	return p, nil
}

// SetUserProfile is used to update the profile information for a
// particular user.
func (mg *MechanicalGreg) SetUserProfile(uid int, p models.UserProfile) error {
	p.ID = 0
	p.UserID = uid
	err := mg.s.Update(&p)
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		return NewConstraintError("No such profile exists with that ID", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}
}

// UsernameExists can be used to check if a username exists before creating a new user.
func (mg *MechanicalGreg) UsernameExists(username string) (models.User, error) {
	var user models.User

	err := mg.s.One("Username", username, &user)
	switch err {
	case nil:
		return user, nil
	case storm.ErrNotFound:
		return models.User{}, NewConstraintError("No such user exists", err, http.StatusNotFound)
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

// SetUserPassword is used to, unsurprisingly, set the user password.
// This is meant to be used during user setup, and uses bcrypt as the
// hashing engine.
func (mg *MechanicalGreg) SetUserPassword(username, password string) error {
	// Compute the hash from the incomming password
	cost := viper.GetInt("auth.hashcost")
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return NewInternalError("BCrypt Failure", err, http.StatusInternalServerError)
	}

	// Save the data
	d := models.AuthData{
		Username: username,
		Provider: "PASSWORD",
		Password: string(hash[:]),
	}
	err = mg.s.Save(&d)
	if err != nil {
		return NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}
	return nil
}

// CheckUserPassword checks a password for a specific user.
func (mg *MechanicalGreg) CheckUserPassword(username, password string) error {
	query := mg.s.Select(q.Eq("Username", username), q.Eq("Provider", "PASSWORD"))

	var ad models.AuthData
	// Safe to use First here because there had better only ever
	// be one of these.
	err := query.First(&ad)
	switch err {
	case nil:
		break
	case storm.ErrNotFound:
		return NewConstraintError("No such user exists with that ID", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(ad.Password), []byte(password)); err != nil {
		return NewAuthError("The authentication information is incorrect", err, http.StatusUnauthorized)
	}
	return nil
}
