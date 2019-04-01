package mechgreg

import (
	"log"
	"net/http"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"

	"github.com/BESTRobotics/registry/internal/models"
)

// RegisterBRCTeam creates a new BRCTeam resource and returns the created
// resource to the caller.
func (mg *MechanicalGreg) RegisterBRCTeam(teamID, seasonID int) (int, error) {
	if _, err := mg.GetTeam(teamID); err != nil {
		return 0, err
	}
	if _, err := mg.GetSeason(seasonID); err != nil {
		return 0, err
	}

	var brcteam models.BRCTeam
	query := mg.s.Select(q.And(q.Eq("TeamID", teamID), q.Eq("SeasonID", seasonID)))
	if err := query.First(&brcteam); err != storm.ErrNotFound {
		return -1, NewConstraintError("BRCTeam already exists", err, http.StatusPreconditionFailed)
	}

	newBRCTeam := &models.BRCTeam{
		TeamID:   teamID,
		SeasonID: seasonID,
	}

	err := mg.s.Save(newBRCTeam)
	switch err {
	case nil:
		return newBRCTeam.ID, nil
	case storm.ErrAlreadyExists:
		return 0, NewConstraintError("A BRCTeam with that ID already exists?", err, http.StatusConflict)
	default:
		log.Println("Error occured while trying to process BRCTeam Registration:", err)
		return 0, NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}
}

// GetBRCTeam loads a BRCTeam by numeric ID.
func (mg *MechanicalGreg) GetBRCTeam(teamID, seasonID int) (models.BRCTeam, error) {
	var brcteam models.BRCTeam
	query := mg.s.Select(q.And(q.Eq("TeamID", teamID), q.Eq("SeasonID", seasonID)))
	err := query.First(&brcteam)
	return mg.handleBRCTeamGet(brcteam, err)
}

// GetBRCTeamByJoinKey gets the number of the team that was requested.
func (mg *MechanicalGreg) GetBRCTeamByJoinKey(key string, seasonID int) (models.BRCTeam, error) {
	var brcteam models.BRCTeam
	query := mg.s.Select(q.And(q.Eq("JoinKey", key), q.Eq("SeasonID", seasonID)))
	err := query.First(&brcteam)
	return mg.handleBRCTeamGet(brcteam, err)
}

// GetBRCTeamBySymbol gets the number of the team by the symbol.
func (mg *MechanicalGreg) GetBRCTeamBySymbol(symbol string, seasonID int) (models.BRCTeam, error) {
	var brcteam models.BRCTeam
	query := mg.s.Select(q.And(q.Eq("Symbol", symbol), q.Eq("SeasonID", seasonID)))
	err := query.First(&brcteam)
	return mg.handleBRCTeamGet(brcteam, err)
}

// updateBRCTeam can update all fields on a BRCTeam struct.
func (mg *MechanicalGreg) updateBRCTeam(teamID, seasonID int, update models.BRCTeam) error {
	update.TeamID = teamID
	update.SeasonID = seasonID

	err := mg.s.Save(&update)
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		return NewConstraintError("This team did not participate in that season", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}
}

// UpdateBRCTeam allows fields to be set that are mutable after creation.
func (mg *MechanicalGreg) UpdateBRCTeam(teamID, seasonID int, update models.BRCTeam) error {
	update.Team = models.Team{}
	update.Season = models.Season{}
	update.TeamID = 0
	update.SeasonID = 0
	update.State = ""
	update.Roster = nil

	return mg.updateBRCTeam(teamID, seasonID, update)
}

// handleBRCTeamGet populates the remaining fields and processes the
// returned error.
func (mg *MechanicalGreg) handleBRCTeamGet(b models.BRCTeam, err error) (models.BRCTeam, error) {
	switch err {
	case nil:
		break
	case storm.ErrNotFound:
		return models.BRCTeam{}, NewConstraintError("This hub did not participate in that season", err, http.StatusNotFound)
	default:
		return models.BRCTeam{}, NewInternalError("An unspecified error has occured", err, http.StatusInternalServerError)
	}

	team, err := mg.GetTeam(b.TeamID)
	if err != nil {
		return models.BRCTeam{}, err
	}
	b.Team = team

	season, err := mg.GetSeason(b.SeasonID)
	if err != nil {
		return models.BRCTeam{}, err
	}
	b.Season = season

	return b, nil
}
