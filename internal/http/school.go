package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/BESTRobotics/registry/internal/models"
)

func (s *Server) newSchool(c echo.Context) error {
	// Perform Authorization Checks
	if err := canManageTeams(extractClaims(c)); err != nil {
		return s.handleError(c, err)
	}

	var school models.School
	if err := c.Bind(&school); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	id, err := s.mg.NewSchool(school)
	if err != nil {
		return s.handleError(c, err)
	}
	school, err = s.mg.GetSchool(id)
	if err != nil {
		return s.handleError(c, err)
	}
	return c.JSON(http.StatusCreated, school)
}

func (s *Server) getSchool(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	school, err := s.mg.GetSchool(int(id))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, school)
}

func (s *Server) getSchools(c echo.Context) error {
	schools, err := s.mg.GetSchools()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, schools)
}

func (s *Server) modSchool(c echo.Context) error {
	// Perform Authorization Checks
	if err := canManageTeams(extractClaims(c)); err != nil {
		return s.handleError(c, err)
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var school models.School
	if err := c.Bind(&school); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	school.ID = int(id)

	err = s.mg.ModSchool(school)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}
