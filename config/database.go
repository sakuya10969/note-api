package config

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"note-api/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("note.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := models.MigrateNotes(database); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = database
}