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
}
