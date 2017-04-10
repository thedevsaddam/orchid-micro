package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //postgres/sqlite3
	"os"
)

// Open returns a DB reference for a data source.
func Connect() *gorm.DB {
	driverName := os.Getenv("DATABASE_DRIVER")
	connection := ""
	//check database driver //Note: This not a good practice. Highly encourage to do this initialization by yourself
	switch driverName {
	case "mysql":
		connection = os.Getenv("DATABASE_USERNAME") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" + os.Getenv("DATABASE_NAME") + "?charset=" + os.Getenv("DATABASE_CHARSET") + "&parseTime=True&loc=Local"
	case "postgres":
		connection = "host=" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + "user=" + os.Getenv("DATABASE_USERNAME") + " dbname=" + os.Getenv("DATABASE_NAME") + " sslmode=disable password=" + os.Getenv("DATABASE_PASSWORD")
	case "sqlite3":
		connection = os.Getenv("DATABASE_HOST")
	}
	db, err := gorm.Open(driverName, connection)
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
