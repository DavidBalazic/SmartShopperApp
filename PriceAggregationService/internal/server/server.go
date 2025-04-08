package server

import (
	"fmt"
	"log"
	"net"

	"github.com/DavidBalazic/SmartShopperApp/config"
	"github.com/DavidBalazic/SmartShopperApp/internal/controllers"
	"github.com/DavidBalazic/SmartShopperApp/internal/proto"
	"github.com/DavidBalazic/SmartShopperApp/internal/rabbitmq"
	"github.com/DavidBalazic/SmartShopperApp/internal/repo"
	"github.com/DavidBalazic/SmartShopperApp/internal/services"
	"google.golang.org/grpc"
)

func StartGRPCServer(cfg *config.Config) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	productRepo := repo.NewMongoProductRepository()
	productService := services.NewProductService(productRepo)
	rabbitPublisher, err := rabbitmq.NewPublisher(cfg.Rabbitmq.Rabbitmq_host, cfg.Rabbitmq.Rabbitmq_queue)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ publisher: %v", err)
	}
	defer rabbitPublisher.Close()
	controller := controllers.NewProductController(productService, rabbitPublisher)

	grpcServer := grpc.NewServer()

	proto.RegisterProductServiceServer(grpcServer, controller)

	fmt.Println("gRPC Server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
