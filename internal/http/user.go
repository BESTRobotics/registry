package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/BESTRobotics/registry/internal/models"
	"github.com/BESTRobotics/registry/internal/token"
)

func (s *Server) newUser(c *gin.Context) {
	// Deserialize the user
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	// Attempt to create the user
	uid, err := s.mg.NewUser(user)
	if err != nil {
		s.handleError(c, err)
		return
	}
	user, err = s.mg.GetUser(uid)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (s *Server) getUser(c *gin.Context) {
	// Fetch the User from the request
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := s.mg.GetUser(int(uid))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *Server) getUsers(c *gin.Context) {
	page := int64(0)
	count := int64(25)
	pageStr := c.Query("page")
	countStr := c.Query("count")
	var err error

	if pageStr != "" {
		page, err = strconv.ParseInt(pageStr, 10, 32)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	if countStr != "" {
		count, err = strconv.ParseInt(countStr, 10, 32)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	}

	set, err := s.mg.GetUserPage(int(page), int(count))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, set)
}

func (s *Server) modUser(c *gin.Context) {
	// Fetch the User from the request
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Perform Authentication Checks
	if err := canModUser(extractClaims(c), int(uid)); err != nil {
		s.handleError(c, err)
		return
	}

	// Deserialize the user
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	user.ID = int(uid)
	user.Username = ""

	err = s.mg.ModUser(user)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// canModUser checks if a user modification is allowed.  This will
// return true in the case of the included user matching the requested
// ID, or the USER_ADMIN capability being present.
func canModUser(claims token.Claims, id int) error {
	if claims.IsEmpty() {
		return newAuthError("Unauthorized", "Claims are empty")
	}

	if claims.User.ID != id {
		return newAuthError("Unauthorized", "You are not authorized to modify this user")
	}
	return nil
}
