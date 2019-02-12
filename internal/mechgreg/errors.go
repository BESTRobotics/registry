package mechgreg

import (
	"errors"
)

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

// ConstraintError is used to handle errors that result from invalid
// or unsatisfiable constraints.
type ConstraintError struct {
	Message string
	Cause   error

	httpCode int
}

// NewConstraintError returns an initialized constriant error.
func NewConstraintError(s string, err error, c int) error {
	return &ConstraintError{
		Message:  s,
		Cause:    err,
		httpCode: c,
	}
}

func (c *ConstraintError) Error() string {
	return c.Message
}

// Code returns the http status code that this error should represent.
func (c *ConstraintError) Code() int {
	return c.httpCode
}

// InternalError is used to type errors that are beyond the scope of
// other errors and signify an immediate abort in processing.
type InternalError struct {
	Message string
	Cause   error

	httpCode int
}

// NewInternalError returns an internal error structure initialized and populated.
func NewInternalError(s string, err error, c int) error {
	return &InternalError{
		Message:  s,
		Cause:    err,
		httpCode: c,
	}
}

func (e *InternalError) Error() string {
	return e.Message
}

// Code returns the http status code that this error should represent.
func (e *InternalError) Code() int {
	return e.httpCode
}
