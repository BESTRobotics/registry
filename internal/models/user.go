package models

import (
	"time"
	//"github.com/jinzhu/gorm"
)

// User models the fields that are directly part of a human user of
// the system.
type User struct {
	// Version metadata that needs to be assigned to the user to
	// allow more advanced data management.
	Meta Meta

	// The UID is a numeric ID for the user that is used to
	// uniquely identify them throughout the system.  This field
	// is required to be set at all times and is the primary key
	// for the user.
	UID int `gorm:"AUTO_INCREMENT;primary_key"`

	// The user can have a username, these must be unique within
	// the system, but can be changed at any time.
	Username string `gorm:"not null;unique"`

	// The user is required to have a valid address to receive
	// mail.
	EMail string

	// Type is the type of user that this represents.  This can be
	// things like "STUDENT" or "TEACHER" or "VOLUNTEER" etc.
	Type string

	// FirstName is the first part of a user's name, it is not
	// guaranteed to be a single word.
	FirstName string

	// LastName is the second part of a users's name, it is not
	// guaranteed to be a single word.
	LastName string

	// Birthdate contains the user's birthdate, this is used to
	// calculate whether or not the user can sign things.
	Birthdate time.Time
}
