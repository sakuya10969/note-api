// services/noteService.go
package services

import (
    "note-api/config"
    "note-api/models"
)

func CreateNote(note *models.Note) error {
    return config.DB.Create(note).Error
}

func GetNotes() ([]models.Note, error) {
    var notes []models.Note
    err := config.DB.Find(&notes).Error
    return notes, err
}

func GetNoteByID(id uint) (models.Note, error) {
    var note models.Note
    err := config.DB.First(&note, id).Error
    return note, err
}

func UpdateNote(note *models.Note) error {
    return config.DB.Save(note).Error
}

func DeleteNoteByID(id uint) error {
    return config.DB.Delete(&models.Note{}, id).Error
}
