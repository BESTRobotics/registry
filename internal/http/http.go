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
	s.ws.POST("v1/account/login/local", s.loginLocalUser)

	s.ws.GET("v1/users", s.getUsers)
	s.ws.POST("v1/users", s.newUser)
	s.ws.GET("v1/users/:uid", s.getUser)
	s.ws.PUT("v1/users/:uid", s.modUser)
	s.ws.GET("v1/users/:uid/capabilities", s.getUserCapabilities)
	s.ws.PUT("v1/users/:uid/capabilities", s.addUserCapability)
	s.ws.DELETE("v1/users/:uid/capabilities", s.delUserCapability)

	s.ws.GET("v1/token/:id", s.getToken)
	s.ws.GET("v1/token-inspect", s.inspectToken)

	s.ws.GET("v1/seasons", s.getSeasons)
	s.ws.POST("v1/seasons", s.newSeason)
	s.ws.GET("v1/seasons/:id", s.getSeason)
	s.ws.PUT("v1/seasons/:id/update", s.modSeason)
	s.ws.PUT("v1/seasons/:id/archive", s.archiveSeason)

	s.ws.GET("v1/hubs", s.getHubs)
	s.ws.POST("v1/hubs", s.newHub)
	s.ws.GET("v1/hubs/:id", s.getHub)
	s.ws.PUT("v1/hubs/:id/update", s.modHub)
	s.ws.PUT("v1/hubs/:id/deactivate", s.deactivateHub)
	s.ws.PUT("v1/hubs/:id/activate", s.activateHub)
	s.ws.GET("v1/hubs/:id/director", s.getHubDirector)
	s.ws.PUT("v1/hubs/:id/director", s.setHubDirector)
	s.ws.PUT("v1/hubs/:id/admins", s.addHubAdmin)
	s.ws.DELETE("v1/hubs/:id/admins/delete/:uid", s.delHubAdmin)
	s.ws.GET("v1/hubs/:id/teams", s.getTeamsForHub)

	s.ws.POST("v1/hubs/:id/brc/:season", s.registerBRCHub)
	s.ws.GET("/v1/hubs/:id/brc/:season", s.getBRCHub)
	s.ws.PUT("/v1/hubs/:id/brc/:season/update", s.updateBRCHub)
	s.ws.POST("/v1/hubs/:id/brc/:season/bri-approve", s.approveBRCHub)

	s.ws.GET("v1/teams", s.getTeams)
	s.ws.POST("v1/teams", s.newTeam)
	s.ws.GET("v1/teams/:id", s.getTeam)
	s.ws.PUT("v1/teams/:id", s.modTeam)
	s.ws.GET("v1/teams/:id/coach", s.getTeamCoach)
	s.ws.PUT("v1/teams/:id/coach", s.setTeamCoach)
	s.ws.PUT("v1/teams/:id/mentors", s.addTeamMentor)
	s.ws.DELETE("v1/teams/:id/mentors/:uid", s.delTeamMentor)
	s.ws.PUT("v1/teams/:id/home", s.setTeamHome)
	s.ws.GET("v1/teams/:id/home", s.getTeamHome)
	s.ws.PUT("v1/teams/:id/approve", s.approveTeam)
	s.ws.PUT("v1/teams/:id/deactivate", s.deactivateTeam)
	s.ws.PUT("v1/teams/:id/activate", s.activateTeam)

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
