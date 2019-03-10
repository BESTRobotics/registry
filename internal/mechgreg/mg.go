package mechgreg

// New returns a new mechanical greg. The mechanical greg abstraction
// is a convenient way to refer to the behaviors and things the server
// needs to do.  Kind of like a mechanical turk, the mechanical greg
// performs tasks that might otherwise be done by a human, and does
// them with better repeatability and accuracy than a human could.
func New(rb ResourceBundle) (*MechanicalGreg, error) {
	mg := MechanicalGreg{
		s: rb.StormDB,
	}
	return &mg, nil
}
