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

	log.Println("Requesting token for ID", id)

	invs, err := s.getInvolvements(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	user, err := s.mg.GetUser(int(id))
	if err != nil {
		s.handleError(c, err)
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

	token, err := s.tkn.Generate(claims, token.GetConfig())
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

// extractClaims fishes out the token and asserts the type back to
// something sane.
func extractClaims(c *gin.Context) token.Claims {
	cl, exists := c.Get("authinfo")
	if !exists {
		// Bail out now with empty claims.  Its safe to bail
		// with no error because empty claims can't be used
		// for anything.  Auth checks implicitly fail.
		return token.Claims{}
	}
	claims, ok := cl.(token.Claims)
	if !ok {
		log.Println("Something that wasn't authinfo was in the token:", cl)
		return token.Claims{}
	}
	return claims
}
