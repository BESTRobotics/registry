package http

import (
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
