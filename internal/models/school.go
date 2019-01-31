package models

// School represents the information that needs to be known about a
// school that is sponsoring a team.
type School struct {
	// Name of the school as it appears on the school's official
	// paperwork.
	Name string

	// Address of the school, as it appears on the school's
	// official paperwork.
	Address string

	// Website for the school.
	Website string
}
