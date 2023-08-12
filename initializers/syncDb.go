package initializers

import "gostud/models"

func SyncDb() {
	DB.AutoMigrate(
		&models.User{},
	)

}
