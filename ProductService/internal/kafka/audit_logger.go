package kafka

import (
    "context"
    "encoding/json"
    "log"
    "time"

    "github.com/segmentio/kafka-go"
    "github.com/DavidBalazic/SmartShopperApp/internal/dtos"
)

type AuditLogger interface {
    PublishAuditLog(ctx context.Context, logEntry dtos.AuditLog) error
    Close() error
}

type KafkaPublisher struct {
    writer *kafka.Writer
}

func NewKafkaPublisher(brokers []string, topic string) (*KafkaPublisher, error) {
    return &KafkaPublisher{
        writer: &kafka.Writer{
            Addr:     kafka.TCP(brokers...),
            Topic:    topic,
            Balancer: &kafka.LeastBytes{},
        },
    }, nil
}

func (p *KafkaPublisher) PublishAuditLog(ctx context.Context, logEntry dtos.AuditLog) error {
    data, err := json.Marshal(logEntry)
    if err != nil {
        return err
    }

    actorID := logEntry.Actor.ID

    msg := kafka.Message{
        Key:   []byte(actorID),
        Value: data,
        Time:  time.Now(),
    }

    err = p.writer.WriteMessages(ctx, msg)
    if err != nil {
        log.Printf("Kafka publish error: %v", err)
        return err
    }
    return nil
}

func (p *KafkaPublisher) Close() error {
    return p.writer.Close()
}