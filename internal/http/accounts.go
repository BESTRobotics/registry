package http

import (
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/BESTRobotics/registry/internal/mail"
	"github.com/BESTRobotics/registry/internal/models"
	"github.com/BESTRobotics/registry/internal/token"
)

func (s *Server) registerLocalUser(c echo.Context) error {
	// Decode the user struct and add the authenticating
	// information to an authdata which needs to be seperated from
	// the user meta.  After that send a mail with the magic link
	// to activate the account.
	var regRequest struct {
		U        models.User
		Password string
	}
	if err := c.Bind(&regRequest); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// User doesn't exist, time to create the user:
	userID, err := s.mg.NewUser(regRequest.U)
	if err != nil {
		return s.handleError(c, err)
	}

	// And now we set the authdata (password in this case).
	if err := s.mg.SetUserPassword(userID, regRequest.Password); err != nil {
		return s.handleError(c, err)
	}

	// Send the letter
	l, err := mail.RenderLetter("new-local-user", &mail.LetterContext{LocalMessage: "Foo"})
	if err != nil {
		log.Println("Error trying to mail new user:", err)
		return c.NoContent(http.StatusInternalServerError)
	}
	l.AddTo(mail.UserToAddress(regRequest.U))
	l.Subject = "Activate Your BRI Registry Account"
	if err := s.po.SendMail(l); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// TODO Message the user via Email that they need to use the
	// activation link.
	return c.NoContent(http.StatusCreated)
}

// loginLocaluser logs in the user using the authentication
// information provided and redirects the browser into the app with a
// token unless a boomerang target was specified in the query string.
func (s *Server) loginLocalUser(c echo.Context) error {
	// target := "/app/login#t=%s"

	var ld struct {
		EMail    string
		Password string
	}
	if err := c.Bind(&ld); err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Get user by email.  This should be reasonably fast, but is
	// easily the slowest part of this since the user emails
	// aren't indexed.
	user, err := s.mg.GetUserByEMail(ld.EMail)
	if err != nil {
		return s.handleError(c, err)
	}

	// Check the user password, if its successful then generate a
	// token and hand it off.
	err = s.mg.CheckUserPassword(user.ID, ld.Password)
	if err != nil {
		return s.handleError(c, err)
	}

	// Get the token
	log.Println(user)
	tkn, err := s.generateToken(user.ID, token.GetConfig())
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, tkn)
}
