package main

import (
	"log"
	"github.com/DavidBalazic/SmartShopperApp/config"
	"github.com/DavidBalazic/SmartShopperApp/internal/server"
	"github.com/DavidBalazic/SmartShopperApp/internal/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := db.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer db.Disconnect()

	server.StartGRPCServer(cfg)
}