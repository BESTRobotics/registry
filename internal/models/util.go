package models

import (
	"bytes"
	"time"
)

// DateTime wraps the time.Time but returns null if its the zero time.
// Special thanks to GitHub user Nilium for this reflection free time
// interface and helping to debug the data storage problems.
type DateTime time.Time

// MarshalJSON implements the JSON marshaller interface.
func (t DateTime) MarshalJSON() ([]byte, error) {
	if t.Time().Equal(time.Time{}) {
		return []byte("null"), nil
	}
	return t.Time().MarshalJSON()
}

// UnmarshalJSON implements the JSON marshaller interface.
func (t *DateTime) UnmarshalJSON(data []byte) error {
	if bytes.Compare(data, []byte("null")) == 0 {
		t = &DateTime{}
		return nil
	}
	// Unmarshal into tm using the time.Time UnmarshalJSON (same effect as
	// json.Unmarshal(data, &tm) here, just no reflection involved).
	var tm time.Time
	if err := tm.UnmarshalJSON(data); err != nil {
		return err
	}
	*t = DateTime(tm)
	return nil
}

// Time returns the underlying time type.
func (t DateTime) Time() time.Time {
	return time.Time(t)
}
