package main

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	log "github.com/sirupsen/logrus"
)

var url = "https://www.backmarket.de/"

func main()  {

	type price struct {
		Price string
		Link string
	}

	var prices []price
	c := colly.NewCollector(
		colly.AllowedDomains("www.backmarket.de"),
		colly.Debugger(&debug.LogDebugger{}))

	c.OnHTML("a", func(e *colly.HTMLElement) {
		p := price{
			Price: e.Text,
			Link:  "wip",
		}
		prices = append(prices, p)
	})


	c.OnError(func(_ *colly.Response, err error) {
		log.Info("Something went wrong", err)
	})
	

	c.OnRequest(func(request *colly.Request) {
		log.Infof("Visiting %v", request)
	})

	err := c.Visit(url)
	if err != nil {
		log.WithError(err)
	}
	c.Wait()

	log.Info(c.AllowedDomains)


}