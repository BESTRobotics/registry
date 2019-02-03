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

func (s *Server) newHub(c *gin.Context) {
	var hub models.Hub
	if err := c.ShouldBindJSON(&hub); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	id, err := s.mg.NewHub(hub)
	switch err {
	case nil:
		break
	case mechgreg.ErrResourceExists:
		c.AbortWithError(http.StatusConflict, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.Set("Location", fmt.Sprintf("/v1/hubs/%d", id))
	c.Status(http.StatusCreated)
}

func (s *Server) getHub(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	hub, err := s.mg.GetHub(int(id))
	switch err {
	case nil:
		break
	case mechgreg.ErrNoSuchResource:
		c.AbortWithError(http.StatusNotFound, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, hub)
}

func (s *Server) getHubs(c *gin.Context) {
	allStr := c.Query("include-inactive")
	all := false
	if allStr != "" {
		all = true
	}

	set, err := s.mg.GetHubs(all)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, set)
}

func (s *Server) modHub(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var hub models.Hub
	if err := c.ShouldBindJSON(&hub); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	hub.ID = int(id)

	switch s.mg.ModHub(hub) {
	case nil:
		break
	case mechgreg.ErrNoSuchResource:
		c.AbortWithError(http.StatusNotFound, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (s *Server) deactivateHub(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	switch s.mg.DeactivateHub(int(id)) {
	case nil:
		break
	case mechgreg.ErrNoSuchResource:
		c.AbortWithError(http.StatusNotFound, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (s *Server) activateHub(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	switch s.mg.ActivateHub(int(id)) {
	case nil:
		break
	case mechgreg.ErrNoSuchResource:
		c.AbortWithError(http.StatusNotFound, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (s *Server) setHubDirector(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	switch s.mg.SetHubDirector(int(id), user) {
	case nil:
		break
	case mechgreg.ErrNoSuchResource:
		c.AbortWithError(http.StatusNotFound, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (s *Server) getHubDirector(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	hub, err := s.mg.GetHubDirector(int(id))
	switch err {
	case nil:
		break
	case mechgreg.ErrNoSuchResource:
		c.AbortWithError(http.StatusNotFound, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, hub)
}

func (s *Server) addHubAdmin(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	switch s.mg.AddHubAdmin(int(id), user) {
	case nil:
		break
	case mechgreg.ErrNoSuchResource:
		c.AbortWithError(http.StatusNotFound, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (s *Server) delHubAdmin(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	switch s.mg.DelHubAdmin(int(id), user) {
	case nil:
		break
	case mechgreg.ErrNoSuchResource:
		c.AbortWithError(http.StatusNotFound, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.Status(http.StatusNoContent)
}
