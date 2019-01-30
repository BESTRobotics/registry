package models

// A Team is an entity that is sponsored by a school and participates
// in one or more BEST programs.  A team has people that are on it,
// but this is stored seperately since it changes from year to year.
type Team struct {
	// A team has exactly one coach.  This is the person who the
	// school has designated as responsible for the team's
	// actions, and is the point of contact for delivering
	// paperwork for the team and other key functions.
	Coach User

	// Like the hub director, its rare the coach does it on their
	// own.  They have Mentors that help them out to run the team
	// and keep things moving along.  Mentors have similar powers
	// in the system to the Coach.
	Mentor []User

	// As stated above, the team must be associated with a school,
	// so that is stored here as a string.
	SchoolName string
}
