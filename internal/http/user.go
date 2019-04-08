package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/BESTRobotics/registry/internal/models"
	"github.com/BESTRobotics/registry/internal/token"
)

func (s *Server) newUser(c echo.Context) error {
	// Deserialize the user
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Attempt to create the user
	uid, err := s.mg.NewUser(user)
	if err != nil {
		return s.handleError(c, err)

	}
	user, err = s.mg.GetUser(uid)
	if err != nil {
		return s.handleError(c, err)

	}

	return c.JSON(http.StatusCreated, user)
}

func (s *Server) getUser(c echo.Context) error {
	// Fetch the User from the request
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())

	}

	user, err := s.mg.GetUser(int(uid))
	if err != nil {
		return s.handleError(c, err)

	}

	return c.JSON(http.StatusOK, user)
}

func (s *Server) setProfile(c echo.Context) error {
	// Fetch the User from the request
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())

	}

	// Perform Authentication Checks
	if err := canModUser(extractClaims(c), int(uid)); err != nil && extractClaims(c).User.ID != int(uid) {
		return s.handleError(c, err)
	}

	var p models.UserProfile
	if err := c.Bind(&p); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if err := s.mg.SetUserProfile(int(uid), p); err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusOK)
}

func (s *Server) getProfile(c echo.Context) error {
	// Fetch the User from the request
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())

	}

	// Perform authorization checks
	if err := isAuthenticated(extractClaims(c)); err != nil {
		return c.String(http.StatusUnauthorized, "Insufficient authorization")
	}

	p, err := s.mg.GetUserProfile(int(uid))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, p)
}

func (s *Server) getUsers(c echo.Context) error {
	page := int64(0)
	count := int64(25)
	pageStr := c.QueryParam("page")
	countStr := c.QueryParam("count")
	var err error

	if pageStr != "" {
		page, err = strconv.ParseInt(pageStr, 10, 32)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())

		}
	}

	if countStr != "" {
		count, err = strconv.ParseInt(countStr, 10, 32)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
	}

	set, err := s.mg.GetUserPage(int(page), int(count))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, set)
}

func (s *Server) modUser(c echo.Context) error {
	// Fetch the User from the request
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())

	}

	// Perform Authentication Checks
	if err := canModUser(extractClaims(c), int(uid)); err != nil {
		return s.handleError(c, err)

	}

	// Deserialize the user
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user.ID = int(uid)
	user.Username = ""
	user.Capabilities = nil

	err = s.mg.ModUser(user)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) getUserCapabilities(c echo.Context) error {
	// Fetch the User from the request
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := s.mg.GetUser(int(uid))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, user.Capabilities)
}

func (s *Server) addUserCapability(c echo.Context) error {
	// Perform Authentication Checks
	if err := canModCapabilities(extractClaims(c)); err != nil {
		return s.handleError(c, err)
	}

	// Fetch the User from the request
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var cap models.Capability
	if err := c.Bind(&cap); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := s.mg.GetUser(int(uid))
	if err != nil {
		return s.handleError(c, err)
	}
	user.GrantCapability(cap)
	err = s.mg.ModUser(user)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) delUserCapability(c echo.Context) error {
	// Perform Authentication Checks
	if err := canModCapabilities(extractClaims(c)); err != nil {
		return s.handleError(c, err)

	}

	// Fetch the User from the request
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	var cap models.Capability
	if err := c.Bind(&cap); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	user, err := s.mg.GetUser(int(uid))
	if err != nil {
	}

	user.RemoveCapability(cap)
	err = s.mg.ModUser(user)
	if err != nil {
		return s.handleError(c, err)

	}
	return c.NoContent(http.StatusNoContent)
}

// canModUser checks if a user modification is allowed.  This will
// return true in the case of the included user matching the requested
// ID, or the USER_ADMIN capability being present.
func canModUser(claims token.Claims, id int) error {
	if claims.IsEmpty() {
		return newAuthError("Unauthorized", "Claims are empty")
	}

	if claims.User.ID != id {
		return newAuthError("Unauthorized", "You are not authorized to modify this user")
	}
	return nil
}

// canModCapabilities takes care of returning an error if the user
// lacks sufficient power to make a change to capabilities.  Since
// capabilities gate the ability to gain special powers, the ability
// to add and remove capabilities from them is only available to a
// SuperAdmin.
func canModCapabilities(claims token.Claims) error {
	if claims.IsEmpty() {
		return newAuthError("Unauthorized", "Claims are empty")
	}

	if !claims.User.HasCapability(models.CapSuperAdmin) {
		return newAuthError("Unauthorized", "Only a SuperAdmin can modify capabilities!")
	}
	return nil
}
