package mechgreg

import (
	"net/http"

	"github.com/asdine/storm"

	"github.com/BESTRobotics/registry/internal/models"
)

// NewSeason creates a new season within the seasons table.
func (mg *MechanicalGreg) NewSeason(s models.Season) (int, error) {
	err := mg.s.Save(&s)
	switch err {
	case nil:
		return s.ID, nil
	case storm.ErrAlreadyExists:
		return 0, NewConstraintError("A season with that name already exists", err, http.StatusConflict)
	default:
		return 0, NewInternalError("An unspecified failure has occured", err, http.StatusInternalServerError)
	}
}

// GetSeason requests a single season from the database via the
// season's ID.
func (mg *MechanicalGreg) GetSeason(id int) (models.Season, error) {
	var season models.Season

	err := mg.s.One("ID", id, &season)
	switch err {
	case nil:
		return season, nil
	case storm.ErrNotFound:
		return models.Season{}, NewConstraintError("No season exists with that ID", err, http.StatusNotFound)
	default:
		return models.Season{}, NewInternalError("An unspecified failure has occured", err, http.StatusInternalServerError)
	}
}

// GetSeasons returns all seasons that are not archived.  To return
// *all* seasons the all parameter should be set to true.
func (mg *MechanicalGreg) GetSeasons(all bool) ([]models.Season, error) {
	var out []models.Season
	var err error

	if all {
		err = mg.s.All(&out)
	} else {
		err = mg.s.Find("Archived", false, &out)
	}
	switch err {
	case nil:
		return out, nil
	case storm.ErrNotFound:
		return []models.Season{}, nil
	default:
		return nil, err
	}
}

// ModSeason updates an existing season.  An error is returned if the
// season requested does not exist.
func (mg *MechanicalGreg) ModSeason(s models.Season) error {
	err := mg.s.Update(&s)
	switch err {
	case nil:
		return nil
	case storm.ErrNotFound:
		return NewConstraintError("No season exists with that ID", err, http.StatusNotFound)
	default:
		return NewInternalError("An unspecified failure has occured", err, http.StatusInternalServerError)
	}
}

// ArchiveSeason freezes a season in time.  Seasons can't ever be
// deleted, but archiving them removes them from the set that's shown
// by default.
func (mg *MechanicalGreg) ArchiveSeason(id int) error {
	return mg.ModSeason(models.Season{ID: id, Archived: true})
}
