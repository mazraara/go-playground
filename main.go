package main

import (
    "notes-app/controllers"
    "github.com/gin-gonic/gin"
    "html/template"
)

func truncate(s string, length int) string {
    if len(s) <= length {
        return s
    }
    return s[:length] + "..."
}

func main() {
    r := gin.Default()

    templ := template.Must(template.New("").Funcs(template.FuncMap{
        "truncate": truncate,
    }).ParseGlob("templates/*"))

    r.SetHTMLTemplate(templ)

    r.GET("/", controllers.ListNotes)
    r.GET("/add", controllers.ShowAddNote)
    r.POST("/add", controllers.AddNote)
    r.GET("/edit/:id", controllers.ShowEditNote)
    r.POST("/edit/:id", controllers.UpdateNote)
    r.GET("/delete/:id", controllers.DeleteNote)

    r.Run(":8080")
}