package migrations

import (
	"test-go/config"
	"test-go/models"
)

func CarMigration() {
	config.DB.AutoMigrate(&models.Car{})
}
