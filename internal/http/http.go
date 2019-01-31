package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// New returns a new http.Server or dies trying.
func New(mg MechGreg) (*Server, error) {
	s := Server{
		mg: mg,
	}
	s.g = gin.New()

	s.g.GET("/status", s.statusPage)

	v1 := s.g.Group("v1/")
	{
		v1.GET("/users", s.getUsers)
		v1.POST("/users", s.newUser)
		v1.GET("/users/:uid", s.getUser)
		v1.PUT("/users/:uid", s.modUser)
		v1.DELETE("/users/:uid", s.delUser)

		v1.GET("/seasons", s.getSeasons)
		v1.POST("/seasons", s.newSeason)
		v1.GET("/seasons/:id", s.getSeason)
		v1.PUT("/seasons/update/:id", s.modSeason)
		v1.PUT("/seasons/archive/:id", s.archiveSeason)
	}

	return &s, nil
}

// Serve is the entrypoint to the http package that serves the API.
func (s *Server) Serve() {
	if viper.GetBool("dev.extweb") {
		log.Println("Using external webroot", viper.GetString("dev.webroot"))
		s.g.StaticFS("/app", http.Dir(viper.GetString("dev.webroot")))
	}

	bindstr := fmt.Sprintf("%s:%d",
		viper.GetString("http.bind"),
		viper.GetInt("http.port"))

	s.g.Run(bindstr)
}

func (s *Server) statusPage(c *gin.Context) {
	c.String(http.StatusOK, "System OK")
}
