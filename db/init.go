package db

import (
	"fmt"
	"sre-goapi/config"
	logger "sre-goapi/utils"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetAppDB() *gorm.DB {
	dbURL := config.DBUrl[9:] // removing sqlite://

	appDB, err := gorm.Open(sqlite.Open(dbURL), &gorm.Config{})
	if err != nil {
		fmt.Print("ERROR: Error connecting to database")
		logger.Panic("Database connection failed")
	}

	return appDB
}
