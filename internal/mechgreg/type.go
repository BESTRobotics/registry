package mechgreg

import (
	"github.com/jinzhu/gorm"
)

// The ResourceBundle contains all the resources that we'd like
// Mechanical Greg to abstract for us.  This includes the database,
// any storage layers, and any ugly program logic that is arcane and
// should be hidden.
type ResourceBundle struct {
	DB *gorm.DB
}

// MechanicalGreg is a convenience type that binds all the methods of
// the abstraction together.
type MechanicalGreg struct {
	rb *ResourceBundle
}
