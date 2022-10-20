package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sk25469/scribble_backend/pkg/config"
	"github.com/sk25469/scribble_backend/pkg/controllers"
)

func RegisterAllRoutes() {

	mrouter := config.GetWebSocketRouter()
	router := config.GetHTTPRouter()
	// main html page
	router.GET("/", controllers.GetHomePageRoute)

	// convertes http requests in a ws
	router.GET("/ws", func(c *gin.Context) {
		mrouter.HandleRequest(c.Writer, c.Request)
	})

	RegisterSocketRoutes(mrouter)

	router.Run(":5000")
}
