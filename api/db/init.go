package db

import (
	"fmt"
	"sre-goapi/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetAppDB() *gorm.DB {
	dbURL := config.DBUrl[9:] // removing sqlite://

	appDB, err := gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Print("ERROR: Error connecting to database")
	}

	return appDB
}
