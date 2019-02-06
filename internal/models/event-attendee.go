package models

// The EventAttendee joins the event to its attendees and forms the
// many-to-many relationship.
type EventAttendee struct {
	// ID is the numeric ID for this attendance.
	ID int `storm:"increment"`

	// Type is the type of attendee, represented as a string.
	// This is intended to allow the filtering of attendees to a
	// type that can then be processed in bulk.
	Type string

	// Event is the Event that this attendee is going to.
	Event Event

	// User is a user is an individual going to an event.
	User User

	// TeamRoster is a team that is going to an event as an entire
	// group.
	Team TeamRoster
}
