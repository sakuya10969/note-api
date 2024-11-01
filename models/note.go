package models

import "gorm.io/gorm"

type Note struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
	Content string `json:"content"`
}

func MigrateNotes(db *gorm.DB) error {
	return db.AutoMigrate(&Note{})
}