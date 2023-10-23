package main

import (
	"crawler-financial/models"
	"crawler-financial/utils"
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

var (
	ListFuturesActifs []models.FuturesActifs
)

func main() {
	c := colly.NewCollector(
		colly.Async(true),
		colly.MaxDepth(4),
	)
	sites := utils.ReadCSV()
	c.SetRequestTimeout(120 * time.Second)
	c.OnHTML("#list-res-table", func(h *colly.HTMLElement) {
		h.ForEach("tr", func(i int, h *colly.HTMLElement) {
			future := models.NewFuture()
			h.ForEach("td", func(i int, h *colly.HTMLElement) {
				switch {
				case i == 0:
					future.Symbol = h.Text
				case i == 1:
					future.Name = h.Text
				case i == 2:
					future.LastPrice = models.ParserStringToFloat64(h.Text)

				case i == 3:
					future.MarketTime = h.Text
				case i == 4:
					future.Change = h.Text
				case i == 5:
					future.PercentChange = h.Text
				case i == 6:
					// future.Volume = models.ParserStringToFloat64(h.Text)
					future.Volume = h.Text
				case i == 7:
					future.TotalValue = h.Text
				}
			})

			ListFuturesActifs = append(ListFuturesActifs, *future)
		})
		fmt.Println(ListFuturesActifs)

	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visita register", r.StatusCode)
	})
	fmt.Println(sites)
	for _, site := range sites {
		c.Visit(site)
		c.Wait()
	}
}
