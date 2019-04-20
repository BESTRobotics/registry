package mechgreg

import (
	"log"
	"net/http"

	"github.com/asdine/storm"

	"github.com/BESTRobotics/registry/internal/models"
)

// NewTeam creates a new team.
func (mg *MechanicalGreg) NewTeam(t models.Team) (int, error) {
	t.BRIApproved = false

	t.HomeHubID = t.HomeHub.ID

	err := mg.s.Save(&t)
	switch err {
	case nil:
		return t.ID, nil
	case storm.ErrAlreadyExists:
		return 0, NewConstraintError("A team already exists with that ID", err, http.StatusConflict)
	default:
		return 0, NewInternalError("An unspecified internal error has occured", err, http.StatusInternalServerError)
	}
}

// GetTeam returns a single team.
func (mg *MechanicalGreg) GetTeam(id int) (models.Team, error) {
	var team models.Team

	err := mg.s.One("ID", id, &team)
	switch err {
	case nil:
		break
	case storm.ErrNotFound:
		return models.Team{}, NewConstraintError("No team exists with that ID", err, http.StatusNotFound)
	default:
		return models.Team{}, NewInternalError("An unspecified internal error has occured", err, http.StatusInternalServerError)
	}

	var coaches []models.User
	for i := range team.Coaches {
		coach, err := mg.GetUser(team.Coaches[i].ID)
		if err != nil {
			log.Println("Error loading mentor", err)
			continue
		}
		coaches = append(coaches, coach)
	}
	team.Coaches = coaches

	hub, err := mg.GetHub(team.HomeHub.ID)
	if err != nil {
		log.Println("Hub not loadable for team:", err)
	}
	team.HomeHub = hub

	return team, nil
}

// GetTeams returns all non-archived teams by default.  To include
// inactive teams set the parameter to true.
func (mg *MechanicalGreg) GetTeams(includeInactive bool) ([]models.Team, error) {
	var tmp []models.Team
	var out []models.Team
	var err error

	if includeInactive {
		err = mg.s.All(&tmp)
	} else {
		err = mg.s.Find("InactiveSince", models.DateTime{}, &tmp)
	}

	switch err {
	case nil:
		break
	case storm.ErrNotFound:
		// In this specific case, notfound actually means
		// there are no teams satisfying the query.
		return []models.Team{}, nil
	default:
		return nil, NewInternalError("An unspecified internal error has occured", err, http.StatusInternalServerError)
	}

	// This looks rather innefficient, but remember that the
	// backing boltdb is memory mapped, and the alternative would
	// be to duplicate code from the GetHub function.
	for i := range tmp {
		t, err := mg.GetTeam(tmp[i].ID)
		if err != nil {
			log.Println("Error loading team:", err)
			continue
		}
		out = append(out, t)
	}

	return out, nil
}

// GetTeamsForUser returns all the hubs that a use has power over in
// some way, shape, or form.
func (mg *MechanicalGreg) GetTeamsForUser(userID int) ([]models.Team, error) {
	involvements := make(map[int]models.Team)

	// Query for all hubs that have ever been, then figure out if
	// any of them have this person.
	teams, err := mg.GetTeams(true)
	if err != nil {
		return nil, err
	}

	// Iterate through the teams and find any that have this user
	// as a director or admin.  This isn't N^2 even though it
	// looks like it!
	for i := range teams {
		for j := range teams[i].Coaches {
			if teams[i].Coaches[j].ID == userID {
				involvements[teams[i].ID] = teams[i]
			}
		}
	}

	// Downconvert to just a list
	var out []models.Team
	for _, team := range involvements {
		out = append(out, team)
	}
	return out, nil
}

// GetTeamsForHub returns all the teams that are homed to a particular
// hub, active or not.
func (mg *MechanicalGreg) GetTeamsForHub(hubID int) ([]models.Team, error) {
	var tmp []models.Team
	var out []models.Team
	var err error

	err = mg.s.Find("HomeHubID", hubID, &tmp)

	switch err {
	case nil:
		break
	case storm.ErrNotFound:
		// In this specific case, notfound actually means
		// there are no teams satisfying the query.
		return []models.Team{}, nil
	default:
		return nil, NewInternalError("An unspecified internal error has occured", err, http.StatusInternalServerError)
	}

	// This looks rather innefficient, but remember that the
	// backing boltdb is memory mapped, and the alternative would
	// be to duplicate code from the GetTeam function.
	for i := range tmp {
		t, err := mg.GetTeam(tmp[i].ID)
		if err != nil {
			log.Println("Error loading team:", err)
			continue
		}
		out = append(out, t)
	}

	return out, nil
}

// ModTeam is like ModTeam, but doesn't null certain fields, making it
// suitable for internal use.
func (mg *MechanicalGreg) ModTeam(t models.Team) error {
	// Some fields need to be pulled for indexing.
	t.HomeHubID = t.HomeHub.ID

	err := mg.s.Update(&t)
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		return NewConstraintError("No team exists with that ID", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified internal error has occured", err, http.StatusInternalServerError)
	}
}
