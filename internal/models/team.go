package models

import "time"

// A Team is an entity that is sponsored by a school and participates
// in one or more BEST programs.  A team has people that are on it,
// but this is stored seperately since it changes from year to year.
type Team struct {
	// ID is the numeric ID for a team.  This is distinct from the
	// number that is assigned when a team participates, and is
	// instead the number that is used for internal tracking.
	ID int `storm:"increment"`

	// While some teams will come up with a name ever year for how
	// they wish to be represented, the entity that is a team
	// needs to have a name we can call them.
	StaticName string

	// Every team has a "Home" Hub where they normally are
	// present, and that hub director is the point of contact for
	// the team.
	HomeHub Hub

	// A team has exactly one coach.  This is the person who the
	// school has designated as responsible for the team's
	// actions, and is the point of contact for delivering
	// paperwork for the team and other key functions.
	Coach User

	// Like the hub director, its rare the coach does it on their
	// own.  They have Mentors that help them out to run the team
	// and keep things moving along.  Mentors have similar powers
	// in the system to the Coach.
	Mentors []User

	// As stated above, the team must be associated with a school,
	// so that is stored here.
	School School

	// Teams can also be inactive, which means they won't appear
	// in lookups or otherwise be available.  This is different
	// from being without a season to participate in.
	InactiveSince time.Time

	// Its handy to know how long teams have been around.
	Founded time.Time
}
