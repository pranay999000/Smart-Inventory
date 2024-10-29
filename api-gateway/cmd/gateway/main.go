package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/pranay999000/smart-inventory/api-gateway/handler"
	"github.com/pranay999000/smart-inventory/api-gateway/middleware"
	productproto "github.com/pranay999000/smart-inventory/product-service/proto/product"
	businessproto "github.com/pranay999000/smart-inventory/user-service/proto/business"
	userproto "github.com/pranay999000/smart-inventory/user-service/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env: %v\n", err)
	}

	log.Printf("Connecting to UserService on Port: %s", os.Getenv("USER_SERVER_ADDR"))

	userConn, err := grpc.NewClient("localhost" + os.Getenv("USER_SERVER_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to user service: %v", err)
	}

	productConn, err := grpc.NewClient("localhost" + os.Getenv("PRODUCT_SERVER_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}

	defer userConn.Close()

	userGRPCClient := userproto.NewUserServiceClient(userConn)
	businessGRPCClient := businessproto.NewBusinessServiceClient(userConn)
	productGRPCClient := productproto.NewProductServiceClient(productConn)

	userHandler := handler.NewUserHandler(&userGRPCClient)
	businessHandler := handler.NewBusinessHandler(&businessGRPCClient)
	productHandler := handler.NewProductHandler(&productGRPCClient)

	http.HandleFunc("/api/v1/signup", middleware.LoggingMiddleware(userHandler.SignUpHandler))
	http.HandleFunc("/api/v1/signin", middleware.LoggingMiddleware(userHandler.SigninHandler))
	
	http.HandleFunc("/api/v1/business/create", middleware.LoggingMiddleware(middleware.AuthenticationMiddleware(businessHandler.CreateBusinessHandler)))

	http.HandleFunc("/api/v1/product/create", middleware.LoggingMiddleware(middleware.AuthenticationMiddleware(productHandler.CreateProductHandler)))

	log.Printf("API Gateway running in post: %d", 3000)
	log.Fatal(http.ListenAndServe(":3000", nil))

}