package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/BESTRobotics/registry/internal/models"
	"github.com/BESTRobotics/registry/internal/token"
)

func (s *Server) newSeason(c echo.Context) error {
	if err := canManageSeasons(extractclaims(c)); err != nil {
		return s.handleError(c, err)
	}

	var season models.Season
	if err := c.Bind(&season); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	id, err := s.mg.NewSeason(season)
	if err != nil {
		return s.handleError(c, err)
	}
	season, err = s.mg.GetSeason(id)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusCreated, season)
}

func (s *Server) getSeason(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	season, err := s.mg.GetSeason(int(id))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, season)
}

func (s *Server) getSeasons(c echo.Context) error {
	allStr := c.QueryParam("all")
	all := false
	if allStr != "" {
		all = true
	}

	set, err := s.mg.GetSeasons(all)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, set)
}

func (s *Server) modSeason(c echo.Context) error {
	if err := canManageSeasons(extractclaims(c)); err != nil {
		return s.handleError(c, err)
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var season models.Season
	if err := c.Bind(&season); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	season.ID = int(id)

	err = s.mg.ModSeason(season)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) archiveSeason(c echo.Context) error {
	if err := canManageSeasons(extractclaims(c)); err != nil {
		return s.handleError(c, err)
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = s.mg.ArchiveSeason(int(id))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

// canManageSeasons handles the actions around seasons.  These actions
// are reserved to the superadmins because of how many other parts of
// the system key off of whether or not a season is live.
func canManageSeasons(claims token.Claims) error {
	if claims.IsEmpty() {
		return newAuthError("Unauthorized", "Claims are empty")
	}
	if claims.User.HasCapability(models.CapSuperAdmin) {
		return nil
	}
	return newAuthError("Unauthorized", "You must be a SuperAdmin to do that!")
}
