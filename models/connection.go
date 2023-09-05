package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDataBase() *gorm.DB {
  err := godotenv.Load(".env"); if err != nil {
	  log.Fatalf("Error loading .env file")
    panic(err)
	}	

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")
  DbSsl := os.Getenv("DB_SSL_MODE")
  DbTimeZone := os.Getenv("DB_TIME_ZONE")

	DBURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", 
		DbHost, DbUser, DbPassword, DbName, DbPort, DbSsl, DbTimeZone)
	
	DB, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{}); if err != nil {
		fmt.Println("Cannot connect to database ", err.Error())
		log.Fatal("connection error:", err)
	}
	return DB
}