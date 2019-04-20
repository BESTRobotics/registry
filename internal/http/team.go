package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/BESTRobotics/registry/internal/models"
	"github.com/BESTRobotics/registry/internal/token"
)

func (s *Server) newTeam(c echo.Context) error {
	// Perform Authorization Checks
	clms := extractClaims(c)
	if err := isAuthenticated(clms); err != nil {
		return s.handleError(c, err)
	}

	var team models.Team
	if err := c.Bind(&team); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Make sure the requesting user is the first coach.
	team.Coaches = []models.User{models.User{ID: clms.User.ID}}

	// Actually create the team.
	id, err := s.mg.NewTeam(team)
	if err != nil {
		return s.handleError(c, err)
	}

	team, err = s.mg.GetTeam(id)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusCreated, team)
}

func (s *Server) getTeam(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	team, err := s.mg.GetTeam(int(id))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, team)
}

func (s *Server) getTeams(c echo.Context) error {
	allStr := c.QueryParam("include-inactive")
	all := false
	if allStr != "" {
		all = true
	}

	set, err := s.mg.GetTeams(all)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, set)
}

func (s *Server) getTeamsForHub(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	set, err := s.mg.GetTeamsForHub(int(id))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, set)
}

func (s *Server) modTeam(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var team models.Team
	if err := c.Bind(&team); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Perform Authorization Checks
	if err := canModTeam(extractClaims(c), int(id)); err != nil {
		return s.handleError(c, err)
	}

	team.ID = int(id)

	err = s.mg.ModTeam(team)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

// canModTeam checks whether the appropriate claims are available to
// modify the team.  General modifications can be done by a mentor, so
// we check all teams the user may be allowed to do something on.
func canModTeam(claims token.Claims, teamID int) error {
	if claims.IsEmpty() {
		return newAuthError("Unauthorized", "Claims are empty")
	}
	if claims.User.HasCapability(models.CapSuperAdmin) {
		return nil
	}
	if claims.User.HasCapability(models.CapTeamAdmin) {
		return nil
	}

	for i := range claims.Teams {
		if claims.Teams[i] == teamID {
			return nil
		}
	}
	return newAuthError("Unauthorized", "You do not have the appropriate clearance to modify this team!")
}

// canManageTeams tells whether or not the requestor is allowed to
// handle things like creating or archiving teams or setting the team
// coach.
func canManageTeams(claims token.Claims) error {
	if claims.IsEmpty() {
		return newAuthError("Unauthorized", "Claims are empty")
	}
	if claims.User.HasCapability(models.CapSuperAdmin) {
		return nil
	}
	if !claims.User.HasCapability(models.CapTeamAdmin) {
		return newAuthError("Unauthorized", "You must posess CapTeamAdmin to do that!")
	}
	return nil
}

// permitCoachActions figures out if the requesting user is the coach
// of this team and can take actions that are reserved for the coach.
func permitCoachActions(claims token.Claims, team models.Team) error {
	if claims.IsEmpty() {
		return newAuthError("Unauthorized", "Claims are empty")
	}
	if claims.User.HasCapability(models.CapSuperAdmin) {
		return nil
	}
	if claims.User.HasCapability(models.CapTeamAdmin) {
		return nil
	}
	for i := range claims.Teams {
		for j := range team.Coaches {
			if claims.Teams[i] == team.Coaches[j].ID {
				return nil
			}
		}
	}
	return newAuthError("Unauthorized", "You must be a team coach to do that!")
}

// permitHomeHubActions provides a way for the home hub to take some
// other managerial actions on the team that are not permitted to the
// team itself.
func permitHomeHubActions(claims token.Claims, team models.Team) error {
	if claims.IsEmpty() {
		return newAuthError("Unauthorized", "Claims are empty")
	}
	if claims.User.HasCapability(models.CapSuperAdmin) {
		return nil
	}
	if claims.User.HasCapability(models.CapTeamAdmin) {
		return nil
	}

	for i := range claims.Hubs {
		if claims.Hubs[i] == team.HomeHub.ID {
			return nil
		}
	}
	return newAuthError("Unauthorized", "You must be part of the home hub to do that!")
}
