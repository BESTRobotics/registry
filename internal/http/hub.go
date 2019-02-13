package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

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
	if err != nil {
		s.handleError(c, err)
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
	if err != nil {
		s.handleError(c, err)
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
		s.handleError(c, err)
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

	err = s.mg.ModHub(hub)
	if err != nil {
		s.handleError(c, err)
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

	err = s.mg.DeactivateHub(int(id))
	if err != nil {
		s.handleError(c, err)
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

	err = s.mg.ActivateHub(int(id))
	if err != nil {
		s.handleError(c, err)
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

	err = s.mg.SetHubDirector(int(id), user)
	if err != nil {
		s.handleError(c, err)
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
	if err != nil {
		s.handleError(c, err)
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

	err = s.mg.AddHubAdmin(int(id), user)
	if err != nil {
		s.handleError(c, err)
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

	err = s.mg.DelHubAdmin(int(id), user)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
