package migrations

import (
	"backend/internal/models"

	"gorm.io/gorm"
)

type MassAutoMigrationResult struct {
	Status bool
	Err    error
}

func AutoMigrateAll(db *gorm.DB) MassAutoMigrationResult {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		result := MassAutoMigrationResult{
			Status: false,
			Err:    err,
		}
		return result
	}
	result := MassAutoMigrationResult{
		Status: true,
		Err:    err,
	}
	return result
}
