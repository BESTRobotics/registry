package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/BESTRobotics/registry/internal/mechgreg"
	"github.com/BESTRobotics/registry/internal/models"
)

func (s *Server) newSchool(c *gin.Context) {
	var school models.School
	if err := c.ShouldBindJSON(&school); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	id, err := s.mg.NewSchool(school)
	switch err {
	case nil:
		break
	case mechgreg.ErrResourceExists:
		c.AbortWithError(http.StatusConflict, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.Set("Location", fmt.Sprintf("/v1/schools/%d", id))
	c.Status(http.StatusCreated)
}

func (s *Server) getSchool(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	school, err := s.mg.GetSchool(int(id))
	switch err {
	case nil:
		break
	case mechgreg.ErrNoSuchResource:
		c.AbortWithError(http.StatusNotFound, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
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

	switch s.mg.ModSchool(school) {
	case nil:
		break
	case mechgreg.ErrNoSuchResource:
		c.AbortWithError(http.StatusNotFound, err)
		return
	default:
		c.AbortWithError(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	c.Status(http.StatusNoContent)
}
