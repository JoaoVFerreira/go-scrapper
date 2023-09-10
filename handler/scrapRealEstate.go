package handler

import (
	"github.com/JoaoVFerreira/GOScrapper/db"
	"github.com/gocolly/colly/v2"
)


const baseURL = "https://www.fundsexplorer.com.br/funds/"
const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"

func ScrapRealEstate() {
	c1 := colly.NewCollector()
	c1.IgnoreRobotsTxt = true
	c1.UserAgent = userAgent

	c1.OnHTML(".tickerBox", func(h *colly.HTMLElement) {
    desc := h.DOM.Find(".tickerBox__desc").Text()
    ticker := h.DOM.Find(".tickerBox__title").Text()
		title := h.DOM.Find(".tickerBox__type").Text()

		fund := db.RealStateFund{}
		fund.Code = ticker
		fund.Type = title
		fund.Description = desc
		fund.SaveFund()
	})
	c1.Visit(baseURL)
}