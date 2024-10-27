package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

var ReadDB *sql.DB
var WriteDB *sql.DB

func ConnectReadDB() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s?parseTime=true", "root", os.Getenv("MYSQL_MASTER_ROOT_PASSWORD"), 3300, os.Getenv("MYSQL_DATABASE"))

	ReadDB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to Read MySQL: %v\n", err)
	}

	ReadDB.SetMaxOpenConns(25)
	ReadDB.SetMaxIdleConns(25)
	ReadDB.SetConnMaxLifetime(time.Minute)

	if err = ReadDB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v\n", err)
	}

	fmt.Println("Read MySQL Connected")
}

func CloseReadDB() {
	if err := ReadDB; err != nil {
		log.Fatalf("Failed to connect to Read MYSQL: %v\n", err)
	}
}

func ConnectWriteDB() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(localhost:%d)/%s?parseTime=true", "root", os.Getenv("MYSQL_MASTER_ROOT_PASSWORD"), 3306, os.Getenv("MYSQL_DATABASE"))

	WriteDB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to Write MySQL: %v\n", err)
	}

	WriteDB.SetMaxOpenConns(25)
	WriteDB.SetMaxIdleConns(25)
	WriteDB.SetConnMaxLifetime(time.Minute)

	if err = WriteDB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v\n", err)
	}

	fmt.Println("Write MySQL Connected")
}

func CloseWriteDB() {
	if err := WriteDB; err != nil {
		log.Fatalf("Failed to connect to Write MYSQL: %v\n", err)
	}
}