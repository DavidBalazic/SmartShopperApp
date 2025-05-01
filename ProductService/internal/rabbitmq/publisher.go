package rabbitmq

import (
	"encoding/json"
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher interface {
	PublishSingleProduct(message map[string]interface{}) error
	PublishMultipleProducts(messages []map[string]interface{}) error
}

type RabbitMQPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

// NewPublisher creates a new Publisher instance, connects to RabbitMQ, and declares a queue
func NewPublisher(rabbitmqHost, rabbitmqPort, rabbitmqUser, rabbitmqPass, RabbitmqQueue string) (*RabbitMQPublisher, error) {
	// Establish connection to RabbitMQ
	conn, err := amqp.Dial("amqp://" + rabbitmqUser + ":" + rabbitmqPass + "@" + rabbitmqHost + ":" + rabbitmqPort + "/")
	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %v", err)
		return nil, err
	}

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Failed to open a channel: %v", err)
		return nil, err
	}

	// Declare a queue
	queue, err := ch.QueueDeclare(
		RabbitmqQueue, // Queue name
		true,      // Durable: survives server restart
		false,     // Auto-delete: deletes when no consumers are connected
		false,     // Exclusive: used by only this connection
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		log.Printf("Failed to declare a queue: %v", err)
		return nil, err
	}

	return &RabbitMQPublisher{
		conn:    conn,
		channel: ch,
		queue:   queue,
	}, nil
}

// Publish sends a message to the RabbitMQ queue
func (p *RabbitMQPublisher) PublishSingleProduct(message map[string]interface{}) error {
	// Convert the message to a byte slice (example: JSON encoding)
	// Here we're assuming the message is a map and you need to serialize it as JSON
	body, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err)
		return err
	}

	// Publish the message to the queue
	err = p.channel.Publish(
		"",        // Default exchange
		p.queue.Name, // Queue name
		false,    // Mandatory: delivery fails if no queues are bound
		false,    // Immediate: ensure immediate delivery
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}

	log.Printf("Message published: %v", message)
	return nil
}

func (p *RabbitMQPublisher) PublishMultipleProducts(message []map[string]interface{}) error {
	// Convert the message to a byte slice (example: JSON encoding)
	// Here we're assuming the message is a map and you need to serialize it as JSON
	body, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err)
		return err
	}

	// Publish the message to the queue
	err = p.channel.Publish(
		"",        // Default exchange
		p.queue.Name, // Queue name
		false,    // Mandatory: delivery fails if no queues are bound
		false,    // Immediate: ensure immediate delivery
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}

	log.Printf("Message published: %v", message)
	return nil
}

func (p *RabbitMQPublisher) Close() {
	p.channel.Close()
	p.conn.Close()
}