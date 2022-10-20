package main

import (
	"github.com/sk25469/scribble_backend/pkg/config"
	"github.com/sk25469/scribble_backend/pkg/routes"
)

func main() {
	config.Init()
	routes.RegisterAllRoutes()
}
