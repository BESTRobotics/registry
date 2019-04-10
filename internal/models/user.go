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

	// Capabilities confer special powers within the system that a
	// user might have in addition to those granted by owning a
	// particular resource.
	Capabilities []Capability

	// For convenience, the User has a profile embedded, but not
	// loaded at all times.
	*UserProfile
}

// The UserProfile contains information that is not public on a user,
// but that may need to be made available to various systems
// internally.
type UserProfile struct {
	// Not really useful for anything, but for consistency we give
	// the UserProfile an int.
	ID int `storm:"increment"`

	// The profile knows which user it belongs to
	UserID int `storm:"unique,index"`

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
}

// A Student is a special type that handles an "owned" profile on a
// primary account.
type Student struct {
	// Allows us to later keep track of what student this is.
	ID int `storm:"increment"`

	// The ID of the user account that owns this student record is
	// kept seperate.
	UserID int

	// Students have names that are likely distinct from that of
	// the account holder.
	FirstName string
	LastName  string

	// A student might have their own email account which should
	// get stuff sent to it, but this is by default blank.
	Email string

	// Asking and storing this here makes it significantly more
	// efficient than asking people later, and increases the
	// likelyhood that people will actually fill out the surveys
	// that happen during the season.  String typed to support
	// future expansion to an arbitrarily large set of values.
	Race   string
	Gender string
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

	// CapHubAdmin permits the holder to create new hubs, manage
	// the archival of existing ones, and set the director on
	// hubs.
	CapHubAdmin

	// CapTeamAdmin permits the holder to create new teams, manage
	// artifacts for existing ones, and set the coach on teams.
	CapTeamAdmin

	// CapIDMax isn't so much a capability as an upper bound to
	// iterate to.  It doesn't confer powers, but provides an
	// upper point to the range to validate to.
	CapIDMax
)
