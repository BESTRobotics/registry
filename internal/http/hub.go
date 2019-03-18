package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/BESTRobotics/registry/internal/models"
	"github.com/BESTRobotics/registry/internal/token"
)

func (s *Server) newHub(c echo.Context) error {
	// Perform Authorization Checks
	if err := canManageHubs(extractClaims(c)); err != nil {
		return s.handleError(c, err)
	}

	var hub models.Hub
	if err := c.Bind(&hub); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	id, err := s.mg.NewHub(hub)
	if err != nil {
		return s.handleError(c, err)
	}
	hub, err = s.mg.GetHub(id)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusCreated, hub)
}

func (s *Server) getHub(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	hub, err := s.mg.GetHub(int(id))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, hub)
}

func (s *Server) getHubs(c echo.Context) error {
	allStr := c.QueryParam("include-inactive")
	all := false
	if allStr != "" {
		all = true
	}

	set, err := s.mg.GetHubs(all)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, set)
}

func (s *Server) modHub(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Perform Authorization Checks
	if err := canModHub(extractClaims(c), int(id)); err != nil {
		return s.handleError(c, err)
	}

	var hub models.Hub
	if err := c.Bind(&hub); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	hub.ID = int(id)
	// Admins has to be set by the correct endpoint, director has
	// to be set with the right endpoint as well.
	hub.Admins = nil
	hub.Director = models.User{}

	err = s.mg.ModHub(hub)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) deactivateHub(c echo.Context) error {
	// Perform Authorization Checks
	if err := canManageHubs(extractClaims(c)); err != nil {
		return s.handleError(c, err)
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = s.mg.DeactivateHub(int(id))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) activateHub(c echo.Context) error {
	// Perform Authorization Checks
	if err := canManageHubs(extractClaims(c)); err != nil {
		return s.handleError(c, err)
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = s.mg.ActivateHub(int(id))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) setHubDirector(c echo.Context) error {
	// Perform Authorization Checks
	if err := canManageHubs(extractClaims(c)); err != nil {
		return s.handleError(c, err)
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = s.mg.SetHubDirector(int(id), user)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) getHubDirector(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	hub, err := s.mg.GetHubDirector(int(id))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, hub)
}

func (s *Server) addHubAdmin(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Perform Authorization Checks
	hub, err := s.mg.GetHub(int(id))
	if err != nil {
		return s.handleError(c, err)
	}
	if err := permitDirectorActions(extractClaims(c), hub); err != nil {
		return s.handleError(c, err)
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = s.mg.AddHubAdmin(int(id), user)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) delHubAdmin(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	// Perform Authorization Checks
	hub, err := s.mg.GetHub(int(id))
	if err != nil {
		return s.handleError(c, err)
	}
	if err := permitDirectorActions(extractClaims(c), hub); err != nil {
		return s.handleError(c, err)
	}

	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = s.mg.DelHubAdmin(int(id), models.User{ID: int(uid)})
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
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
