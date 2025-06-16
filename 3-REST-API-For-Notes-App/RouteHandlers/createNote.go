package handlers

import (
	"net/http"

	database "github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/3-REST-API-For-Notes-App/DB_Dir"
	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {
	var note database.Note

	err := c.ShouldBindJSON(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	statment, err := database.DB.Prepare("INSERT INTO notes (title, content) VALUES (?,?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer statment.Close()

	res, err := statment.Exec(note.Title, note.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, _ := res.LastInsertId()

	c.JSON(http.StatusCreated, gin.H{
		"id":      id,
		"title":   note.Title,
		"content": note.Content,
	})

}
