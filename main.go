package main

import (
	"fmt"
	"log"

	"github.com/JoaoVFerreira/GOScrapper/db"
	"github.com/gocolly/colly/v2"
)

func main() {
	db.ConnectToDataBase()
	c := colly.NewCollector()
	c.IgnoreRobotsTxt = true
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"

	c.OnHTML(".tickerBox", func(h *colly.HTMLElement) {
    desc := h.DOM.Find(".tickerBox__desc").Text()
    ticker := h.DOM.Find(".tickerBox__title").Text()
		title := h.DOM.Find(".tickerBox__type").Text()

		fund := db.RealStateFund{}
		fund.Code = ticker
		fund.Type = title
		fund.Description = desc
		fund.SaveFund()
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Connection", "keep-alive")
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, "scraped!")
	})

	c.Visit("https://www.fundsexplorer.com.br/funds")
} 
