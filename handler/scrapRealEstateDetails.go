package handler

import (
	"log"
	"strings"

	"github.com/JoaoVFerreira/GOScrapper/db"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

const completedParagraphTexts = 26;

func ScrapRealEstateDetail() {
	c1 := colly.NewCollector()
	c1.IgnoreRobotsTxt = true
	c1.UserAgent = userAgent
	
	f := db.RealStateFund{}
	funds, err := f.GetAll(); if err != nil {
		log.Println("Error retrieving funds from the database: ", err)
		panic(err)
	}
	var pTexts []string
	c1.OnHTML(".indicators__box", func(h *colly.HTMLElement) {
			h.DOM.Find("p").Each(func(_ int, p *goquery.Selection) {
				pText := p.Text()
				pTexts = append(pTexts, pText)
			})

			if len(pTexts) == completedParagraphTexts {
				fund := db.RealStateFund{}
				liquidity := pTexts[1]
				dy := pTexts[5]
				pvp := pTexts[15] 

				url := h.Request.URL.String()
				urlSplit := strings.Split(url, "/")
				code := urlSplit[len(urlSplit)-1]

				fund.Pvp = pvp
				fund.Liquidity = liquidity
				fund.DividendYield = dy
				fund.Code = code
				fund.UpdateFund(dy, pvp, liquidity)
			}
	})

	c1.OnError(func(response *colly.Response, err error) {
    log.Println("Request URL:", response.Request.URL, "\nError:", err)
	})

	for _, fund := range funds {
		url := baseURL + fund.Code
		c1.Visit(url)

		pTexts = []string{}
	}
	c1.Visit(baseURL)
}