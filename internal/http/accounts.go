package http

import (
	"log"
	"time"
	"net/http"
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/BESTRobotics/registry/internal/mail"
	"github.com/BESTRobotics/registry/internal/models"
	"github.com/BESTRobotics/registry/internal/token"
)

func init() {
	viper.SetDefault("account.activation.timeout", time.Hour * 24)
}

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
	u, err := s.mg.GetUser(userID)
	if err != nil {
		return s.handleError(c, err)
	}

	// And now we set the authdata (password in this case).
	if err := s.mg.SetUserPassword(userID, regRequest.Password); err != nil {
		return s.handleError(c, err)
	}

	actURL := s.genActivationString(u)

	// Send the letter
	l, err := mail.RenderLetter("new-local-user", &mail.LetterContext{LocalMessage: actURL})
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

// genActivationURL generates a URL which can be used to activate a
// user's account.
func (s *Server) genActivationString(u models.User) string {
	expTime := time.Now().Add(viper.GetDuration("account.activation.timeout")).Unix()
	raw := fmt.Sprintf("%d.%d", u.ID, expTime)
	log.Println(raw)
	h := hmac.New(sha256.New, []byte(viper.GetString("account.hmac_secret")))
	h.Write([]byte(raw))
	sha := hex.EncodeToString(h.Sum(nil))
	out := fmt.Sprintf("%s.%s", raw, sha)
	log.Println(out)
	return out
}

func (s Server) checkActivationString() {

}

// activateUser flips the activation bit and allows a user to sign in.
// This handler expects to get a URL with a signed hmac which will
// specify the user to activate, and the time interval over which to
// permit the activation.
func (s *Server) activateUser(c echo.Context) error {

	return nil
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

	if user.Active == nil || !*user.Active {
		return c.NoContent(http.StatusPreconditionFailed)
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
