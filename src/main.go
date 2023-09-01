package main

import (
	"go-learning/src/Client"
	"go-learning/src/Interfaces"
)

func main() {
	Client.RunServer(Interfaces.SystemInterface{
		CorsAllowOrigin: "*",
		CorsAllowHeader: "*",
		Port:            3000,
	})
}
