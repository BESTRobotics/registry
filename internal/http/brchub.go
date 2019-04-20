package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/BESTRobotics/registry/internal/models"
)

func (s *Server) registerBRCHub(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	seasonStr := c.Param("season")
	season, err := strconv.ParseInt(seasonStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Check that the season exists
	if _, err := s.mg.GetSeason(int(season)); err != nil {
		return s.handleError(c, err)
	}

	// Perform Authorization Checks
	if err := canModHub(extractClaims(c), int(id)); err != nil {
		return s.handleError(c, err)
	}

	if _, err := s.mg.RegisterBRCHub(int(id), int(season)); err != nil {
		return s.handleError(c, err)
	}

	brchub, err := s.mg.GetBRCHub(int(id), int(season))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusCreated, brchub)
}

func (s *Server) getBRCHub(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	seasonStr := c.Param("season")
	season, err := strconv.ParseInt(seasonStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	brchub, err := s.mg.GetBRCHub(int(id), int(season))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, brchub)
}

func (s *Server) getBRCHubs(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	set, err := s.mg.GetBRCHubs(int(id))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, set)
}

func (s *Server) updateBRCHub(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	seasonStr := c.Param("season")
	season, err := strconv.ParseInt(seasonStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Perform Authorization Checks
	if err := canModHub(extractClaims(c), int(id)); err != nil {
		return s.handleError(c, err)
	}

	var update models.BRCHub
	if err := c.Bind(&update); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := s.mg.UpdateBRCHub(int(id), int(season), update); err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}
