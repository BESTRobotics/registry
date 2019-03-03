package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/BESTRobotics/registry/internal/models"
)

func (s *Server) newSchool(c *gin.Context) {
	// Perform Authorization Checks
	if err := canManageTeams(extractClaims(c)); err != nil {
		s.handleError(c, err)
		return
	}

	var school models.School
	if err := c.ShouldBindJSON(&school); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	id, err := s.mg.NewSchool(school)
	if err != nil {
		s.handleError(c, err)
		return
	}
	school, err = s.mg.GetSchool(id)
	if err != nil {
		s.handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, school)
}

func (s *Server) getSchool(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	school, err := s.mg.GetSchool(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, school)
}

func (s *Server) getSchools(c *gin.Context) {
	schools, err := s.mg.GetSchools()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, schools)
}

func (s *Server) modSchool(c *gin.Context) {
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

	school.ID = int(id)

	err = s.mg.ModSchool(school)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
