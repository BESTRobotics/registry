package models

import "time"

// An Event is the core of the scheduler system.  It encompases both
// the idea of how many entities may attend an event, as well as its
// location and sponsors.
type Event struct {
	// The ID is the numeric key for the event.
	ID int `storm:"increment"`

	// Name is the human readable name of the event.
	Name string

	// Description is the human readable description of the event.
	Description string

	// Location is a well formatted location string that should be
	// a fully qualified address.
	Location string

	// StartTime is when the event commences.
	StartTime time.Time

	// EndTime likewise is when the event ends.
	EndTime time.Time

	// Hub points to the hub that is sponsoring this event.  All
	// events are required to be sponsored by hubs.
	Hub Hub
}
