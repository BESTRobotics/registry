package mechgreg

import (
	"github.com/asdine/storm"

	"github.com/BESTRobotics/registry/internal/mail"
)

// MechanicalGreg is a convenience type that binds all the methods of
// the abstraction together.
type MechanicalGreg struct {
	s  *storm.DB
	po mail.Mailer
}

// ResourceBundle is a type for bundling up resources that should be
// handed off for management by a mechgreg instance.
type ResourceBundle struct {
	StormDB *storm.DB
	Mailer  mail.Mailer
}
