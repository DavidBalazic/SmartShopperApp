package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func ConnectDB() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("MongoDB Client Error: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("MongoDB Connection Error: %v", err)
	}

	db = client.Database("smart_shopper")
	fmt.Println("Connected to MongoDB")
}

func GetDB() *mongo.Database {
	if db == nil {
		log.Fatal("Database not initialized. Call ConnectDB() first.")
	}
	return db
}