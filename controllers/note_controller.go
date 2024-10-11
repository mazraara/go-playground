package controllers

import (
    "net/http"
    "strconv"
    "notes-app/models"
    "github.com/gin-gonic/gin"
)

func ListNotes(c *gin.Context) {
    notes := models.GetAllNotes()
    c.HTML(http.StatusOK, "list.html", gin.H{
        "notes": notes,
    })
}

func ShowAddNote(c *gin.Context) {
    c.HTML(http.StatusOK, "add.html", nil)
}

func AddNote(c *gin.Context) {
    content := c.PostForm("content")

    _, err := models.AddNote(content)
    if err != nil {
        c.HTML(http.StatusBadRequest, "add.html", gin.H{
            "error": err.Error(),
            "content": content, // Send back the content to preserve user input
        })
        return
    }

    c.Redirect(http.StatusSeeOther, "/")
}

func ShowEditNote(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    note, found := models.GetNote(id)
    if !found {
        c.AbortWithStatus(http.StatusNotFound)
        return
    }
    c.HTML(http.StatusOK, "edit.html", note)
}

func UpdateNote(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    content := c.PostForm("content")

    updated, err := models.UpdateNote(id, content)
    if err != nil {
        c.HTML(http.StatusBadRequest, "edit.html", gin.H{
            "error": err.Error(),
            "ID": id,
            "Content": content, // Send back the content to preserve user input
        })
        return
    }

    if !updated {
        c.AbortWithStatus(http.StatusNotFound)
        return
    }

    c.Redirect(http.StatusSeeOther, "/")
}

func DeleteNote(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    if models.DeleteNote(id) {
        c.Redirect(http.StatusSeeOther, "/")
    } else {
        c.AbortWithStatus(http.StatusNotFound)
    }
}