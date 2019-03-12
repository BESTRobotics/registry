package mail

import (
	"errors"
)

var (
	// ErrUnknownMailer is returned if a mailer is requested but
	// isn't known to the system.
	ErrUnknownMailer = errors.New("Mailer implementation is not known")
)
