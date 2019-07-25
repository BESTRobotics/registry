package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/BESTRobotics/registry/internal/models"
)

func (s *Server) getStudents(c echo.Context) error {
	// Fetch the User from the request
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())

	}

	// Perform authorization check
	if err := isAuthenticated(extractClaims(c)); err != nil {
		return c.String(http.StatusUnauthorized, "You must be logged in to do that")
	}

	out, err := s.mg.GetStudents(int(uid))
	if err != nil {
		return s.handleError(c, err)
	}
	return c.JSON(http.StatusOK, out)
}

func (s *Server) getStudent(c echo.Context) error {
	// Fetch the User from the request
	sidStr := c.Param("sid")
	sid, err := strconv.ParseInt(sidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())

	}

	// Perform authorization check
	if err := isAuthenticated(extractClaims(c)); err != nil {
		return c.String(http.StatusUnauthorized, "You must be logged in to do that")
	}

	out, err := s.mg.GetStudent(int(sid))
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusOK, out)
}

func (s *Server) newStudent(c echo.Context) error {
	// Fetch the User from the request
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())

	}

	// Perform authorization check; the student must be created
	// under the namespace of the logged in user.
	if err := isAuthenticated(extractClaims(c)); err != nil || int(uid) != extractClaims(c).User.ID {
		return c.String(http.StatusUnauthorized, "You must be logged in to do that")
	}

	var st models.Student
	if err := c.Bind(&st); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	st, err = s.mg.PutStudent(int(uid), st)
	if err != nil {
		return s.handleError(c, err)
	}

	return c.JSON(http.StatusCreated, st)
}

func (s *Server) modStudent(c echo.Context) error {
	// Fetch the User from the request
	uidStr := c.Param("uid")
	uid, err := strconv.ParseInt(uidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())

	}
	sidStr := c.Param("sid")
	sid, err := strconv.ParseInt(sidStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())

	}

	// Perform authorization check; the student must be either
	// under the namespace of the logged in user, or they must be
	// capable of editing users at large.
	if err := canModUser(extractClaims(c), int(uid)); err != nil && int(uid) != extractClaims(c).User.ID {
		return c.String(http.StatusUnauthorized, "You must be logged in to do that")
	}
	// Make sure the edit in question is on a student that's owned
	// by this account, otherwise the above check will allow
	// editing any student.
	set, err := s.mg.GetStudents(int(uid))
	if err != nil {
		return s.handleError(c, err)
	}
	owned := false
	for _, st := range set {
		if st.UserID == int(uid) {
			owned = true
		}
	}
	if err := canModUser(extractClaims(c), int(uid)); err != nil && owned == false {
		return c.String(http.StatusUnauthorized, "You can only modify students on your own account!")
	}

	var st models.Student
	if err := c.Bind(&st); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	st.ID = int(sid)

	if _, err := s.mg.PutStudent(int(uid), st); err != nil {
		return s.handleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}
