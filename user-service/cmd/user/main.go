package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/pranay999000/smart-inventory/user-service/internal/database"
	"github.com/pranay999000/smart-inventory/user-service/internal/repository"
	"github.com/pranay999000/smart-inventory/user-service/internal/service"
	userproto "github.com/pranay999000/smart-inventory/user-service/proto"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env: %v\n", err)
	}

	mongoClient, err := database.ConnectToMongoDB(os.Getenv("MONGODB_URI"))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	userCollection := mongoClient.Database(os.Getenv("MONGO_DATABASE_NAME")).Collection("users")

	userRepo := repository.NewUserRepository(mongoClient, userCollection)
	userService := service.NewUserService(userRepo)

	var opts []grpc.ServerOption
	srv := grpc.NewServer(opts...)

	userproto.RegisterUserServiceServer(srv, userService)

	lis, err := net.Listen("tcp", os.Getenv("USER_SERVER_ADDR"))
	if err != nil {
		log.Fatalf("Failed to listen to %s: %v", os.Getenv("USER_SERVER_ADDR"), err)
	}

	log.Printf("User Service running on port %s", os.Getenv("USER_SERVER_ADDR"))

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}