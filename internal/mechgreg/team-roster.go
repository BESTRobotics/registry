package mechgreg

import (
	"log"

	"github.com/asdine/storm"

	"github.com/BESTRobotics/registry/internal/models"
)

// NewTeamRoster creates a new roster and associates it to a team and a
// particular season.
func (mg *MechanicalGreg) NewTeamRoster(t models.TeamRoster) (int, error) {
	if _, err := mg.GetTeam(t.Team.ID); err != nil {
		return 0, err
	}

	if _, err := mg.GetSeason(t.Season.ID); err != nil {
		return 0, err
	}

	// We only save the IDs in this structure.
	t.Team = models.Team{ID: t.Team.ID}
	t.Season = models.Season{ID: t.Season.ID}

	switch mg.s.Save(&t) {
	case nil:
		return t.ID, nil
	case storm.ErrAlreadyExists:
		return 0, ErrResourceExists
	default:
		return 0, ErrInternal
	}
}

// GetTeamRoster returns a specific roster filled out completely.
func (mg *MechanicalGreg) GetTeamRoster(id int) (models.TeamRoster, error) {
	var roster models.TeamRoster

	switch mg.s.One("ID", id, &roster) {
	case nil:
		break
	case storm.ErrNotFound:
		return models.TeamRoster{}, ErrNoSuchResource
	default:
		return models.TeamRoster{}, ErrInternal
	}

	t, err := mg.GetTeam(roster.Team.ID)
	if err != nil {
		return models.TeamRoster{}, err
	}

	s, err := mg.GetSeason(roster.Season.ID)
	if err != nil {
		return models.TeamRoster{}, err
	}

	var members []models.User
	for i := range roster.Members {
		user, err := mg.GetUser(roster.Members[i].ID)
		if err != nil {
			log.Println("Error loading user:", err)
			continue
		}
		members = append(members, user)
	}

	roster.Team = t
	roster.Season = s
	roster.Members = members

	return roster, nil
}


// JoinTeam provides the mechanism by which a team roster gets its
// members.  The user is validated for existence, and the key is used
// to summon the roster to join.
func (mg *MechanicalGreg) JoinTeam(u models.User, k string) error {
	// Check user first, then get the roster.
	if _, err := mg.GetUser(u.ID); err != nil {
		return err
	}

	var r models.TeamRoster
	switch mg.s.One("JoinKey", k, &r) {
	case nil:
		break
	case storm.ErrNotFound:
		return ErrNoSuchResource
	default:
		return ErrInternal
	}

	// Patch in the new member.  This is idempotent to take care
	// of spam clicking a join button.
	r.Members = patchUserSlice(r.Members, true, u)

	return mg.modTeamRoster(r)
}

// LeaveTeam performs the opposite action of JoinTeam, but takes a
// rosterID since it should be known at this point.
func (mg *MechanicalGreg) LeaveTeam(rosterID int, u models.User) error {
	r, err := mg.GetTeamRoster(rosterID)
	if err != nil {
		return err
	}

	r.Members = patchUserSlice(r.Members, false, u)

	return mg.modTeamRoster(r)
}

// SetTeamRosterJoinKey sets a new key to join with.
func (mg *MechanicalGreg) SetTeamRosterJoinKey(rosterID int, k string) error {
	r, err := mg.GetTeamRoster(rosterID)
	if err != nil {
		return err
	}

	r.JoinKey = k
	return mg.modTeamRoster(r)
}

func (mg *MechanicalGreg) modTeamRoster(r models.TeamRoster) error {
	switch mg.s.Update(&r) {
	case nil:
		return nil
	case storm.ErrNotFound:
		return ErrNoSuchResource
	default:
		return ErrInternal
	}
}
