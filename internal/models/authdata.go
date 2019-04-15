package models

import (
	"fmt"
	"strings"
)

// AuthData is what's used to actually log in users to the system.
type AuthData struct {
	// Though not directly used, the record still needs an index.
	ID int `storm:"id,increment"`

	// Username points to the user that this belongs to.  Its just
	// a pointer here, so just a string.
	UserID int `storm:"index"`

	// Provider here specifies what fields need to be filled in
	// for this authenticator.
	Provider string

	// Password is the hash of a users's password.
	Password string

	// ExternalUserID maps to some external system's idea of who
	// you are.
	ExternalUserID string
}

// Validate figures out if the data in the structure is valid for the
// provider that is declared.  It returns an error if it is not.
func (a AuthData) Validate() error {
	if a.UserID == 0 {
		return fmt.Errorf("userID must be set")
	}

	switch strings.ToUpper(a.Provider) {
	case "PASSWORD":
		if a.Password == "" {
			return fmt.Errorf("the password hash must be set")
		}
	case "OAUTH2":
		if a.ExternalUserID == "" {
			return fmt.Errorf("the external user id must be set")
		}
	}
	return nil
}
