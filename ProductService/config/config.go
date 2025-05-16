package config

import (
	"os"
	"strings"
	
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DB    MongoConfig
	Rabbitmq  RabbitMQConfig
	Kafka KafkaConfig
}

type RabbitMQConfig struct {
	Rabbitmq_queue string
	Rabbitmq_host string
	Rabbitmq_port string
	Rabbitmq_user string
	Rabbitmq_pass string	
}

type MongoConfig struct {
	URL      string
}

type KafkaConfig struct {
	Brokers []string
	Topic  string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Rabbitmq: RabbitMQConfig{
			Rabbitmq_queue: os.Getenv("RABBITMQ_QUEUE"),
			Rabbitmq_host: os.Getenv("RABBITMQ_HOST"),
			Rabbitmq_port: os.Getenv("RABBITMQ_PORT"),
			Rabbitmq_user: os.Getenv("RABBITMQ_USER"),
			Rabbitmq_pass: os.Getenv("RABBITMQ_PASSWORD"),
		},
		DB: MongoConfig{
			URL: os.Getenv("MONGO_URI"),
		},
		Kafka: KafkaConfig{
			Brokers: strings.Split(os.Getenv("KAFKA_BROKER"), ","),
			Topic:   os.Getenv("KAFKA_TOPIC"),
		},
	}

	return cfg, nil
}