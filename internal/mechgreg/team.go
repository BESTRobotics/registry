package mechgreg

import (
	"log"
	"net/http"
	"time"

	"github.com/asdine/storm"

	"github.com/BESTRobotics/registry/internal/mail"
	"github.com/BESTRobotics/registry/internal/models"
)

// NewTeam creates a new team.
func (mg *MechanicalGreg) NewTeam(t models.Team) (int, error) {
	// These fields need special handling and have to be set via
	// dedicated interfaces.
	t.Coach = nil
	t.Mentors = nil
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

	var mentors []models.User
	for i := range team.Mentors {
		mentor, err := mg.GetUser(team.Mentors[i].ID)
		if err != nil {
			log.Println("Error loading mentor", err)
			continue
		}
		mentors = append(mentors, mentor)
	}
	team.Mentors = mentors

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
		for j := range teams[i].Coach {
			if teams[i].Coach[j].ID == userID {
				involvements[teams[i].ID] = teams[i]
			}
		}
		for j := range teams[i].Mentors {
			if teams[i].Mentors[j].ID == userID {
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

// ModTeam modifies a team.  It protects certain specialized fields
// that require different mechanisms to set.
func (mg *MechanicalGreg) ModTeam(team models.Team) error {
	team.Coach = nil
	team.Mentors = nil

	// BRIApproved needs to be pulled and fed across any updates.
	current, err := mg.GetTeam(team.ID)
	if err != nil {
		return err
	}
	team.BRIApproved = current.BRIApproved

	return mg.modTeam(team)
}

// modTeam is like ModTeam, but doesn't null certain fields, making it
// suitable for internal use.
func (mg *MechanicalGreg) modTeam(t models.Team) error {
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

// AddTeamCoach adds a coach to the team which can perform most, but
// not all, actions that the coach can perform.
func (mg *MechanicalGreg) AddTeamCoach(id int, u models.User) error {
	team, err := mg.GetTeam(id)
	if err != nil {
		return err
	}

	team.Coach = patchUserSlice(team.Coach, true, u)
	if err := mg.modTeam(team); err != nil {
		return err
	}

	l := mail.NewLetter()
	l.AddTo(mail.UserToAddress(u))
	l.Subject = "You're now a coach!"
	l.Body = "Thanks for helping to further STEM education, you're now a team coach."

	return mg.po.SendMail(l)
}

// DelTeamCoach removes a coach from the listed team.
func (mg *MechanicalGreg) DelTeamCoach(id int, u models.User) error {
	team, err := mg.GetTeam(id)
	if err != nil {
		return err
	}

	coachs := patchUserSlice(team.Coach, false, u)
	err = mg.s.UpdateField(&models.Team{ID: id}, "Coach", coachs)
	switch err {
	case nil:
		break
	case storm.ErrNotFound:
		return NewConstraintError("No team exists with that ID", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified internal error has occured", err, http.StatusInternalServerError)
	}

	l := mail.NewLetter()
	l.AddTo(mail.UserToAddress(u))
	l.Subject = "You've been removed from a team"
	l.Body = "Thanks for your time, you're no longer listed as a coach on " + team.StaticName + "."

	return mg.po.SendMail(l)
}

// AddTeamMentor adds a mentor to the team which can perform most, but
// not all, actions that the coach can perform.
func (mg *MechanicalGreg) AddTeamMentor(id int, u models.User) error {
	team, err := mg.GetTeam(id)
	if err != nil {
		return err
	}

	team.Mentors = patchUserSlice(team.Mentors, true, u)
	if err := mg.modTeam(team); err != nil {
		return err
	}

	l := mail.NewLetter()
	l.AddTo(mail.UserToAddress(u))
	l.Subject = "You're now a mentor!"
	l.Body = "Thanks for helping to further STEM education, you're now a team mentor."

	return mg.po.SendMail(l)
}

// DelTeamMentor removes a mentor from the listed team.
func (mg *MechanicalGreg) DelTeamMentor(id int, u models.User) error {
	team, err := mg.GetTeam(id)
	if err != nil {
		return err
	}

	mentors := patchUserSlice(team.Mentors, false, u)
	err = mg.s.UpdateField(&models.Team{ID: id}, "Mentors", mentors)
	switch err {
	case nil:
		break
	case storm.ErrNotFound:
		return NewConstraintError("No team exists with that ID", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified internal error has occured", err, http.StatusInternalServerError)
	}

	l := mail.NewLetter()
	l.AddTo(mail.UserToAddress(u))
	l.Subject = "You've been removed from a team"
	l.Body = "Thanks for your time, you're no longer listed as a mentor on " + team.StaticName + "."

	return mg.po.SendMail(l)
}

// DeactivateTeam allows us to mark a team as "dead".  When this
// happens the team is still in the system, but doesn't show up in
// most queries anymore.
func (mg *MechanicalGreg) DeactivateTeam(id int) error {
	return mg.modTeam(models.Team{ID: id, InactiveSince: models.DateTime(time.Now())})
}

// ActivateTeam brings a team back from an inactive state.
func (mg *MechanicalGreg) ActivateTeam(id int) error {
	// Needs to use UpdateField in order to explicitely zero the
	// value.
	err := mg.s.UpdateField(&models.Team{ID: id}, "InactiveSince", time.Time{})
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		return NewConstraintError("No team exists with that ID", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified internal error has occured", err, http.StatusInternalServerError)
	}
}

// SetTeamHome sets the hub that this team calls home.  All teams have
// a home hub as their point of contact with the rest of BEST, and
// this is the one for this team.
func (mg *MechanicalGreg) SetTeamHome(id int, h models.Hub) error {
	if _, err := mg.GetHub(h.ID); err != nil {
		return err
	}

	return mg.modTeam(models.Team{ID: id, HomeHub: models.Hub{ID: h.ID}})
}

// GetTeamHome returns the home hub for the team.
func (mg *MechanicalGreg) GetTeamHome(id int) (models.Hub, error) {
	h, err := mg.GetHub(id)
	if err != nil {
		return models.Hub{}, err
	}
	return h, nil
}

// ApproveTeam approves a team which unlocks other abilities such as
// BRC registration acceptance.
func (mg *MechanicalGreg) ApproveTeam(id int) error {
	t := models.Team{
		ID:          id,
		BRIApproved: true,
	}

	return mg.modTeam(t)
}
