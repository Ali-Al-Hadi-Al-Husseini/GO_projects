package router

import (
	handlers "github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/3-REST-API-For-Notes-App/RouteHandlers"
	"github.com/gin-gonic/gin"
)

func RunRouter() {
	router := gin.Default()
	router.POST("note", handlers.CreateNote)
	router.GET("notes", handlers.GetAllNotes)
	router.GET("note/:id", handlers.GetlNoteByID)
	router.Run()
}
