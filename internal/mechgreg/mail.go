package mechgreg

import (
	"log"

	"github.com/BESTRobotics/registry/internal/mail"
)

func (mg *MechanicalGreg) mailTeamCoaches(teamID int, template, subj string, ctx mail.LetterContext) error {
	t, err := mg.GetTeam(teamID)
	if err != nil {
		log.Println("Error mailing team coaches, team load:", err)
		return err
	}

	l, err := mail.RenderLetter(template, &ctx)
	if err != nil {
		log.Println("Error mailing team coaches, rendering:", err)
		return err
	}

	for i := range t.Coaches {
		l.AddTo(mail.UserToAddress(t.Coaches[i]))
	}

	l.Subject = subj

	if err := mg.po.SendMail(l); err != nil {
		log.Println("Error mailing team coaches, sending:", err)
		return err
	}
	return nil
}
