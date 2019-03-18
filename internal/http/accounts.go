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

	// If this returns nil then the user already exists and we
	// abort.
	if _, err := s.mg.UsernameExists(regRequest.U.Username); err == nil {
		return c.NoContent(http.StatusConflict)
	}

	// User doesn't exist, time to create the user:
	_, err := s.mg.NewUser(regRequest.U)
	if err != nil {
		return s.handleError(c, err)
	}

	// And now we set the authdata (password in this case).
	if err := s.mg.SetUserPassword(regRequest.U.Username, regRequest.Password); err != nil {
		return s.handleError(c, err)
	}

	l := mail.NewLetter()
	l.AddTo(mail.UserToAddress(regRequest.U))
	l.Subject = "Activate Your BRI Registry Account"
	l.Body = "Some magic link in here to activate this account..."

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
		Username string
		Password string
	}
	if err := c.Bind(&ld); err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Figure out who this user is.
	user, err := s.mg.UsernameExists(ld.Username)
	if err != nil {
		return s.handleError(c, err)
	}

	// Check the user password, if its successful then generate a
	// token and hand it off.
	err = s.mg.CheckUserPassword(ld.Username, ld.Password)
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
