package token

import (
	"github.com/BESTRobotics/registry/internal/models"
)

// Claims represents the claims that are made to the system.
type Claims struct {
	User  models.User
	Hubs  []int
	Teams []int
}
