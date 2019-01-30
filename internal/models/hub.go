package models

import (
	"time"
)

// The Hub is an organization that may run events on behalf of BEST.
type Hub struct {
	// ID is the primary key for the Hub.
	ID int `gorm:"AUTO_INCREMENT;primary_key"`

	// A hub must have a name that is unique.  While this name can
	// be changed, its not particularly recommended.
	Name string `gorm:"not null;unique"`

	// The hub director is the primary point of contact for the
	// hub and is responsible for its smooth running.  Changing
	// this field requires the approval of BRI.
	Director User

	// A director can't do it alone, they have people that help
	// them.  These are the admins and can act as the director for
	// the hub in many actions.  They can be added and removed by
	// the director.
	Admins []User

	// Location stores the City and State of a hub.  Since the hub
	// itself is usually only rooted to a city, this is stored as
	// a string.
	Location string

	// The Description is the blurb for the hub.  This is a good
	// place to add a history for the hub or talk about the goals
	// of the specific hub.
	Description string

	// Founded stores the founding date for a hub to tell how old
	// it is.  Though noone's keeping score, there's nonetheless
	// something exciting about a hub that's been around longer
	// than most of its competitors have been alive.
	Founded time.Time
}
