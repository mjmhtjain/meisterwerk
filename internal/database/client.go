package database

import (
	"github.com/mjmhtjain/meisterwerk/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresDB *gorm.DB = nil

func NewDBClient(config *config.DatabaseConfig) (*gorm.DB, error) {
	if postgresDB != nil {
		return postgresDB, nil
	}

	db, err := gorm.Open(postgres.Open(config.GetDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Get the underlying *sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Run migrations
	if err := RunMigrations(sqlDB); err != nil {
		return nil, err
	}

	postgresDB = db

	return postgresDB, nil
}
