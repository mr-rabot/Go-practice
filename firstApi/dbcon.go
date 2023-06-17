package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func conDb() {
	db, err := gorm.Open(sqlite.Open("sdata.db"), &gorm.Config{})

	if err != nil {
		panic("failed to Connect database")
	}
	err = db.AutoMigrate(&Students{})
	if err != nil {
		return
	}

	DB = db

}
