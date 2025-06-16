package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	database "github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/3-REST-API-For-Notes-App/DB_Dir"
	"github.com/gin-gonic/gin"
)

func UpdateNoteById(c *gin.Context) {
	id := c.Param("id")
	var note database.Note

	noteId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = database.DB.QueryRow("SELECT id, title, content, created_at, updated_at FROM notes WHERE id = ?", noteId).
		Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	err = c.BindJSON(&note)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid note update"})
		return
	}
	statment, err := database.DB.Prepare(`
	UPDATE notes 
	SET title = ?, content = ?, updated_at = CURRENT_TIMESTAMP 
	WHERE id = ?
	`)
	defer statment.Close()

	_, err = statment.Exec(note.Title, note.Content, noteId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":      noteId,
		"title":   note.Title,
		"content": note.Content,
	})
}
