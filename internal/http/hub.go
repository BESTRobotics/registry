package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/BESTRobotics/registry/internal/models"
	"github.com/BESTRobotics/registry/internal/token"
)

func (s *Server) newHub(c *gin.Context) {
	// Perform Authorization Checks
	if err := canManageHubs(extractClaims(c)); err != nil {
		s.handleError(c, err)
		return
	}

	var hub models.Hub
	if err := c.ShouldBindJSON(&hub); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	id, err := s.mg.NewHub(hub)
	if err != nil {
		s.handleError(c, err)
		return
	}
	hub, err = s.mg.GetHub(id)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, hub)
}

func (s *Server) getHub(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	hub, err := s.mg.GetHub(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, hub)
}

func (s *Server) getHubs(c *gin.Context) {
	allStr := c.Query("include-inactive")
	all := false
	if allStr != "" {
		all = true
	}

	set, err := s.mg.GetHubs(all)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, set)
}

func (s *Server) modHub(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Perform Authorization Checks
	if err := canModHub(extractClaims(c), int(id)); err != nil {
		s.handleError(c, err)
		return
	}

	var hub models.Hub
	if err := c.ShouldBindJSON(&hub); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	hub.ID = int(id)
	// Admins has to be set by the correct endpoint, director has
	// to be set with the right endpoint as well.
	hub.Admins = nil
	hub.Director = models.User{}

	err = s.mg.ModHub(hub)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Server) deactivateHub(c *gin.Context) {
	// Perform Authorization Checks
	if err := canManageHubs(extractClaims(c)); err != nil {
		s.handleError(c, err)
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = s.mg.DeactivateHub(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Server) activateHub(c *gin.Context) {
	// Perform Authorization Checks
	if err := canManageHubs(extractClaims(c)); err != nil {
		s.handleError(c, err)
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = s.mg.ActivateHub(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Server) setHubDirector(c *gin.Context) {
	// Perform Authorization Checks
	if err := canManageHubs(extractClaims(c)); err != nil {
		s.handleError(c, err)
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	err = s.mg.SetHubDirector(int(id), user)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Server) getHubDirector(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	hub, err := s.mg.GetHubDirector(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, hub)
}

func (s *Server) addHubAdmin(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Perform Authorization Checks
	hub, err := s.mg.GetHub(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}
	if err := permitDirectorActions(extractClaims(c), hub); err != nil {
		s.handleError(c, err)
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	err = s.mg.AddHubAdmin(int(id), user)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Server) delHubAdmin(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Perform Authorization Checks
	hub, err := s.mg.GetHub(int(id))
	if err != nil {
		s.handleError(c, err)
		return
	}
	if err := permitDirectorActions(extractClaims(c), hub); err != nil {
		s.handleError(c, err)
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	err = s.mg.DelHubAdmin(int(id), user)
	if err != nil {
		s.handleError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

// canModHub checks whether appropriate claims are available to modify
// a hub.  General modifications can be done by a hub admin, so we
// check all hubs the user may be allowed to handle.
func canModHub(claims token.Claims, hubID int) error {
	if claims.IsEmpty() {
		return newAuthError("Unauthorized", "Claims are empty")
	}
	if claims.User.HasCapability(models.CapSuperAdmin) {
		return nil
	}

	for i := range claims.Hubs {
		if claims.Hubs[i] == hubID {
			return nil
		}
	}
	return newAuthError("Unauthorized", "You do not have the appropriate clearance to modify this hub!")
}

// canManageHubs tells whether or not the requestor is allowed to
// handle things like creating and archiving hubs, or setting the hub
// director.
func canManageHubs(claims token.Claims) error {
	if claims.IsEmpty() {
		return newAuthError("Unauthorized", "Claims are empty")
	}
	if !claims.User.HasCapability(models.CapHubAdmin) {
		return newAuthError("Unauthorized", "You must posess CapHubAdmin to do that!")
	}
	return nil
}

// isHubDirector figures out if this user has a claim as the director
// of this hub.
func permitDirectorActions(claims token.Claims, hub models.Hub) error {
	if claims.IsEmpty() {
		return newAuthError("Unauthorized", "Claims are empty")
	}
	if claims.User.HasCapability(models.CapHubAdmin) {
		// Short circuit if they're can HubAdmin
		return nil
	}
	for i := range claims.Hubs {
		if claims.Hubs[i] == hub.Director.ID {
			return nil
		}
	}
	return newAuthError("Unauthorized", "You must be a hub director to do that!")
}
