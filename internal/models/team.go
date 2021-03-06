package models

// A Team is an entity that is sponsored by a school and participates
// in one or more BEST programs.  A team has people that are on it,
// but this is stored seperately since it changes from year to year.
type Team struct {
	// ID is the numeric ID for a team.  This is distinct from the
	// number that is assigned when a team participates, and is
	// instead the number that is used for internal tracking.
	ID int `storm:"increment"`

	// SchoolName is the name of the host school.
	SchoolName string

	// SchoolAddress is the address at which the school can get
	// physical mail.
	SchoolAddress string

	// Website is the website for the school or team, it doesn't
	// really matter which.
	Website string

	// While some teams will come up with a name ever year for how
	// they wish to be represented, the entity that is a team
	// needs to have a name we can call them.
	StaticName string

	// Every team has a "Home" Hub where they normally are
	// present, and that hub director is the point of contact for
	// the team.  The ID is stored seperately to make it indexed.
	HomeHubID int `storm:"index"`
	HomeHub   Hub

	// Teams in general have only a single coach, but in some
	// schools its desirable for the principal or other
	// administrator to be a coach for continuity reasons.
	Coaches []User

	// Teams can also be inactive, which means they won't appear
	// in lookups or otherwise be available.  This is different
	// from being without a season to participate in.
	InactiveSince DateTime

	// Its handy to know how long teams have been around.
	Founded DateTime

	// BRIApproved determines whether or not this team has been
	// approved and accepted.
	BRIApproved bool
}
