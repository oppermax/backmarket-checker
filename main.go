package main

import (
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

var url = "https://www.backmarket.de/ipad-air-3-gebraucht.html"

func main()  {

	type price struct {
		Price string
		Link string
	}

	var prices []price
	c := colly.NewCollector(
		colly.AllowedDomains("www.backmarket.de"), )

	c.OnHTML(".price", func(e *colly.HTMLElement) {
		p := price{
			Price: e.Text,
			Link:  "wip",
		}
		prices = append(prices, p)
		log.Info(p)
	})


	c.OnError(func(_ *colly.Response, err error) {
		log.Info("Something went wrong", err)
	})
	

	c.OnRequest(func(request *colly.Request) {
		log.Infof("Visiting %v", request.URL)
	})

	err := c.Visit(url)
	if err != nil {
		log.WithError(err)
	}
	c.Wait()



}