package config

import (
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	// "go.uber.org/zap"
)

var mrouter *melody.Melody
var router *gin.Engine

// var logger *zap.Logger

// initialize melody and gin routers
func Init() {
	mrouter = melody.New()
	router = gin.Default()
	// var err error
	// logger, err = zap.NewDevelopment()
	// if err != nil {
	// 	log.Fatal("error starting zap logger")
	// }

}

func GetWebSocketRouter() *melody.Melody {
	return mrouter
}

func GetHTTPRouter() *gin.Engine {
	return router
}

// func GetLogger() *zap.Logger {
// 	return logger
// }
