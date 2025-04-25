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
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer(cfg *config.Config) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	rabbitPublisher, err := rabbitmq.NewPublisher(cfg.Rabbitmq.Rabbitmq_host, cfg.Rabbitmq.Rabbitmq_port, cfg.Rabbitmq.Rabbitmq_user, cfg.Rabbitmq.Rabbitmq_pass, cfg.Rabbitmq.Rabbitmq_queue)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ publisher: %v", err)
	}
	defer rabbitPublisher.Close()
	productRepo := repo.NewMongoProductRepository()
	productService := services.NewProductService(productRepo, rabbitPublisher)
	controller := controllers.NewProductController(productService)

	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	proto.RegisterProductServiceServer(grpcServer, controller)

	fmt.Println("gRPC Server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
