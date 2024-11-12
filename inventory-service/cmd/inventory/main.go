package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/pranay999000/smart-inventory/inventory-service/internal/database"
	"github.com/pranay999000/smart-inventory/inventory-service/internal/domain"
	"github.com/pranay999000/smart-inventory/inventory-service/internal/repository"
	"github.com/pranay999000/smart-inventory/inventory-service/internal/service"
	inventoryproto "github.com/pranay999000/smart-inventory/inventory-service/proto/inventory"
	productproto "github.com/pranay999000/smart-inventory/product-service/proto/product"
	userproto "github.com/pranay999000/smart-inventory/user-service/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env: %v\n", err)
	}

	database.ConnectReadDB()
	database.ConnectWriteDB()

	domain.InventoryMigrate()

	userConn, err := grpc.NewClient("localhost" + os.Getenv("USER_SERVER_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}

	productConn, err := grpc.NewClient("localhost" + os.Getenv("PRODUCT_SERVER_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}

	userGRPCClient := userproto.NewUserServiceClient(userConn)
	productGRPCClient := productproto.NewProductServiceClient(productConn)

	inventoryRepo := repository.NewInventoryRepo(database.ReadDB, database.WriteDB)

	inventoryService := service.NewInventoryServiceServer(inventoryRepo, &userGRPCClient, &productGRPCClient)

	var opts []grpc.ServerOption
	srv := grpc.NewServer(opts...)

	inventoryproto.RegisterInventoryServiceServer(srv, inventoryService)

	lis, err := net.Listen("tcp", os.Getenv("INVENTORY_SERVER_ADDR"))
	if err != nil {
		log.Fatalf("Failed to listen to %s: %v", os.Getenv("INVENTORY_SERVER_ADDR"), err)
	}

	log.Printf("Inventory Service running on port %s", os.Getenv("INVENTORY_SERVER_ADDR"))

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}