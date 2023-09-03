package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()
	c.IgnoreRobotsTxt = true
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"

	c.OnHTML("", func(h *colly.HTMLElement) {
		fmt.Println("Texto do botão:", h.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Connection", "keep-alive")
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
		body := string(r.Body)
		fmt.Println("Body:", body)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, "scraped!")
	})

	c.Visit("https://www.fundsexplorer.com.br/")
} 
