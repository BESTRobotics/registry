package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/BESTRobotics/registry/internal/models"
	"github.com/BESTRobotics/registry/internal/token"
)

type involvements struct {
	Hubs  []models.Hub
	Teams []models.Team
}

func (s *Server) getInvolvements(userID int) (involvements, error) {
	var invs involvements
	var err error

	invs.Hubs, err = s.mg.GetHubsForUser(userID)
	if err != nil {
		return involvements{}, err
	}

	invs.Teams, err = s.mg.GetTeamsForUser(userID)
	if err != nil {
		return involvements{}, err
	}

	return invs, nil
}

func (s *Server) getToken(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	token, err := s.generateToken(int(id), token.GetConfig())
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, token)
}

func (s *Server) validateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tknStr := c.Request().Header.Get("Authorization")

		// If there was no token, don't try to extract it.  In the
		// case that there is a need to use authenticating
		// information, it will be obvious in later functions that
		// there is no authinfo in the context.
		if tknStr == "" {
			return next(c)
		}

		claims, err := s.tkn.Validate(tknStr)
		if err != nil {
			log.Println(err)
			status := struct {
				Message string
				Cause   string
			}{
				Message: "Bad token",
				Cause:   fmt.Sprint(err),
			}
			return c.JSON(http.StatusUnauthorized, status)
		}

		c.Set("authinfo", claims)
		return next(c)
	}
}

func (s *Server) inspectToken(c echo.Context) error {
	claims := extractClaims(c)
	return c.JSON(http.StatusOK, claims)
}

// generateToken figures out the contents of a token with the given
// configuration.  This is meant as a general purpose function to get
// tokens.
func (s *Server) generateToken(id int, cfg token.Config) (string, error) {
	log.Println("Requesting token for ID", id)

	invs, err := s.getInvolvements(id)
	if err != nil {
		return "", err
	}

	user, err := s.mg.GetUser(id)
	if err != nil {
		return "", err
	}

	var hubIDs []int
	for i := range invs.Hubs {
		hubIDs = append(hubIDs, invs.Hubs[i].ID)
	}

	var teamIDs []int
	for i := range invs.Teams {
		teamIDs = append(teamIDs, invs.Teams[i].ID)
	}

	claims := token.Claims{
		User:  user,
		Hubs:  hubIDs,
		Teams: teamIDs,
	}
	return s.tkn.Generate(claims, cfg)
}

func (s *Server) renewToken(c echo.Context) error {
	// TODO: This extends the lifetime of the token, but to do
	// otherwise means we'll need to fix things in the underlying
	// token system.
	claims := extractClaims(c)
	if err := isAuthenticated(claims); err != nil {
		return s.handleError(c, err)
	}

	tkn, err := s.generateToken(claims.User.ID, token.GetConfig())
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, tkn)
}
