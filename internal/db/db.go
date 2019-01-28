package db

import (
	"log"

	"github.com/BESTRobotics/registry/internal/models"

	"github.com/jinzhu/gorm"
)

// Open an instance of the database.
func Open() (*gorm.DB, error) {
	log.Println("Aquiring database handle")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return nil, err
	}

	log.Println("Performing schema auto-migration")
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Meta{})
	return db, nil
}
