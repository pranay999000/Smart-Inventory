package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ReadDB *gorm.DB
var WriteDB *gorm.DB

func ConnectReadDB() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", os.Getenv("MYSQL_MASTER_ROOT_PASSWORD"), 3300, os.Getenv("MYSQL_DATABASE"))

	ReadDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Read MySQL: %v\n", err)
	}

	log.Println("Read MySQL Connected")
}

func CloseReadDB() {
	sqlDB, err := ReadDB.DB()
	if err != nil {
		log.Fatalf("Error retriving database instanace: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
}

func ConnectWriteDB() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", os.Getenv("MYSQL_MASTER_ROOT_PASSWORD"), 3306, os.Getenv("MYSQL_DATABASE"))

	WriteDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Write MySQL: %v\n", err)
	}

	log.Println("Write MySQL Connected")
}

func CloseWriteDB() {
	sqlDB, err := WriteDB.DB()
	if err != nil {
		log.Fatalf("Error retriving database instanace: %v", err)
	}

	if err := sqlDB.Close(); err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
}