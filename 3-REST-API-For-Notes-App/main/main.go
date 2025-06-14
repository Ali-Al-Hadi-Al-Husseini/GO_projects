package restapifornotesapp

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.POST("note", func(c *gin.Context) {

	})
}
