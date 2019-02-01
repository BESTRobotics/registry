package mechgreg

import (
	"github.com/BESTRobotics/registry/internal/models"
)

// NewSeason creates a new season within the seasons table.
func (mg *MechanicalGreg) NewSeason(s models.Season) (int, error) {
	return 0, nil
}

// GetSeason requests a single season from the database via the
// season's ID.
func (mg *MechanicalGreg) GetSeason(id int) (models.Season, error) {
	return models.Season{}, nil
}

// GetSeasons returns all seasons that are not archived.  To return
// *all* seasons the all parameter should be set to true.
func (mg *MechanicalGreg) GetSeasons(all bool) ([]models.Season, error) {
	return nil, nil
}

// ModSeason updates an existing season.  An error is returned if the
// season requested does not exist.
func (mg *MechanicalGreg) ModSeason(s models.Season) error {
	return nil
}

// ArchiveSeason freezes a season in time.  Seasons can't ever be
// deleted, but archiving them removes them from the set that's shown
// by default.
func (mg *MechanicalGreg) ArchiveSeason(id int) error {
	return nil
}
