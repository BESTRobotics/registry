package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/BESTRobotics/registry/internal/models"
)

func (s *Server) getBRCTeams(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	set, err := s.mg.GetBRCTeams(int(id))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, set)
}

func (s *Server) registerBRCTeam(c echo.Context) error {
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

	// Validate preconditions
	if _, err := s.mg.GetTeam(int(id)); err != nil {
		return s.handleError(c, err)
	}
	if _, err := s.mg.GetSeason(int(id)); err != nil {
		return s.handleError(c, err)
	}

	// Perform Authorization Checks
	if err := canModTeam(extractClaims(c), int(id)); err != nil {
		return s.handleError(c, err)
	}

	// Perform registration
	if _, err := s.mg.RegisterBRCTeam(int(id), int(season)); err != nil {
		return s.handleError(c, err)
	}

	// Retrieve and return
	brcteam, err := s.mg.GetBRCTeam(int(id), int(season))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusCreated, brcteam)
}

func (s *Server) getBRCTeam(c echo.Context) error {
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

	brcteam, err := s.mg.GetBRCTeam(int(id), int(season))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, brcteam)
}

func (s *Server) updateBRCTeam(c echo.Context) error {
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
	if err := canModTeam(extractClaims(c), int(id)); err != nil {
		return s.handleError(c, err)
	}

	var t models.BRCTeam
	if err := c.Bind(&t); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := s.mg.UpdateBRCTeam(int(id), int(season), t); err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) joinBRCTeam(c echo.Context) error {
	// Perform Authorization Checks
	claims := extractClaims(c)
	if err := isAuthenticated(claims); err != nil {
		return s.handleError(c, err)
	}

	var req struct {
		JoinKey  string
		SeasonID int
		StudentID int
	}
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	t, err := s.mg.GetBRCTeamByJoinKey(req.JoinKey, req.SeasonID)
	if err != nil {
		return s.handleError(c, err)
	}

	if err := s.mg.JoinBRCTeam(t.ID, t.SeasonID, req.StudentID); err != nil {
		return s.handleError(c, err)
	}
	return c.NoContent(http.StatusNoContent)
}

func (s *Server) leaveBRCTeam(c echo.Context) error {
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
	studentStr := c.Param("student")
	studentID, err := strconv.ParseInt(studentStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Perform Authorization Checks
	claims := extractClaims(c)
	if err := isAuthenticated(claims); err != nil {
		return s.handleError(c, err)
	}

	t, err := s.mg.GetBRCTeam(int(id), int(season))
	if err != nil {
		return s.handleError(c, err)
	}

	if err := permitCoachActions(extractClaims(c), t.Team); err != nil && int(userID) != claims.User.ID {
		return c.String(http.StatusUnauthorized, "You are not authorized to remove that person")
	}

	if err := s.mg.LeaveBRCTeam(t.ID, t.SeasonID, studentID); err != nil {
		return s.handleError(c, err)
	}
	return c.NoContent(http.StatusNoContent)
}
