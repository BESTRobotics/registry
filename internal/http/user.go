package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/BESTRobotics/registry/internal/mechgreg"
	"github.com/BESTRobotics/registry/internal/models"
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
	switch err {
	case nil:
		break
	case mechgreg.ErrUserExists:
		c.AbortWithError(http.StatusConflict, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.Set("Location", fmt.Sprintf("/v1/users/%d", uid))
	c.Status(http.StatusCreated)
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
	switch err {
	case nil:
		break
	case mechgreg.ErrNoSuchUser:
		c.AbortWithError(http.StatusNotFound, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *Server) getUsers(c *gin.Context) {
	minUID := int64(0)
	count := int64(25)
	minUIDStr := c.Query("min")
	countStr := c.Query("count")
	var err error

	if minUIDStr != "" {
		minUID, err = strconv.ParseInt(minUIDStr, 10, 32)
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

	set, err := s.mg.GetUserPage(int(minUID), int(count))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, set)
}

func (s *Server) modUser(c *gin.Context) {
	// Deserialize the user
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	// Fetch the User from the request
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user.UID = int(uid)
	user.Username = ""

	switch s.mg.ModUser(user) {
	case nil:
		break
	case mechgreg.ErrNoSuchUser:
		c.AbortWithError(http.StatusNotFound, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (s *Server) delUser(c *gin.Context) {
	// Fetch the User from the request
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	switch s.mg.DelUser(int(uid)) {
	case nil:
		break
	case mechgreg.ErrNoSuchUser:
		c.AbortWithError(http.StatusNotFound, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.Status(http.StatusNoContent)
}
