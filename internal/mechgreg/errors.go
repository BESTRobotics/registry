package mechgreg

import "errors"

var (
	// ErrResourceExists is to be returned in the case of an existing
	// resource.
	ErrResourceExists = errors.New("The specified resource already exists")

	// ErrNoSuchResource is to be returned in the case of a
	// request for a resource that doesn't exist.
	ErrNoSuchResource = errors.New("No resource exists with the specified selector")

	// ErrInternal is returned when something unforeseen goes
	// wrong.  The original error must be dumped to the log before
	// returning this error.
	ErrInternal = errors.New("An internal error has occured")
)
