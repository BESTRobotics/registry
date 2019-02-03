package http

import (
	"github.com/gin-gonic/gin"

	"github.com/BESTRobotics/registry/internal/models"
)

// The Server represents the type that all methods are bound on within
// the webserver.
type Server struct {
	mg MechGreg
	g  *gin.Engine
}

// The MechGreg or "Mechanical Greg" interface defines all the actions
// we might want to be able to take on the server.  This allows us
// some level of isolation between layers if we need to split the
// server up later on.
type MechGreg interface {
	NewUser(models.User) (int, error)
	GetUser(int) (models.User, error)
	ModUser(models.User) error
	GetUserPage(int, int) ([]models.User, error)

	NewSeason(models.Season) (int, error)
	GetSeason(int) (models.Season, error)
	GetSeasons(bool) ([]models.Season, error)
	ModSeason(models.Season) error
	ArchiveSeason(int) error

	NewHub(models.Hub) (int, error)
	GetHub(int) (models.Hub, error)
	GetHubs(bool) ([]models.Hub, error)
	ModHub(models.Hub) error
	DeactivateHub(int) error
	ActivateHub(int) error
	SetHubDirector(int, models.User) error
	GetHubDirector(int) (models.User, error)
	AddHubAdmin(int, models.User) error
	DelHubAdmin(int, models.User) error
}
