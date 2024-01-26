package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database *gorm.DB
var err error

func DatabaseInit() {
	database, err = gorm.Open(sqlite.Open("myapp.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func DB() *gorm.DB {
	return database
}
