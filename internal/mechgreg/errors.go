package mechgreg

import "errors"

var (
	// ErrUserExists is to be returned in the case of an existing
	// user.  Specificially in the case that the primary key
	// collides.
	ErrUserExists = errors.New("The specified user already exists")

	// ErrNoSuchUser is to be returned in the case of a request
	// for a user that doesn't exist.
	ErrNoSuchUser = errors.New("No user exists with the specified ID")

	// ErrInternal is returned when something unforeseen goes
	// wrong.  The original error must be dumped to the log before
	// returning this error.
	ErrInternal = errors.New("An internal error has occured")
)
