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

// GetBRCTeams returns all BRCTeams that a team owns.  This is useful
// to list all the teams on a landing page.
func (mg *MechanicalGreg) GetBRCTeams(teamID int) ([]models.BRCTeam, error) {
	var out []models.BRCTeam

	err := mg.s.Find("TeamID", teamID, &out)
	switch err {
	case nil:
		break
	case storm.ErrNotFound:
		return []models.BRCTeam{}, nil
	default:
		return nil, NewInternalError("An unspecified internal error has occured", err, http.StatusInternalServerError)
	}

	for i := range out {
		mg.populateBRCTeam(&out[i])
	}
	return out, nil
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

// JoinBRCTeam handles the adding of a user to a BRCTeam.Roster.
func (mg *MechanicalGreg) JoinBRCTeam(teamID, seasonID, userID int) error {
	return mg.handleBRCTeamJoinLeave(teamID, seasonID, userID, true)
}

// LeaveBRCTeam handles the adding of a user to a BRCTeam.Roster.
func (mg *MechanicalGreg) LeaveBRCTeam(teamID, seasonID, userID int) error {
	return mg.handleBRCTeamJoinLeave(teamID, seasonID, userID, false)
}


func (mg *MechanicalGreg) handleBRCTeamJoinLeave(teamID, seasonID, userID int, join bool) error {
	t, err := mg.GetBRCTeam(teamID, seasonID)
	if err != nil {
		return err
	}

	t.Roster = patchUserSlice(t.Roster, join, models.User{ID: userID})

	return mg.updateBRCTeam(t.ID, t.SeasonID, t)
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

	mg.populateBRCTeam(&b)

	return b, nil
}

func (mg *MechanicalGreg) populateBRCTeam(t *models.BRCTeam) error {
	team, err := mg.GetTeam(t.TeamID)
	if err != nil {
		return err
	}
	t.Team = team

	season, err := mg.GetSeason(t.SeasonID)
	if err != nil {
		return err
	}
	t.Season = season

	return nil
}
