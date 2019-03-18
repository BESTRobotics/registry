package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/BESTRobotics/registry/internal/models"
)

func (s *Server) newEvent(c echo.Context) error {
	var event models.Event
	if err := c.Bind(&event); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	id, err := s.mg.NewEvent(event)
	if err != nil {
		return s.handleError(c, err)
	}
	event, err = s.mg.GetEvent(id)
	if err != nil {
		return s.handleError(c, err)
	}
	return c.JSON(http.StatusCreated, event)
}

func (s *Server) getEvent(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	event, err := s.mg.GetEvent(int(id))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, event)
}

func (s *Server) getEvents(c echo.Context) error {
	set, err := s.mg.GetEvents()
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, set)
}

func (s *Server) modEvent(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var event models.Event
	if err := c.Bind(&event); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	event.ID = int(id)

	err = s.mg.ModEvent(event)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}
