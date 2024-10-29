package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/pranay999000/smart-inventory/product-service/internal/database"
	"github.com/pranay999000/smart-inventory/product-service/internal/domain"
	"github.com/pranay999000/smart-inventory/product-service/internal/repository"
	"github.com/pranay999000/smart-inventory/product-service/internal/service"
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

	database.ConnectWriteDB()
	database.ConnectReadDB()

	domain.Migrate()

	productRepo := repository.NewProductRepo(database.ReadDB, database.WriteDB)

	userConn, err := grpc.NewClient("localhost" + os.Getenv("USER_SERVER_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}

	userGRPCClient := userproto.NewUserServiceClient(userConn)

	productService := service.NewProductService(productRepo, &userGRPCClient)

	var opts []grpc.ServerOption
	srv := grpc.NewServer(opts...)

	productproto.RegisterProductServiceServer(srv, productService)

	lis, err := net.Listen("tcp", os.Getenv("PRODUCT_SERVER_ADDR"))
	if err != nil {
		log.Fatalf("Failed to listen to %s: %v", os.Getenv("PRODUCT_SERVER_ADDR"), err)
	}

	log.Printf("Product Service running in port %s", os.Getenv("PRODUCT_SERVER_ADDR"))

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}