package models

import "time"

// Meta models certain information that all other types are required to have.
type Meta struct {
	// Its very useful to know when an object was created.
	Created time.Time

	// Version allows us to resolve conflicts between objects that
	// have been removed and modified externally.
	Version int
}
