package handlers

import (
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
	}

}
