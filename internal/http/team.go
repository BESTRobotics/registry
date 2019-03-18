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
	if err := canManageTeams(extractClaims(c)); err != nil {
		return s.handleError(c, err)
	}

	var team models.Team
	if err := c.Bind(&team); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

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

func (s *Server) modTeam(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Perform Authorization Checks
	if err := canModTeam(extractClaims(c), int(id)); err != nil {
		return s.handleError(c, err)
	}

	var team models.Team
	if err := c.Bind(&team); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	team.ID = int(id)

	err = s.mg.ModTeam(team)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) setTeamSchool(c echo.Context) error {
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

	err = s.mg.SetTeamSchool(int(id), school)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) getTeamSchool(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	school, err := s.mg.GetTeamSchool(int(id))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, school)
}

func (s *Server) setTeamCoach(c echo.Context) error {
	// Perform Authorization Checks
	if err := canManageTeams(extractClaims(c)); err != nil {
		return s.handleError(c, err)
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return s.handleError(c, err)
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = s.mg.SetTeamCoach(int(id), user)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) getTeamCoach(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := s.mg.GetTeamCoach(int(id))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (s *Server) addTeamMentor(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Perform Authorization Checks
	team, err := s.mg.GetTeam(int(id))
	if err != nil {
		return s.handleError(c, err)
	}
	if err := permitCoachActions(extractClaims(c), team); err != nil {
		return s.handleError(c, err)
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = s.mg.AddTeamMentor(int(id), user)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) delTeamMentor(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Perform Authorization Checks
	team, err := s.mg.GetTeam(int(id))
	if err != nil {
		return s.handleError(c, err)
	}
	if err := permitCoachActions(extractClaims(c), team); err != nil {
		return s.handleError(c, err)
	}

	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = s.mg.DelTeamMentor(int(id), models.User{ID: int(uid)})
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) setTeamHome(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Perform Authorization Checks
	team, err := s.mg.GetTeam(int(id))
	if err != nil {
		return s.handleError(c, err)
	}
	if err := permitHomeHubActions(extractClaims(c), team); err != nil {
		return s.handleError(c, err)
	}

	var hub models.Hub
	err = c.Bind(&hub)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = s.mg.SetTeamHome(int(id), hub)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) getTeamHome(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	hub, err := s.mg.GetTeamHome(int(id))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, hub)
}

func (s *Server) deactivateTeam(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Perform Authorization Checks
	team, err := s.mg.GetTeam(int(id))
	if err != nil {
		return s.handleError(c, err)
	}
	if err := permitHomeHubActions(extractClaims(c), team); err != nil {
		return s.handleError(c, err)
	}

	if err := s.mg.DeactivateTeam(int(id)); err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) activateTeam(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Perform Authorization Checks
	team, err := s.mg.GetTeam(int(id))
	if err != nil {
		return s.handleError(c, err)
	}
	if err := permitHomeHubActions(extractClaims(c), team); err != nil {
		return s.handleError(c, err)
	}

	if err = s.mg.ActivateTeam(int(id)); err != nil {
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
		if claims.Teams[i] == team.Coach.ID {
			return nil
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
