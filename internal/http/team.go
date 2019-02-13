package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/BESTRobotics/registry/internal/models"
)

func (s *Server) newTeam(c *gin.Context) {
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

	c.Set("Location", fmt.Sprintf("/v1/teams/%d", id))
	c.Status(http.StatusCreated)
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

	err = s.mg.ActivateTeam(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
