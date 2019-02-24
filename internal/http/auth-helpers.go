package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/BESTRobotics/registry/internal/token"
)

// AuthError represents all authorization errors.
type AuthError struct {
	Message string
	Cause   string
}

func (a AuthError) Error() string {
	return a.Message
}

// Code returns the http status code that this error should represent
// (Always 401)
func (a AuthError) Code() int {
	return http.StatusUnauthorized
}

func newAuthError(m, c string) AuthError {
	return AuthError{
		Message: m,
		Cause:   c,
	}
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
