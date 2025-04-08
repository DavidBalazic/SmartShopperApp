package config

import (
	"context"
	"log"
	"os"
	"sync"
	"time"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	DB    MongoConfig
	Rabbitmq  RabbitMQConfig
}

type RabbitMQConfig struct {
	Rabbitmq_queue string
	Rabbitmq_host string
}

type MongoConfig struct {
	URL      string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Rabbitmq: RabbitMQConfig{
			Rabbitmq_queue: os.Getenv("RABBITMQ_QUEUE"),
			Rabbitmq_host: os.Getenv("RABBITMQ_HOST"),
		},
		DB: MongoConfig{
			URL: os.Getenv("MONGO_URI"),
		},
	}

	return cfg, nil
}

var (
	db     *mongo.Database
	client *mongo.Client
	once   sync.Once
)

func ConnectDB() {
	once.Do(func() {
		// load := godotenv.Load()
		// if load != nil {
		// 	log.Fatalf("Error loading .env file")
		//   }
		  
		uri := os.Getenv("MONGO_URI")
		if uri == "" {
			uri = "mongodb://localhost:27017"
		}
		log.Printf("Using MongoDB URI: %s", uri)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var err error
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatalf("MongoDB Connection Error: %v", err)
		}

		db = client.Database("smart_shopper")
	})
}

func GetDB() *mongo.Database {
	if db == nil {
		log.Fatal("Database not initialized. Call ConnectDB() first.")
	}
	return db
}

func DisconnectDB() {
	if client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := client.Disconnect(ctx)
		if err != nil {
			log.Printf("Error disconnecting MongoDB: %v", err)
		}
	}
}