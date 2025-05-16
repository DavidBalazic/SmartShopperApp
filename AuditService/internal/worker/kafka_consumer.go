package worker

import (
	"AuditService/config"
	"AuditService/internal/models"
	"AuditService/internal/services"
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartKafkaConsumer(svc services.AuditService, cfg *config.Config) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.Kafka.KAFKA_BROKER},
		Topic:   cfg.Kafka.KAFKA_TOPIC,
		GroupID: cfg.Kafka.KAFKA_GROUP,
	})

	ctx := context.Background()

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("Kafka read error: %v", err)
			continue
		}

		log.Printf("Received Kafka message: %s", string(msg.Value))

		var logEntry models.AuditLog
		if err := json.Unmarshal(msg.Value, &logEntry); err != nil {
			log.Printf("JSON parse error: %v", err)
			continue
		}

		if err := svc.HandleAuditLog(logEntry); err != nil {
			log.Printf("Failed to save audit log: %v", err)
		}
	}
}
