package mg

import (
	"log"
	"time"
	"context"
	"fmt"

	"github.com/mailgun/mailgun-go/v3"
	"github.com/spf13/viper"

	"github.com/BESTRobotics/registry/internal/mail"
)

type bestmailgun struct{
	mailgun.Mailgun
}

func init() {
	mail.Register("mailgun", new)
}

func new() (mail.Mailer, error) {
	log.Println("Initializing mailgun mailer")

	m := mailgun.NewMailgun(viper.GetString("mailgun.domain"), viper.GetString("mailgun.privatekey"))
	return &bestmailgun{m}, nil
}

func (mg *bestmailgun) SendMail(l mail.Letter) error {
	ml := mg.NewMessage(fmt.Sprintf("%s",l.From), l.Subject, l.Body, fmt.Sprintf("%s", l.To))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	r, id, err := mg.Send(ctx, ml)
	if err != nil {
		return err
	}

	log.Printf("Sent message %s, got %s", id, r)
	return nil
}
