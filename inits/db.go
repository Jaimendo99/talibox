package inits

import (
	"fmt"
	"talibox/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewConnection(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	return db, nil
}

func InitDB() {
	db, err := NewConnection("test1.db")
	if err != nil {
		fmt.Printf("error opening sqlite:" + err.Error())
		return
	}

	DB = db
}

func Migrate(db *gorm.DB) error {
	modelsToMigrate := []interface{}{
		&models.GenreCatalog{},
		&models.Professional{},
		&models.Movie{},
		&models.Cast{},
		&models.User{},
		&models.WeeklyGrosses{},
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			return err // Or handle the error in a way that suits your application's requirements
		}
	}
	return nil
}
