package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func newDatabase(dbtype, dbconnection string, debug bool) *gorm.DB {
	//db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open(dbtype, dbconnection)

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate if something was changed
	db.AutoMigrate(&dbState{}, &dbTemperature{}, &dbHumidity{})
	db.LogMode(debug)

	return db
}
