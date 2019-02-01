package mechgreg

import (
	"github.com/asdine/storm"

	"github.com/BESTRobotics/registry/internal/models"
)

// NewSeason creates a new season within the seasons table.
func (mg *MechanicalGreg) NewSeason(s models.Season) (int, error) {
	switch mg.s.Save(&s) {
	case nil:
		return s.ID, nil
	case storm.ErrAlreadyExists:
		return 0, ErrResourceExists
	default:
		return 0, ErrInternal
	}
}

// GetSeason requests a single season from the database via the
// season's ID.
func (mg *MechanicalGreg) GetSeason(id int) (models.Season, error) {
	var season models.Season

	switch mg.s.One("ID", id, &season) {
	case nil:
		return season, nil
	case storm.ErrNotFound:
		return models.Season{}, ErrNoSuchResource
	default:
		return models.Season{}, ErrInternal
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
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModSeason updates an existing season.  An error is returned if the
// season requested does not exist.
func (mg *MechanicalGreg) ModSeason(s models.Season) error {
	switch mg.s.Update(&s) {
	case nil:
		return nil
	case storm.ErrNotFound:
		return ErrNoSuchResource
	default:
		return ErrInternal
	}
}

// ArchiveSeason freezes a season in time.  Seasons can't ever be
// deleted, but archiving them removes them from the set that's shown
// by default.
func (mg *MechanicalGreg) ArchiveSeason(id int) error {
	return mg.ModSeason(models.Season{ID: id, Archived: true})
}
