package mechgreg

import (
	"path/filepath"

	"github.com/asdine/storm"
	"github.com/spf13/viper"
)

// New returns a new mechanical greg. The mechanical greg abstraction
// is a convenient way to refer to the behaviors and things the server
// needs to do.  Kind of like a mechanical turk, the mechanical greg
// performs tasks that might otherwise be done by a human, and does
// them with better repeatability and accuracy than a human could.
func New() (*MechanicalGreg, error) {
	dbPath := filepath.Join(viper.GetString("storage.root"), "registry.db")

	s, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}

	mg := MechanicalGreg{
		s: s,
	}
	return &mg, nil
}
