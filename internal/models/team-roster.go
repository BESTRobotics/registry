package models

// The TeamRoster exists year to year and points to a season for a
// given team.  This allows the team to have a flexible membership
// that happens year on year rather than needing to patch the team
// record itself.
type TeamRoster struct {
	// ID is the numeric ID for this particular roster.
	ID int `storm:"increment"`

	// The JoinKey is a semi-secret token that's needed to join
	// the team.  It allows a user to lookup the team and join it
	// in one action.  This has to be set by the coach or a mentor
	// of the team and must be unique.
	JoinKey string `storm:"unique"`

	// Team is the team this roster belongs to.
	Team Team

	// Season is the season this roster is attached to.
	Season Season

	// Members is the membership of the team on this roster.
	Members []User
}
