package main

import (
	"github.com/JoaoVFerreira/GOScrapper/db"
	"github.com/JoaoVFerreira/GOScrapper/output"
)


func main() {
	// Initialize database
	db.ConnectToDataBase()
	// Scrap real estate code
	// handler.ScrapRealEstate()
	// Scrap real estate details by code
	// handler.ScrapRealEstateDetail()
	output.Decision()
} 
