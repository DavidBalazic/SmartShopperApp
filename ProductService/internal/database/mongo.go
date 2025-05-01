package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/DavidBalazic/SmartShopperApp/config"
)

var (
	client *mongo.Client
	db     *mongo.Database
)

func Connect(cfg *config.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(cfg.DB.URL))
	if err != nil {
		return err
	}

	db = client.Database("smart_shopper")
	log.Println("MongoDB connected")
	return nil
}

func GetDB() *mongo.Database {
	if db == nil {
		log.Fatal("MongoDB not initialized.")
	}
	return db
}

func Disconnect() {
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("Error disconnecting MongoDB: %v", err)
		} else {
			log.Println("MongoDB disconnected")
		}
	}
}
