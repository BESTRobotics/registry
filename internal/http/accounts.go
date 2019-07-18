package http

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/BESTRobotics/registry/internal/mail"
	"github.com/BESTRobotics/registry/internal/models"
	"github.com/BESTRobotics/registry/internal/token"
)

func init() {
	viper.SetDefault("account.activation.timeout", time.Hour*24)
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

	// If in the debug mode then activate now.
	if viper.GetBool("dev.activate") {
		u.Active = func(b bool) *bool { return &b }(true)
		if err := s.mg.ModUser(u); err != nil {
			return s.handleError(c, err)
		}
	}

	actURL := s.genActivationString(u)

	msg := "Please visit the following URL to activate your account: " + viper.GetString("core.url") + "v1/account/activate/" + actURL

	// Send the letter
	l, err := mail.RenderLetter("new-local-user", &mail.LetterContext{LocalMessage: msg})
	if err != nil {
		log.Println("Error trying to mail new user:", err)
		return c.NoContent(http.StatusInternalServerError)
	}
	l.AddTo(mail.UserToAddress(regRequest.U))
	l.Subject = "Activate Your BRI Registry Account"
	if err := s.po.SendMail(l); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusCreated)
}

func makeHMAC(s string) string {
	h := hmac.New(sha256.New, []byte(viper.GetString("account.hmac_secret")))
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// genActivationString generates a token which can be used to activate
// a user's account.
func (s *Server) genActivationString(u models.User) string {
	expTime := time.Now().Add(viper.GetDuration("account.activation.timeout")).Unix()
	log.Println(time.Now().Unix(), expTime)
	raw := fmt.Sprintf("%d.%d", u.ID, expTime)
	out := fmt.Sprintf("%s.%s", raw, makeHMAC(raw))
	return out
}

// checkActivationString validates the hmac, and that the token isn't
// expired.
func (s Server) checkActivationString(st string) error {
	parts := strings.Split(st, ".")
	if len(parts) != 3 {
		return errors.New("Malformed activation token")
	}

	t, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return err
	}

	sha := makeHMAC(strings.Join(parts[:2], "."))

	if !hmac.Equal([]byte(sha), []byte(parts[2])) {
		return errors.New("HMAC is invalid")
	}

	if t < time.Now().Unix() {
		return errors.New("Token expired")
	}
	return nil
}

// activateUser flips the activation bit and allows a user to sign in.
// This handler expects to get a URL with a signed hmac which will
// specify the user to activate, and the time interval over which to
// permit the activation.
func (s *Server) activateUser(c echo.Context) error {
	t := c.Param("token")

	log.Println("Reading activation token")
	if err := s.checkActivationString(t); err != nil {
		return s.handleError(c, err)
	}

	// This is safe to do here unchecked because we validated the
	// token structure above.
	uidStr := strings.Split(t, ".")[0]
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())

	}

	log.Println("Attempting to activate")
	user := models.User{
		ID:     int(uid),
		Active: func(b bool) *bool { return &b }(true),
	}
	if err := s.mg.ModUser(user); err != nil {
		return s.handleError(c, err)
	}

	return c.String(http.StatusOK, "Your account has been successfully activated")
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

// requestLocalPasswordReset can be used to request a local password
// reset email.  This will allow a local user to get a password reset
// link which they can then click to get a password emailed to them.
func (s *Server) requestLocalPasswordReset(c echo.Context) error {
	user, err := s.mg.GetUserByEMail(c.Param("email"))
	if err != nil {
		// Always return 204 to prevent user emails from being
		// enumerated.
		return c.NoContent(http.StatusOK)
	}

	tkn := s.genActivationString(user)

	msg := "Please use the following link to reset your password\n" + viper.GetString("core.url") + "v1/account/local/rpass/" + tkn

	l, err := mail.RenderLetter("reset-local-password", &mail.LetterContext{LocalMessage: msg})
	if err != nil {
		log.Println("Error trying to mail reset link:", err)
		return c.NoContent(http.StatusOK)
	}
	l.AddTo(mail.UserToAddress(user))
	l.Subject = "BRI Registry Password Reset"
	if err := s.po.SendMail(l); err != nil {
		log.Println("Error while sending mail:", err)
	}
	return c.NoContent(http.StatusOK)
}

func (s *Server) resetLocalPassword(c echo.Context) error {
	t := c.Param("token")

	if err := s.checkActivationString(t); err != nil {
		return s.handleError(c, err)
	}

	// This is safe to do here unchecked because we validated the
	// token structure above.
	uidStr := strings.Split(t, ".")[0]
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())

	}

	tmpPass := randStringRunes(16)

	if err := s.mg.SetUserPassword(int(uid), tmpPass); err != nil {
		return s.handleError(c, err)
	}

	msg := "Your password has been set to:\n" + tmpPass

	l, err := mail.RenderLetter("new-local-password", &mail.LetterContext{LocalMessage: msg})
	if err != nil {
		log.Println("Error trying to mail reset link:", err)
		return c.NoContent(http.StatusOK)
	}
	user, err := s.mg.GetUser(int(uid))
	if err != nil {
		return s.handleError(c, err)
	}
	l.AddTo(mail.UserToAddress(user))
	l.Subject = "BRI Registry Password Reset"
	if err := s.po.SendMail(l); err != nil {
		log.Println("Error while sending mail:", err)
		return c.String(http.StatusInternalServerError, "An error has occured, please try again later")
	}

	return c.String(http.StatusOK, "Consult your email for furthur instructions")
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// From StackOverflow
func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
