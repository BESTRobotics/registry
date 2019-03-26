package http

import (
	"github.com/labstack/echo"

	"github.com/BESTRobotics/registry/internal/mail"
	"github.com/BESTRobotics/registry/internal/models"
	"github.com/BESTRobotics/registry/internal/token"
)

// The Server represents the type that all methods are bound on within
// the webserver.
type Server struct {
	mg  MechGreg
	tkn *token.RSATokenService
	ws  *echo.Echo
	po  mail.Mailer
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
	UsernameExists(string) (models.User, error)
	SetUserPassword(string, string) error
	CheckUserPassword(string, string) error

	NewSeason(models.Season) (int, error)
	GetSeason(int) (models.Season, error)
	GetSeasons(bool) ([]models.Season, error)
	ModSeason(models.Season) error
	ArchiveSeason(int) error

	NewHub(models.Hub) (int, error)
	GetHub(int) (models.Hub, error)
	GetHubs(bool) ([]models.Hub, error)
	GetHubsForUser(int) ([]models.Hub, error)
	ModHub(models.Hub) error
	DeactivateHub(int) error
	ActivateHub(int) error
	SetHubDirector(int, models.User) error
	GetHubDirector(int) (models.User, error)
	AddHubAdmin(int, models.User) error
	DelHubAdmin(int, models.User) error

	RegisterBRCHub(int, int) (int, error)
	GetBRCHub(int, int) (models.BRCHub, error)
	GetBRCHubs(int) ([]models.BRCHub, error)
	UpdateBRCHub(int, int, models.BRCHub) error
	ApproveBRCHub(int, int, bool) error

	NewTeam(models.Team) (int, error)
	GetTeam(int) (models.Team, error)
	GetTeams(bool) ([]models.Team, error)
	GetTeamsForUser(int) ([]models.Team, error)
	GetTeamsForHub(int) ([]models.Team, error)
	ModTeam(models.Team) error
	SetTeamCoach(int, models.User) error
	GetTeamCoach(int) (models.User, error)
	AddTeamMentor(int, models.User) error
	DelTeamMentor(int, models.User) error
	SetTeamHome(int, models.Hub) error
	GetTeamHome(int) (models.Hub, error)
	DeactivateTeam(int) error
	ActivateTeam(int) error
	ApproveTeam(int) error

	NewEvent(models.Event) (int, error)
	ModEvent(models.Event) error
	GetEvent(int) (models.Event, error)
	GetEvents() ([]models.Event, error)
}
