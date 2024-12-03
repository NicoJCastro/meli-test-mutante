package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	err = database.AutoMigrate(&DNARecord{})
	if err != nil {
		return fmt.Errorf("failed to auto-migrate database: %w", err)
	}

	DB = database
	return nil
}

func SaveDNARecord(dna DNARecord) error {
	return DB.Create(&dna).Error
}

type DNARecord struct {
	ID       uint   `gorm:"primaryKey"`
	Sequence string `gorm:"unique"`
	IsMutant bool
}
