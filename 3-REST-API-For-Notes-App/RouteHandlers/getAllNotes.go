package handlers

import (
	"net/http"

	database "github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/3-REST-API-For-Notes-App/DB_Dir"
	"github.com/gin-gonic/gin"
)

func GetAllNotes(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, title, content, created_at, updated_at from notes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to query notes"})
	}

	defer rows.Close()

	var notes []database.Note

	for rows.Next() {
		var note database.Note

		err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to scan notes"})
		}
		notes = append(notes, note)
	}
	c.JSON(http.StatusOK, notes)
}
