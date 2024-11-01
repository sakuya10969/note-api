// controllers/noteController.go
package controllers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "note-api/models"
    "note-api/services"
)

func CreateNoteHandler(c *gin.Context) {
    var note models.Note
    if err := c.ShouldBindJSON(&note); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := services.CreateNote(&note); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, note)
}

func GetNotesHandler(c *gin.Context) {
    notes, err := services.GetNotes()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, notes)
}

func GetNoteHandler(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
        return
    }

    note, err := services.GetNoteByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
        return
    }
    c.JSON(http.StatusOK, note)
}

func UpdateNoteHandler(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
        return
    }

    var note models.Note
    if err := c.ShouldBindJSON(&note); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    note.ID = uint(id)

    if err := services.UpdateNote(&note); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, note)
}

func DeleteNoteHandler(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
        return
    }

    if err := services.DeleteNoteByID(uint(id)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}
