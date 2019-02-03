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

		v1.GET("/seasons", s.getSeasons)
		v1.POST("/seasons", s.newSeason)
		v1.GET("/seasons/:id", s.getSeason)
		v1.PUT("/seasons/:id/update", s.modSeason)
		v1.PUT("/seasons/:id/archive", s.archiveSeason)

		v1.GET("/hubs", s.getHubs)
		v1.POST("/hubs", s.newHub)
		v1.GET("/hubs/:id", s.getHub)
		v1.PUT("/hubs/:id/update", s.modHub)
		v1.PUT("/hubs/:id/deactivate", s.deactivateHub)
		v1.PUT("/hubs/:id/activate", s.activateHub)
		v1.GET("/hubs/:id/director", s.getHubDirector)
		v1.PUT("/hubs/:id/director", s.setHubDirector)
		v1.PUT("/hubs/:id/admin", s.addHubAdmin)
		v1.DELETE("/hubs/:id/admin", s.delHubAdmin)
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
