package db

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type RealStateFund struct {
	gorm.Model
	ID           	uuid.UUID 			`gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Code         	string    			`gorm:"size:255;not null;unique;index" json:"code"`
	Type         	string    			`gorm:"size:255;not null;" json:"type"`
	Description  	string    			`gorm:"size:255;not null;" json:"description"`
	DividendYield string   				`gorm:"size:255;" json:"dividend_yield"`
	Pvp          	string    			`gorm:"size:255;" json:"pvp"`
	Liquidity    	string    			`gorm:"size:255;" json:"liquidity"`
	CreatedAt    	time.Time 			`json:"created_at"`
	UpdatedAt    	time.Time 			`json:"updated_at"`
	DeletedAt    	gorm.DeletedAt  `json:"deleted_at,omitempty"`
}

func (f *RealStateFund) SaveFund() (*RealStateFund, error) {
	err := DB.Create(&f).Error; if err != nil {
		return &RealStateFund{}, err
	}
	return f, nil
}

func (f *RealStateFund) BeforeSave(tx *gorm.DB) error {
	if f.Code == "" || f.Type == "" || f.Description == "" {
		return errors.New("All fields must be filled in.")
	}
	
	if strings.Contains(f.Type, "Indefinido") {
		return errors.New("Field 'Type' is not allowed to contain 'Indefinido'.")
	}

	bdr := "B"
	if strings.HasSuffix(f.Code, bdr) {
		return errors.New("Field 'Code' is not allowed to be a brazilian depositary receipt 'B'.")
	}
	
	return nil
}

func (f *RealStateFund) GetAll() ([]*RealStateFund, error) {
	funds := []*RealStateFund{}
	err := DB.Find(&funds).Error; if err != nil {
		return nil, err
	} 
	return funds, nil
}

func (f *RealStateFund) UpdateFund(dy, pvp, liquidity string) (error) {
	err := DB.Model(f).Where("code = ?", f.Code).Updates(map[string]interface{}{
		"dividend_yield": dy,
		"pvp":            pvp,
		"liquidity":      liquidity,
	}).Error; if err != nil {
		return err
	}
	return nil
}

func (f *RealStateFund) FindByCode(code string) (*RealStateFund, error) {
	var fund RealStateFund
	err := DB.Where("code = ?", code).First(&fund).Error
	if err != nil {
		return nil, err 
	}
	return &fund, nil
}

func (f *RealStateFund) BeforeUpdate(tx *gorm.DB) error {
	if f.DividendYield == "" || f.Pvp == "" || f.Liquidity == "" {
		return errors.New("All fields must be filled in.")
	}

	dy := strings.ReplaceAll(f.DividendYield, " ", "")
	pvp := strings.ReplaceAll(f.Pvp, " ", "")
	l := strings.ReplaceAll(f.Liquidity, " ", "")

	if strings.Contains(dy, "N/A") || strings.Contains(pvp, "N/A" ) || strings.Contains(l, "N/A" ) {
		return errors.New("Field is not allowed to contain 'N/A'.")
	}

	f.DividendYield = dy
	f.Liquidity = l
	f.Pvp = pvp
	return nil
}

func (f *RealStateFund) FindAllWithData() ([]RealStateFund, error) {
	var funds []RealStateFund
	err := DB.Where("pvp IS NOT NULL AND liquidity IS NOT NULL AND dividend_yield IS NOT NULL").Find(&funds).Error
	if err != nil {
		return nil, err 
	}
	return funds, nil 
}

func ConnectToDataBase() {
  err := godotenv.Load(".env"); if err != nil {
	  log.Fatalf("Error loading .env file")
    panic(err)
	}	

	DbHost      := os.Getenv("DB_HOST")
	DbUser      := os.Getenv("DB_USER")
	DbPassword  := os.Getenv("DB_PASSWORD")
	DbName      := os.Getenv("DB_NAME")
	DbPort      := os.Getenv("DB_PORT")
	DbSsl       := os.Getenv("DB_SSL_MODE")
	DbTimeZone  := os.Getenv("DB_TIME_ZONE")
	
	DBURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", 
		DbHost, DbUser, DbPassword, DbName, DbPort, DbSsl, DbTimeZone)
	
	DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{}); if err != nil {
		fmt.Println("Cannot connect to database ", err.Error())
		log.Fatal("connection error:", err)
	}
	DB.AutoMigrate(&RealStateFund{})
}