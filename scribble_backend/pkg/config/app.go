package config

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

var mrouter *melody.Melody
var router *gin.Engine

// initialize melody and gin routers
func Init() {
	mrouter = melody.New()
	router = gin.Default()
}

func GetWebSocketRouter() *melody.Melody {
	return mrouter
}

func GetHTTPRouter() *gin.Engine {
	return router
}
