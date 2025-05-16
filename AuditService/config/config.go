package config

import (
	"os"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DB    MongoConfig
	Kafka  KafkaConfig
}

type KafkaConfig struct {
	KAFKA_TOPIC string
	KAFKA_BROKER string
	KAFKA_GROUP string
}

type MongoConfig struct {
	URL      string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Kafka: KafkaConfig{
			KAFKA_TOPIC: os.Getenv("KAFKA_TOPIC"),
			KAFKA_BROKER: os.Getenv("KAFKA_BROKER"),
			KAFKA_GROUP: os.Getenv("KAFKA_GROUP"),
		},
		DB: MongoConfig{
			URL: os.Getenv("MONGO_URI"),
		},
	}

	return cfg, nil
}