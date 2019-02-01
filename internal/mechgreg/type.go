package mechgreg

import (
	"github.com/asdine/storm"
)

// MechanicalGreg is a convenience type that binds all the methods of
// the abstraction together.
type MechanicalGreg struct {
	s *storm.DB
}
