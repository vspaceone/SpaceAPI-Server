package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func newDatabase(dbtype, dbconnection string) *gorm.DB {
	//db, err := gorm.Open("sqlite3", "test.db")
	db, err := gorm.Open(dbtype, dbconnection)

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate if something was changed
	db.AutoMigrate(&state{}, &temperature{}, &humidity{})
	db.LogMode(true)

	return db
}
