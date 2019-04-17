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
	GetUserByEMail(string) (models.User, error)
	ModUser(models.User) error
	GetUserPage(int, int) ([]models.User, error)
	FillUserProfile(*models.User) error
	GetUserProfile(int) (models.UserProfile, error)
	SetUserProfile(int, models.UserProfile) error
	SetUserPassword(int, string) error
	CheckUserPassword(int, string) error
	GetStudent(int) (models.Student, error)
	GetStudents(int) ([]models.Student, error)
	PutStudent(int, models.Student) error

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

	RegisterBRCTeam(int, int) (int, error)
	GetBRCTeams(int) ([]models.BRCTeam, error)
	GetBRCTeam(int, int) (models.BRCTeam, error)
	GetBRCTeamByJoinKey(string, int) (models.BRCTeam, error)
	GetBRCTeamBySymbol(string, int) (models.BRCTeam, error)
	UpdateBRCTeam(int, int, models.BRCTeam) error
	JoinBRCTeam(int, int, int) error
	LeaveBRCTeam(int, int, int) error

	NewEvent(models.Event) (int, error)
	ModEvent(models.Event) error
	GetEvent(int) (models.Event, error)
	GetEvents() ([]models.Event, error)
}
