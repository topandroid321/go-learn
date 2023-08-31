package main

import (
	ApiRoutes "go-learning/src/Api"
	"go-learning/src/Client"
	"go-learning/src/Interfaces"
)

func main() {
	ApiRoutes.InitRoutes()
	Client.RunServer(Interfaces.SystemInterface{
		CorsAllowOrigin: "http://localhost",
		CorsAllowHeader: "*",
		Port:            3000,
	})
}
