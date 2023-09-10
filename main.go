package main

import (
	"github.com/JoaoVFerreira/go-scrapper/db"
	"github.com/JoaoVFerreira/go-scrapper/handler"
	"github.com/JoaoVFerreira/go-scrapper/output"
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
