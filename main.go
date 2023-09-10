package main

import (
	"github.com/JoaoVFerreira/GOScrapper/db"
	"github.com/JoaoVFerreira/GOScrapper/handler"
	"github.com/JoaoVFerreira/GOScrapper/output"
)


func main() {
	// Initialize database
	db.ConnectToDataBase()
	// Scrap real estate code
	handler.ScrapRealEstate()
	// Scrap real estate details by code
	handler.ScrapRealEstateDetail()
	// Choose best funds and write a json file with the result
	output.Decision()
} 
