package models

// A BRCTeam is an instantiation of a team that competes as part of
// the BEST Robotics Competition program.  This structore stores the
// team components that change year to year.
type BRCTeam struct {
	// The ID is useful for syncing to other systems, but not much
	// else.
	ID int `storm:"increment"`

	// The BRCTeam knows which team it represents, as well as what
	// season it is part of.
	Team   Team
	Season Season

	// For speed reasons, the TeamID and SeasonID are also stored
	// as indexed fields.
	TeamID   int `storm:"index"`
	SeasonID int `storm:"index"`

	// State is a string here to accomodate many as yet
	// unforeseen statuses that a team may take on during thier
	// registration to a particular season and hub.
	State string

	// The JoinKey is a key that lets you join the team.  Its a
	// "secret" value, but is really only meant to allow the
	// joining of a team, so its not that secure.
	JoinKey string `storm:"increment"`

	// TeamSurvey stores the information that needs to be
	// collected with team granularity.  This is statistical
	// information that is related to the school and the team that
	// might change.  For information that is related to
	// individuals or information that isn't going to change (or
	// can be computed) consult the other embedded data.
	TeamSurvey struct {
		SchoolSize string
		SchoolType string
	}

	// Symbol is a BEST Robotics Competition specific tag that can
	// take the place of a team name/number and is a field that is
	// unique within a particular year.  These may change each
	// year, but the team that had a symbol previously has the
	// right to get it again.
	Symbol string `storm:"index"`

	// The BRCTeam contains a roster of students that are a part
	// of the team.
	Roster []User

	// While this could be stored within the roster, its easier to
	// store it here as a seperate ordered list.
	Drivers []User
}
