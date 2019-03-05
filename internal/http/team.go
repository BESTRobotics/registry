package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/BESTRobotics/registry/internal/models"
	"github.com/BESTRobotics/registry/internal/token"
)

func (s *Server) newTeam(c *gin.Context) {
	// Perform Authorization Checks
	if err := canManageTeams(extractClaims(c)); err != nil {
		s.handleError(c, err)
		return
	}

	var team models.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	id, err := s.mg.NewTeam(team)
	if err != nil {
		s.handleError(c, err)
		return
	}
	team, err = s.mg.GetTeam(id)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, team)
}

func (s *Server) getTeam(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	team, err := s.mg.GetTeam(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, team)
}

func (s *Server) getTeams(c *gin.Context) {
	allStr := c.Query("include-inactive")
	all := false
	if allStr != "" {
		all = true
	}

	set, err := s.mg.GetTeams(all)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, set)
}

func (s *Server) modTeam(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Perform Authorization Checks
	if err := canModTeam(extractClaims(c), int(id)); err != nil {
		s.handleError(c, err)
		return
	}

	var team models.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	team.ID = int(id)

	err = s.mg.ModTeam(team)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Server) setTeamSchool(c *gin.Context) {
	// Perform Authorization Checks
	if err := canManageTeams(extractClaims(c)); err != nil {
		s.handleError(c, err)
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var school models.School
	if err := c.ShouldBindJSON(&school); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	err = s.mg.SetTeamSchool(int(id), school)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Server) getTeamSchool(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	school, err := s.mg.GetTeamSchool(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, school)
}

func (s *Server) setTeamCoach(c *gin.Context) {
	// Perform Authorization Checks
	if err := canManageTeams(extractClaims(c)); err != nil {
		s.handleError(c, err)
		return
	}

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

	err = s.mg.SetTeamCoach(int(id), user)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Server) getTeamCoach(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := s.mg.GetTeamCoach(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (s *Server) addTeamMentor(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Perform Authorization Checks
	team, err := s.mg.GetTeam(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}
	if err := permitCoachActions(extractClaims(c), team); err != nil {
		s.handleError(c, err)
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	err = s.mg.AddTeamMentor(int(id), user)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Server) delTeamMentor(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Perform Authorization Checks
	team, err := s.mg.GetTeam(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}
	if err := permitCoachActions(extractClaims(c), team); err != nil {
		s.handleError(c, err)
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	err = s.mg.DelTeamMentor(int(id), user)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Server) setTeamHome(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Perform Authorization Checks
	team, err := s.mg.GetTeam(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}
	if err := permitHomeHubActions(extractClaims(c), team); err != nil {
		s.handleError(c, err)
		return
	}

	var hub models.Hub
	err = c.ShouldBindJSON(&hub)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	err = s.mg.SetTeamHome(int(id), hub)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Server) getTeamHome(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	hub, err := s.mg.GetTeamHome(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, hub)
}

func (s *Server) deactivateTeam(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Perform Authorization Checks
	team, err := s.mg.GetTeam(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}
	if err := permitHomeHubActions(extractClaims(c), team); err != nil {
		s.handleError(c, err)
		return
	}

	err = s.mg.DeactivateTeam(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Server) activateTeam(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Perform Authorization Checks
	team, err := s.mg.GetTeam(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}
	if err := permitHomeHubActions(extractClaims(c), team); err != nil {
		s.handleError(c, err)
		return
	}

	err = s.mg.ActivateTeam(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
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
