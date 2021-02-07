package main

import (
	"os"

	"github.com/nicobianchetti/Meli-Quasar-NB/src/router"
)

const (
	_portDefualt = "5000"
	_ipDefault   = ""
)

func main() {
	httpRouter := router.NewMuxRouter()
	router.SetupRoutesSatellite(httpRouter)

	serverPort := os.Getenv("PORT")

	if serverPort == "" {
		serverPort = _portDefualt
	}

	serverIP := os.Getenv("IP")

	if serverIP == "" {
		serverIP = _ipDefault
	}

	addr := serverIP + ":" + serverPort

	httpRouter.SERVE(addr)
}
