package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to conenct to database!")
	}

	err = database.AutoMigrate(&Todo{})
	if err != nil {
		return
	}

	DB = database
}
