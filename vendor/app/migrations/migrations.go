package migrations

import (
	_ "app/config"
	"app/database"
	"app/helpers/console"
)

func Migrate() {
	console.ConsoleSuccess("Migrating database")
	db := database.Connect()
	db.AutoMigrate(Users{}, Profiles{})
	console.ConsoleSuccess("Migration completed!")
}
