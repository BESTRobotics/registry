package mail

import (
	"fmt"
	"net/mail"

	"github.com/BESTRobotics/registry/internal/models"
)

// UserToAddress composes the address for a user and returns filled in
// struct.
func UserToAddress(u models.User) *mail.Address {
	if u.UserProfile == nil {
		// Set the address without being properly formatted.
		a := &mail.Address{Address: u.EMail}
		return a
	}
	// We can only send nicely formatted mail if the user
	// profile is present.
	aStr := fmt.Sprintf("%s %s <%s>", u.FirstName, u.LastName, u.EMail)
	a, err := mail.ParseAddress(aStr)
	if err != nil {
		// Set the address without being properly formatted.
		a = &mail.Address{Address: u.EMail}
	}
	return a
}
