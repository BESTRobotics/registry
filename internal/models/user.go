package models

//go:generate stringer -type=Capability

// User models the fields that are directly part of a human user of
// the system.
type User struct {
	// The ID is a numeric ID for the user that is used to
	// uniquely identify them throughout the system.  This field
	// is required to be set at all times and is the primary key
	// for the user.
	ID int `storm:"increment"`

	// The user can have a username, these must be unique within
	// the system, but can be changed at any time.
	Username string `storm:"unique,index"`

	// The user is required to have a valid address to receive
	// mail.
	EMail string `storm:"unique,index"`

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
	Birthdate DateTime

	// Capabilities confer special powers within the system that a
	// user might have in addition to those granted by owning a
	// particular resource.
	Capabilities []Capability
}

// HasCapability can be used to check if a user has a particular
// capability.
func (u User) HasCapability(c Capability) bool {
	for i := range u.Capabilities {
		if u.Capabilities[i] == c || u.Capabilities[i] == CapSuperAdmin {
			return true
		}
	}
	return false
}

// GrantCapability can be used to assign a capability to a user.
func (u *User) GrantCapability(c Capability) {
	for i := range u.Capabilities {
		if u.Capabilities[i] == c {
			return
		}
	}
	u.Capabilities = append(u.Capabilities, c)
	return
}

// RemoveCapability removes a given capability from a user.
func (u *User) RemoveCapability(c Capability) {
	var out []Capability
	for i := range u.Capabilities {
		if u.Capabilities[i] == c {
			continue
		}
		out = append(out, u.Capabilities[i])
	}
	u.Capabilities = out
}

// Capability represents certain additional special powers within the
// system that a user can have.  While it would be nicer to assign
// these only ever to groups, this is clean enough since adding groups
// and a more complex directory structure to the already complicated
// registry is impractical.
type Capability int

const (
	// CapSuperAdmin implies the ability to do effectively any change
	// since this short-circuits through the logic that checks for
	// other capabilities.
	CapSuperAdmin Capability = iota

	// CapUserAdmin can modify all fields on a user, except
	// modifying capabilities.
	CapUserAdmin

	// CapUserRead tells the system that you're allowed to read
	// all information on users, not just that which is otherwise
	// publicly visible.
	CapUserRead

	// CapIDMax isn't so much a capability as an upper bound to
	// iterate to.  It doesn't confer powers, but provides an
	// upper point to the range to validate to.
	CapIDMax
)
