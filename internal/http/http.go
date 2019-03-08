package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/BESTRobotics/registry/internal/mechgreg"
	"github.com/BESTRobotics/registry/internal/token"
	"github.com/BESTRobotics/registry/web"
)

// New returns a new http.Server or dies trying.
func New(mg MechGreg, tkn *token.RSATokenService) (*Server, error) {
	s := Server{
		mg:  mg,
		tkn: tkn,
	}
	s.g = gin.New()

	if viper.GetBool("dev.cors") {
		log.Println("CORS is operating in dev mode")
		cfg := cors.DefaultConfig()
		cfg.AllowAllOrigins = true
		cfg.AllowHeaders = append(cfg.AllowHeaders, "Authorization")
		cfg.AllowMethods = append(cfg.AllowMethods, "DELETE")
		s.g.Use(cors.New(cfg))
	}

	s.g.Use(s.validateToken)

	s.g.GET("/status", s.statusPage)

	v1 := s.g.Group("v1/")
	{
		v1.GET("/users", s.getUsers)
		v1.POST("/users", s.newUser)
		v1.GET("/users/:uid", s.getUser)
		v1.PUT("/users/:uid", s.modUser)
		v1.GET("/users/:uid/capabilities", s.getUserCapabilities)
		v1.PUT("/users/:uid/capabilities", s.addUserCapability)
		v1.DELETE("/users/:uid/capabilities", s.delUserCapability)

		v1.GET("/token/:id", s.getToken)
		v1.GET("/token-inspect", s.inspectToken)

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
		v1.PUT("/hubs/:id/admins", s.addHubAdmin)
		v1.DELETE("/hubs/:id/admins/delete/:uid", s.delHubAdmin)

		v1.GET("/schools", s.getSchools)
		v1.POST("/schools", s.newSchool)
		v1.GET("/schools/:id", s.getSchool)
		v1.PUT("/schools/:id/update", s.modSchool)

		v1.GET("/teams", s.getTeams)
		v1.POST("/teams", s.newTeam)
		v1.GET("/teams/:id", s.getTeam)
		v1.PUT("/teams/:id", s.modTeam)
		v1.PUT("/teams/:id/school", s.setTeamSchool)
		v1.GET("/teams/:id/school", s.getTeamSchool)
		v1.GET("/teams/:id/coach", s.getTeamCoach)
		v1.PUT("/teams/:id/coach", s.setTeamCoach)
		v1.PUT("/teams/:id/mentors", s.addTeamMentor)
		v1.DELETE("/teams/:id/mentors/:uid", s.delTeamMentor)
		v1.PUT("/teams/:id/home", s.setTeamHome)
		v1.GET("/teams/:id/home", s.getTeamHome)
		v1.PUT("/teams/:id/deactivate", s.deactivateTeam)
		v1.PUT("/teams/:id/activate", s.activateTeam)

		v1.GET("/events", s.getEvents)
		v1.POST("/events", s.newEvent)
		v1.GET("/events/:id", s.getEvent)
		v1.PUT("/events/:id", s.modEvent)
	}

	return &s, nil
}

// Serve is the entrypoint to the http package that serves the API.
func (s *Server) Serve() {
	if viper.GetBool("dev.extweb") {
		log.Println("###########################################")
		log.Println("#      HARD TO DEBUG PROBLEM WARNING      #")
		log.Println("###########################################")
		log.Println("Using external webroot", viper.GetString("dev.webroot"))
		s.g.StaticFS("/app", http.Dir(viper.GetString("dev.webroot")))
	} else {
		log.Println("Using built in webapp")
		log.Println("If a crash happens rebuild after running go generate ./...")
		s.g.StaticFS("/app", web.HTTP)
	}

	bindstr := fmt.Sprintf("%s:%d",
		viper.GetString("http.bind"),
		viper.GetInt("http.port"))

	s.g.Run(bindstr)
}

func (s *Server) statusPage(c *gin.Context) {
	c.String(http.StatusOK, "System OK")
}

func (s *Server) handleError(c *gin.Context, err error) {
	switch err.(type) {
	case *mechgreg.ConstraintError:
		cerr, ok := err.(*mechgreg.ConstraintError)
		if !ok {
			break
		}
		c.AbortWithStatusJSON(cerr.Code(), err)
		return
	case *mechgreg.InternalError:
		ierr, ok := err.(*mechgreg.InternalError)
		if !ok {
			break
		}
		c.AbortWithStatusJSON(ierr.Code(), err)
		return
	case AuthError:
		aerr, ok := err.(AuthError)
		if !ok {
			break
		}
		c.AbortWithStatusJSON(aerr.Code(), err)
		return
	}
	c.AbortWithError(http.StatusInternalServerError, err)
	return
}
