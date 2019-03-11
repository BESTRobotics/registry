package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

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

func (s *Server) getToken(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := s.generateToken(int(id), token.GetConfig())
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, token)
}

func (s *Server) validateToken(c *gin.Context) {
	tknStr := c.GetHeader("authorization")

	// If there was no token, don't try to extract it.  In the
	// case that there is a need to use authenticating
	// information, it will be obvious in later functions that
	// there is no authinfo in the context.
	if tknStr == "" {
		return
	}

	claims, err := s.tkn.Validate(tknStr)
	if err != nil {
		status := struct {
			Message string
			Cause   string
		}{
			Message: "Bad token",
			Cause:   fmt.Sprint(err),
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, status)
	}

	c.Set("authinfo", claims)
}

func (s *Server) inspectToken(c *gin.Context) {
	claims := extractClaims(c)
	c.JSON(http.StatusOK, claims)
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
