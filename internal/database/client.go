package database

import (
	"github.com/mjmhtjain/meisterwerk/internal/config"
	"github.com/mjmhtjain/meisterwerk/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient(config *config.DatabaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.GetDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(&models.Quote{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
