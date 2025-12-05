package dbops

import (
	"auth-service/loggerconfig"
	"auth-service/models"
	)

func MigrateTables() {
	err := DB.AutoMigrate(&models.UserProfile{})
	if err!=nil {
		loggerconfig.Panic("unable to run migration for user_profile")
	}
}
