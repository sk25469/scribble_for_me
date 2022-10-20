package routes

import (
	"github.com/olahol/melody"
	"github.com/sk25469/scribble_backend/pkg/controllers"
)

func RegisterSocketRoutes(mrouter *melody.Melody) {
	mrouter.HandleConnect(controllers.OnConnect)
	mrouter.HandleDisconnect(controllers.OnDisconnect)
	mrouter.HandleMessage(controllers.OnMessage)
}
