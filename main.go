package main

import (
    "notes-app/controllers"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("templates/*")

    r.GET("/", controllers.ListNotes)
    r.GET("/add", controllers.ShowAddNote)
    r.POST("/add", controllers.AddNote)
    r.GET("/edit/:id", controllers.ShowEditNote)
    r.POST("/edit/:id", controllers.UpdateNote)
    r.GET("/delete/:id", controllers.DeleteNote)

    r.Run(":8080")
}