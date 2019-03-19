package models

import (
	"net/url"
)

// A BRCHub is a hub that is hosting a season of the BEST Robotics
// Competition.
type BRCHub struct {
	// ID as an integer for convenience.
	ID int `storm:"index,increment"`

	// This hub is owned by a normal hub.
	Hub Hub

	// BRC hubs are seasonal.
	Season Season

	// To make this even remotely fast, we also keep the HubID and
	// SeasonID as unstructured ints.
	HubID    int `storm:"index"`
	SeasonID int `storm:"index"`

	// Slice of events as a hub may have more than the minimum, so
	// we give them a slice for convenience.
	Events []Event

	// Assorted metadata about the BRCHub.  This is where all
	// kinds of fun information lives that tracks the hub's
	// participation in a season.
	Meta struct {
		BRIApproved bool
		Sponsors    []struct {
			Name string
			Logo url.URL
		}
	}
}
