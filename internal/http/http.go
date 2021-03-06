package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"

	"github.com/BESTRobotics/registry/internal/mail"
	"github.com/BESTRobotics/registry/internal/mechgreg"
	"github.com/BESTRobotics/registry/internal/token"
	"github.com/BESTRobotics/registry/web"
)

func init() {
	viper.SetDefault("core.url", "http://localhost:8080/")
}

// New returns a new http.Server or dies trying.
func New(mg MechGreg, tkn *token.RSATokenService, po mail.Mailer) (*Server, error) {
	s := Server{
		mg:  mg,
		tkn: tkn,
		po:  po,
	}
	s.ws = echo.New()

	if viper.GetBool("dev.cors") {
		log.Println("CORS is operating in dev mode")
		cors := middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		})
		s.ws.Use(cors)
	}

	s.ws.Use(s.validateToken)

	s.ws.GET("/", func(c echo.Context) error { return c.Redirect(http.StatusPermanentRedirect, "/app/index.html") })

	s.ws.GET("/status", s.statusPage)

	s.ws.POST("v1/account/register/local", s.registerLocalUser)
	s.ws.GET("v1/account/activate/:token", s.activateUser)
	s.ws.GET("v1/account/local/reset/:email", s.requestLocalPasswordReset)
	s.ws.GET("v1/account/local/rpass/:token", s.resetLocalPassword)
	s.ws.POST("v1/account/login/local", s.loginLocalUser)
	s.ws.POST("v1/account/token/renew", s.renewToken)

	s.ws.GET("v1/users", s.getUsers)
	s.ws.POST("v1/users", s.newUser)
	s.ws.GET("v1/users/:uid", s.getUser)
	s.ws.PUT("v1/users/:uid", s.modUser)
	s.ws.GET("v1/users/:uid/profile", s.getProfile)
	s.ws.POST("v1/users/:uid/profile", s.setProfile)
	s.ws.GET("v1/users/:uid/capabilities", s.getUserCapabilities)
	s.ws.PUT("v1/users/:uid/capabilities", s.addUserCapability)
	s.ws.DELETE("v1/users/:uid/capabilities", s.delUserCapability)
	s.ws.GET("v1/users/:uid/students", s.getStudents)
	s.ws.GET("v1/users/:uid/students/:sid", s.getStudent)
	s.ws.POST("v1/users/:uid/students", s.newStudent)
	s.ws.POST("v1/users/:uid/students/:sid", s.modStudent)

	s.ws.GET("v1/seasons", s.getSeasons)
	s.ws.POST("v1/seasons", s.newSeason)
	s.ws.GET("v1/seasons/:id", s.getSeason)
	s.ws.POST("v1/seasons/:id", s.modSeason)

	s.ws.GET("v1/hubs", s.getHubs)
	s.ws.POST("v1/hubs", s.newHub)
	s.ws.GET("v1/hubs/:id", s.getHub)
	s.ws.POST("v1/hubs/:id", s.modHub)
	s.ws.GET("v1/hubs/:id/teams", s.getTeamsForHub)

	s.ws.GET("v1/hubs/:id/brc", s.getBRCHubs)
	s.ws.POST("v1/hubs/:id/brc/:season", s.registerBRCHub)
	s.ws.GET("/v1/hubs/:id/brc/:season", s.getBRCHub)
	s.ws.POST("/v1/hubs/:id/brc/:season/update", s.updateBRCHub)

	s.ws.GET("v1/teams", s.getTeams)
	s.ws.POST("v1/teams", s.newTeam)
	s.ws.GET("v1/teams/:id", s.getTeam)
	s.ws.POST("v1/teams/:id", s.modTeam)

	s.ws.GET("v1/teams/:id/brc", s.getBRCTeams)
	s.ws.POST("v1/teams/:id/brc/:season", s.registerBRCTeam)
	s.ws.GET("v1/teams/:id/brc/:season", s.getBRCTeam)
	s.ws.PUT("v1/teams/:id/brc/:season/update", s.updateBRCTeam)
	s.ws.DELETE("v1/teams/:id/brc/:season/:student", s.leaveBRCTeam)

	s.ws.POST("v1/brc/join", s.joinBRCTeam)
	s.ws.GET("v1/brc/:season/hubs", s.getBRCHubsForSeason)
	s.ws.GET("v1/brc/teams/bystudent/:sid", s.getBRCTeamsByStudent)

	s.ws.GET("v1/events", s.getEvents)
	s.ws.POST("v1/events", s.newEvent)
	s.ws.GET("v1/events/:id", s.getEvent)
	s.ws.PUT("v1/events/:id", s.modEvent)

	return &s, nil
}

// Serve is the entrypoint to the http package that serves the API.
func (s *Server) Serve() {
	if viper.GetBool("dev.extweb") {
		log.Println("###########################################")
		log.Println("#      HARD TO DEBUG PROBLEM WARNING      #")
		log.Println("###########################################")
		log.Println("Using external webroot", viper.GetString("dev.webroot"))
		s.ws.Static("/app/*", viper.GetString("dev.webroot"))
	} else {
		log.Println("Using built in webapp")
		log.Println("If a crash happens rebuild after running go generate ./...")
		s.ws.GET("/app/*", echo.WrapHandler(http.StripPrefix("/app/", web.Handler)))
	}

	bindstr := fmt.Sprintf("%s:%d",
		viper.GetString("http.bind"),
		viper.GetInt("http.port"))

	s.ws.Start(bindstr)
}

func (s *Server) statusPage(c echo.Context) error {
	return c.String(http.StatusOK, "System OK")
}

func (s *Server) handleError(c echo.Context, err error) error {
	fmt.Println(err)
	switch err.(type) {
	case *mechgreg.ConstraintError:
		cerr, ok := err.(*mechgreg.ConstraintError)
		if !ok {
			break
		}
		return c.JSON(cerr.Code(), err)
	case *mechgreg.InternalError:
		ierr, ok := err.(*mechgreg.InternalError)
		if !ok {
			break
		}
		return c.JSON(ierr.Code(), err)
	case *mechgreg.AuthError:
		aerr, ok := err.(*mechgreg.AuthError)
		if !ok {
			break
		}
		return c.JSON(aerr.Code(), err)
	case AuthError:
		aerr, ok := err.(AuthError)
		if !ok {
			break
		}
		return c.JSON(aerr.Code(), err)
	}
	return c.JSON(http.StatusInternalServerError, err)
}
