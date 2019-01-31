package mechgreg

import (
	"log"

	"github.com/BESTRobotics/registry/internal/models"
)

// NewUser creates a new user from the provided user model.  If the
// user already exists an error will be returned instead.
func (mg *MechanicalGreg) NewUser(u models.User) (int, error) {
	var tu models.User
	su := models.User{Username: u.Username}
	if err := mg.rb.DB.Where(&su).First(&tu).Error; err == nil {
		// This worked, so the user already existed
		return 0, ErrResourceExists
	}

	if err := mg.rb.DB.Create(&u).Error; err != nil {
		log.Println(err)
		return 0, ErrInternal
	}
	return u.UID, nil
}

// GetUser returns a user based on the ID in the query.
func (mg *MechanicalGreg) GetUser(uid int) (models.User, error) {
	var user models.User

	if err := mg.rb.DB.First(&user, uid).Error; err != nil {
		log.Println(err)
		return models.User{}, ErrNoSuchResource
	}
	return user, nil
}

// ModUser modifies an existing user to match the state provided.
func (mg *MechanicalGreg) ModUser(u models.User) error {
	log.Println(u)
	var tu models.User
	if err := mg.rb.DB.First(&tu, u.UID).Error; err != nil {
		return ErrNoSuchResource
	}

	if err := mg.rb.DB.Model(&u).Updates(u).Error; err != nil {
		return ErrInternal
	}
	return nil
}

// DelUser removes a user from the server.
func (mg *MechanicalGreg) DelUser(uid int) error {
	if err := mg.rb.DB.Delete(&models.User{UID: uid}).Error; err != nil {
		log.Println(err)
		return ErrInternal
	}
	return nil
}

// GetUserPage returns a slice of users starting with a minID and
// returning count number of users.
func (mg *MechanicalGreg) GetUserPage(minID int, count int) ([]models.User, error) {
	out := []models.User{}

	err := mg.rb.DB.Where("UID >= ?", minID, "22").Limit(count).Find(&out).Error
	if err != nil {
		return []models.User{}, ErrInternal
	}

	return out, nil
}
