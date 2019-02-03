package models

// School represents the information that needs to be known about a
// school that is sponsoring a team.
type School struct {
	// ID is the numerical ID for this school with BEST.  Schools
	// may already have a numeric ID of some kind, but this is the
	// number assigned by BEST.
	ID int `storm:"increment"`

	// Name of the school as it appears on the school's official
	// paperwork.
	Name string

	// Address of the school, as it appears on the school's
	// official paperwork.
	Address string

	// Website for the school.
	Website string
}
