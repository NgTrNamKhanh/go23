package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	//"strings"

	//"math/rand"
	"github.com/gocolly/colly/v2"
)

type Merch struct {
	Title  string `json:"title"`
	Price  string `json:"price"`
	Status string `json:"status"`
	Link   string `json:"link"`
}

func main() {

	merch := Merch{}
	merchs := make([]Merch, 0, 1)
	c := colly.NewCollector(
		colly.AllowedDomains("https://www.lazada.vn", "www.lazada.vn"),
	)

	c.OnRequest(func(r *colly.Request) {
		//r.Headers.Set("Accept-Language", "en-US;q=0,9")
		fmt.Println(fmt.Sprintf("Visiting %s", r.URL))
	})
	c.OnError(func(r *colly.Response, e error) {

		fmt.Printf("Error while scraping %s\n", e.Error())
	})
	c.OnScraped(func(r *colly.Response) {
		merchs = append(merchs, merch)
		merch = Merch{}
	})

	var wg sync.WaitGroup

	var productURLs []string
	go c.OnHTML("a.RfADt", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		fmt.Println(href)
		productURLs = append(productURLs, href)
	})
	wg.Add(1)

	go func() {
		defer wg.Done() // Decrement the WaitGroup counter when done
		err := c.Visit("https://www.lazada.vn/phu-kien-ban-phim/?q=milky%20yellow%20pro")
		if err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
	c.OnHTML("h1.merch_merchTitle__uPHuc", func(e *colly.HTMLElement) {
		merch.Title = e.Text
	})
	c.OnHTML("div.merch_priceContainer__7_rcQ span.Price_salePrice__KnHC9", func(h *colly.HTMLElement) {
		merch.Price = h.Text
	})
	c.OnHTML("div.merch_priceContainer__7_rcQ span.Price_price__bumLD", func(h *colly.HTMLElement) {
		merch.Price = h.Text
	})
	c.OnHTML("div.MerchInfo_container__qXYh_ tbody", func(h *colly.HTMLElement) {
		selection := h.DOM
		childNodes := selection.Children().Nodes
		merch.Status = selection.FindNodes(childNodes[0]).Find("span").Text()
	})

	for _, url := range productURLs {
		productURL := scrapeUrl(url)
		merch.Link = productURL
		c.Visit(productURL)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(merchs)

	content,err := json.Marshal(merchs)
	if err != nil{
		fmt.Println(err.Error())
	}
	os.WriteFile("merch.json", content, 0644)
}

//	func cleanDesc(s string) string {
//		return strings.TrimSpace(s)
//	}
func scrapeUrl(href string) string {
	return "https://www.merchbar.com" + href
}

