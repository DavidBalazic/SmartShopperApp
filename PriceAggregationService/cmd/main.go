package main

import (
	"github.com/DavidBalazic/SmartShopperApp/config"
	"github.com/DavidBalazic/SmartShopperApp/internal/server"
)

func main() {
	config.ConnectDB()
	defer config.DisconnectDB()

	server.StartGRPCServer()
}