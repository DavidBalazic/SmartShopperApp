package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/DavidBalazic/SmartShopperApp/internal/controllers"
	"github.com/DavidBalazic/SmartShopperApp/internal/proto"
	"github.com/DavidBalazic/SmartShopperApp/internal/services"
	"github.com/DavidBalazic/SmartShopperApp/internal/repo"
)

func StartGRPCServer() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	productRepo := repo.NewMongoProductRepository()
	productService := services.NewProductService(productRepo)
	controller := controllers.NewProductController(productService)

	grpcServer := grpc.NewServer()

	proto.RegisterProductServiceServer(grpcServer, controller)

	fmt.Println("gRPC Server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
