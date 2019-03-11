package mechgreg

import (
	"log"

	"github.com/BESTRobotics/registry/internal/models"
)

// New returns a new mechanical greg. The mechanical greg abstraction
// is a convenient way to refer to the behaviors and things the server
// needs to do.  Kind of like a mechanical turk, the mechanical greg
// performs tasks that might otherwise be done by a human, and does
// them with better repeatability and accuracy than a human could.
func New(rb ResourceBundle) (*MechanicalGreg, error) {
	mg := MechanicalGreg{
		s: rb.StormDB,
	}

	if err := mg.s.ReIndex(&models.User{}); err != nil {
		log.Println("Error during indexing", err)
	}
	if err := mg.s.ReIndex(&models.AuthData{}); err != nil {
		log.Println("Error during indexing", err)
	}
	if err := mg.s.ReIndex(&models.Hub{}); err != nil {
		log.Println("Error during indexing", err)
	}
	if err := mg.s.ReIndex(&models.Team{}); err != nil {
		log.Println("Error during indexing", err)
	}
	if err := mg.s.ReIndex(&models.Season{}); err != nil {
		log.Println("Error during indexing", err)
	}
	if err := mg.s.ReIndex(&models.School{}); err != nil {
		log.Println("Error during indexing", err)
	}

	return &mg, nil
}
