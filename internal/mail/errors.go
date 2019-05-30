package mail

import (
	"errors"
)

var (
	// ErrUnknownMailer is returned if a mailer is requested but
	// isn't known to the system.
	ErrUnknownMailer = errors.New("Mailer implementation is not known")

	// ErrNoSuchLetter is returned if a templated letter is
	// requested but isn't known to the system.
	ErrNoSuchLetter = errors.New("No letter with that slug is known")

	// ErrInternal is a catch all for things that are unforseen.
	ErrInternal = errors.New("An unspecified error has occured")
)
