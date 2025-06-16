package router

import (
	createNote "github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/3-REST-API-For-Notes-App/RouteHandlers"
	"github.com/gin-gonic/gin"
)

func runRouter() {
	router := gin.Default()
	router.POST("note", func(c *gin.Context) {
		createNote.CreateNote(c)
	})
	router.Run()
}
