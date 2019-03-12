package null

import (
	"log"

	"github.com/BESTRobotics/registry/internal/mail"
)

type nullMailer struct{}

func init() {
	mail.Register("null", new)
}

func new() (mail.Mailer, error) {
	log.Println("Null mailer has initialized")
	return &nullMailer{}, nil
}

func (nm *nullMailer) SendMail(l mail.Letter) error {
	log.Println("From:", l.From)
	log.Println("To:", l.To)
	log.Println("Subject:", l.Subject)
	log.Println("Body:", l.Body)
	return nil
}
