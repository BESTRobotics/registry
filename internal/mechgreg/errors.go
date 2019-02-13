package mechgreg

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
