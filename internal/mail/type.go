package mail

import (
	"log"
	"net/mail"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("mail.DefaultFrom", "BEST Robotics <no-reply@bestinc.org>")
}

// A Mailer implements the components needed to send email to a
// various persons.
type Mailer interface {
	SendMail(Letter) error
}

// A Letter here means the same thing it does in the real world.  Its
// a message with a from, a to, a subject and a body.
type Letter struct {
	From    *mail.Address
	To      []*mail.Address
	BCC     []*mail.Address
	Subject string
	Body    string
}

// NewLetter creates a new letter with the from address filled in
// automatically.
func NewLetter() Letter {
	a, err := mail.ParseAddress(viper.GetString("mail.DefaultFrom"))
	if err != nil {
		log.Println("DefaultFrom parse failure", err)
		a = &mail.Address{Address: viper.GetString("mail.DefaultFrom")}
	}
	return Letter{From: a}
}

// AddTo adds a single address to the To line of the letter.  It
// understands how to do so without creating duplicate recipients.
func (l *Letter) AddTo(a *mail.Address) {
	for i := range l.To {
		if l.To[i] == a {
			return
		}
	}
	l.To = append(l.To, a)
}

// AddBCC adds a single address to the BCC line of the letter.  It
// understands how to do so without creating duplicate recipients.
func (l *Letter) AddBCC(a *mail.Address) {
	for i := range l.BCC {
		if l.BCC[i] == a {
			return
		}
	}
	l.BCC = append(l.BCC, a)
}

// A Factory returns a configured mailer that is ready to use.
type Factory func() (Mailer, error)
