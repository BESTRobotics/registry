package mechgreg

import (
	"log"

	"github.com/BESTRobotics/registry/internal/models"
)

// NewSeason creates a new season within the seasons table.
func (mg *MechanicalGreg) NewSeason(s models.Season) (int, error) {
	var ts models.Season
	ns := models.Season{Name: s.Name}
	if err := mg.rb.DB.Where(&ns).First(&ts).Error; err == nil {
		// Already exists
		return 0, ErrResourceExists
	}

	if err := mg.rb.DB.Create(&s).Error; err != nil {
		log.Println(err)
		return 0, ErrInternal
	}
	return s.ID, nil
}

// GetSeason requests a single season from the database via the
// season's ID.
func (mg *MechanicalGreg) GetSeason(id int) (models.Season, error) {
	var season models.Season
	if err := mg.rb.DB.First(&season, id).Error; err != nil {
		log.Println(err)
		return models.Season{}, ErrNoSuchResource
	}
	return season, nil
}

// GetSeasons returns all seasons that are not archived.  To return
// *all* seasons the all parameter should be set to true.
func (mg *MechanicalGreg) GetSeasons(all bool) ([]models.Season, error) {
	out := []models.Season{}
	var err error

	if all {
		err = mg.rb.DB.Find(&out).Error
	} else {
		err = mg.rb.DB.Not(&models.Season{Archived: true}).Find(&out).Error
	}
	if err != nil {
		log.Println(err)
		return nil, ErrInternal
	}
	return out, nil
}

// ModSeason updates an existing season.  An error is returned if the
// season requested does not exist.
func (mg *MechanicalGreg) ModSeason(s models.Season) error {
	var ts models.Season
	if err := mg.rb.DB.First(&ts, s.ID).Error; err != nil {
		return ErrNoSuchResource
	}

	if err := mg.rb.DB.Model(&models.Season{}).Updates(s).Error; err != nil {
		log.Println(err)
		return ErrInternal
	}
	return nil
}

// ArchiveSeason freezes a season in time.  Seasons can't ever be
// deleted, but archiving them removes them from the set that's shown
// by default.
func (mg *MechanicalGreg) ArchiveSeason(id int) error {
	var ts models.Season
	if err := mg.rb.DB.First(&ts, id).Error; err != nil {
		return ErrNoSuchResource
	}

	if err := mg.rb.DB.Model(&models.Season{}).Updates(models.Season{ID: id, Archived: true}).Error; err != nil {
		log.Println(err)
		return ErrInternal
	}
	return nil
}
