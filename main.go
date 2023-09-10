package main

import (
	"github.com/JoaoVFerreira/GOScrapper/db"
	"github.com/JoaoVFerreira/GOScrapper/handler"
)


func main() {
	// Initialize database
	db.ConnectToDataBase()
	// Scrap real estate code
	handler.ScrapRealEstate()
	// Scrap real estate details by code
	handler.ScrapRealEstateDetail()
} 
