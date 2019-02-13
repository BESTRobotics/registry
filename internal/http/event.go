package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/BESTRobotics/registry/internal/models"
)

func (s *Server) newEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	id, err := s.mg.NewEvent(event)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.Set("Location", fmt.Sprintf("/v1/events/%d", id))
	c.Status(http.StatusCreated)
}

func (s *Server) getEvent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	event, err := s.mg.GetEvent(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, event)
}

func (s *Server) getEvents(c *gin.Context) {
	set, err := s.mg.GetEvents()
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, set)
}

func (s *Server) modEvent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	event.ID = int(id)

	err = s.mg.ModEvent(event)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
