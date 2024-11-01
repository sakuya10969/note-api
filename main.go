// main.go
package main

import (
    "note-api/config"
    "note-api/controllers"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    config.ConnectDatabase()

    r.POST("/notes", controllers.CreateNoteHandler)
    r.GET("/notes", controllers.GetNotesHandler)
    r.GET("/notes/:id", controllers.GetNoteHandler)
    r.PUT("/notes/:id", controllers.UpdateNoteHandler)
    r.DELETE("/notes/:id", controllers.DeleteNoteHandler)

    r.Run(":8080")
}
