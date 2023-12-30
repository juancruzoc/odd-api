package initializers

import "example/odd-api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}