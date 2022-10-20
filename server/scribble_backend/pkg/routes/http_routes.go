package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"github.com/sk25469/scribble_backend/pkg/controllers"
)

func RegisterAllRoutes() {
	router := gin.Default()
	mrouter := melody.New()

	// main html page
	router.GET("/", controllers.GetHomePageRoute)

	// convertes http requests in a ws
	router.GET("/ws", func(c *gin.Context) {
		mrouter.HandleRequest(c.Writer, c.Request)
	})

	RegisterSocketRoutes(mrouter)

	router.Run(":5000")
}
