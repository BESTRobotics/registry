package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/BESTRobotics/registry/internal/models"
	"github.com/BESTRobotics/registry/internal/token"
)

func (s *Server) registerLocalUser(c *gin.Context) {
	// Decode the user struct and add the authenticating
	// information to an authdata which needs to be seperated from
	// the user meta.  After that send a mail with the magic link
	// to activate the account.
	var regRequest struct {
		U        models.User
		Password string
	}
	if err := c.ShouldBindJSON(&regRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	// If this returns nil then the user already exists and we
	// abort.
	if _, err := s.mg.UsernameExists(regRequest.U.Username); err == nil {
		c.Status(http.StatusConflict)
		return
	}

	// User doesn't exist, time to create the user:
	_, err := s.mg.NewUser(regRequest.U)
	if err != nil {
		s.handleError(c, err)
		return
	}

	// And now we set the authdata (password in this case).
	if err := s.mg.SetUserPassword(regRequest.U.Username, regRequest.Password); err != nil {
		s.handleError(c, err)
		return
	}

	// TODO Message the user via Email that they need to use the
	// activation link.
	c.Status(http.StatusCreated)
}

// loginLocaluser logs in the user using the authentication
// information provided and redirects the browser into the app with a
// token unless a boomerang target was specified in the query string.
func (s *Server) loginLocalUser(c *gin.Context) {
	target := "/app/login#t=%s"

	var ld struct {
		Username string
		Password string
	}
	if err := c.ShouldBindJSON(&ld); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	// Figure out who this user is.
	user, err := s.mg.UsernameExists(ld.Username)
	if err != nil {
		s.handleError(c, err)
		return
	}

	// Check the user password, if its successful then generate a
	// token and hand it off.
	err = s.mg.CheckUserPassword(ld.Username, ld.Password)
	if err != nil {
		s.handleError(c, err)
		return
	}

	// Get the token
	log.Println(user)
	tkn, err := s.generateToken(user.ID, token.GetConfig())
	if err != nil {
		s.handleError(c, err)
		return
	}

	// Put the token into the URL
	target = fmt.Sprintf(target, tkn)
	c.Redirect(http.StatusMovedPermanently, target)
}
