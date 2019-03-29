package models

// The Season type parents all events in a series.  This includes the
// concept of a program as well as that of a competition season.
type Season struct {
	// ID is the primary key for the Season.
	ID int `storm:"increment"`

	// Season names are the human readable name.  This should have
	// a friendly name like "BEST Robotics Competition 2019"
	Name string `storm:"unique"`

	// Archived determines if this season is archived or not.
	// Archived seasons lock all resources within them.
	Archived bool

	// State allows the season to take on program specific states
	// such as OPEN or REGISTRATION_CLOSED.
	State string

	// Program allows us to determine what program this season
	// represents.  Not all seasons necessarily have programs, but
	// programs usually do have seasons.
	Program Program
}

// Program allows us to know what program this season is referencing.
type Program int

const (
	// ProgramBRC is the BEST Robotics Competition
	ProgramBRC Program = iota
)
