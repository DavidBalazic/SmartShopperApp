package main

import (
	"log"
	"context"
	"time"

	"AuditService/config"
	"AuditService/internal/repository"
	"AuditService/internal/services"
	"AuditService/internal/worker"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.DB.URL))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	collection := client.Database("auditlogdb").Collection("auditlogs")

	repo := repository.NewMongoRepository(collection)
	auditService := services.NewAuditService(repo)

	worker.StartKafkaConsumer(auditService, cfg)
}