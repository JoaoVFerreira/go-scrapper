package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type RealStateFund struct {
	gorm.Model
	ID   uint  `gorm:"primaryKey" json:"id"`
	Code string `gorm:"size:255;not null;unique;index" json:"code"`
	Type string `gorm:"size:255;not null;" json:"type"`
	Description string `gorm:"size:255;not null;" json:"description"`
	CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (f * RealStateFund) SaveFund() (*RealStateFund, error) {
	err := DB.Create(&f).Error; if err != nil {
		return &RealStateFund{}, err
	}
	return f, nil
}

func ConnectToDataBase() {
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
	DB.AutoMigrate(&RealStateFund{})
}