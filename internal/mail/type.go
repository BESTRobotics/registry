package mail

import (
	"net/mail"
)

// A Mailer implements the components needed to send email to a
// various persons.
type Mailer interface {
	SendMail(Letter) error
}

// A Letter here means the same thing it does in the real world.  Its
// a message with a from, a to, a subject and a body.
type Letter struct {
	From    mail.Address
	To      mail.Address
	Subject string
	Body    string
}

// A Factory returns a configured mailer that is ready to use.
type Factory func() (Mailer, error)
