package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteNoteById(c *gin.Context) {
	id := c.Param("id")

	noteId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
}
