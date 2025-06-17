package handlers

import (
	"net/http"
	"strconv"

	database "github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/3-REST-API-For-Notes-App/DB_Dir"
	"github.com/gin-gonic/gin"
)

func DeleteNoteById(c *gin.Context) {
	id := c.Param("id")

	noteId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	statement, err := database.DB.Prepare("DELETE FROM notes Where id= ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "incorrect id"})
		return
	}
	defer statement.Close()
	_, err = statement.Exec(noteId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "failed to excute querrty"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"deleted": "true"})
}
