package main

import (
	"github.com/DavidBalazic/SmartShopperApp/config"
	"github.com/DavidBalazic/SmartShopperApp/server"
)

func main() {
	config.ConnectDB()
	server.StartGRPCServer()
}