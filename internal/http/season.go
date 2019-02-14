package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/BESTRobotics/registry/internal/models"
)

func (s *Server) newSeason(c *gin.Context) {
	var season models.Season
	if err := c.ShouldBindJSON(&season); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	id, err := s.mg.NewSeason(season)
	if err != nil {
		s.handleError(c, err)
		return
	}
	season, err = s.mg.GetSeason(id)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, season)
}

func (s *Server) getSeason(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	season, err := s.mg.GetSeason(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, season)
}

func (s *Server) getSeasons(c *gin.Context) {
	allStr := c.Query("all")
	all := false
	if allStr != "" {
		all = true
	}

	set, err := s.mg.GetSeasons(all)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, set)
}

func (s *Server) modSeason(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var season models.Season
	if err := c.ShouldBindJSON(&season); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	season.ID = int(id)

	err = s.mg.ModSeason(season)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Server) archiveSeason(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = s.mg.ArchiveSeason(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
