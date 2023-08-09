package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly/v2"
	//"net/http"
)

func main() {
	url := "https://shopee.tw/api/v2/search_items/?by=pop&limit=30&match_id=1819984&newest=0&order=desc&page_type=shop&shop_categoryids=9271157&version=2"

	// Create a new collector
	c := colly.NewCollector()

	// Set custom headers
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:73.0) Gecko/20100101 Firefox/73.0")
		r.Headers.Set("X-Requested-With", "XMLHttpRequest")
		r.Headers.Set("Referer", "https://shopee.tw/shop/1819984/search?shopCollection=9271157")
	})

	// Variable to store the JSON response
	var data map[string]interface{}

	// Callback function to handle the JSON response
	c.OnResponse(func(r *colly.Response) {
		err := json.Unmarshal(r.Body, &data)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		// Extract and print data
		items := data["items"].([]interface{})
		for _, item := range items {
			itemMap := item.(map[string]interface{})
			fmt.Println("name:", itemMap["name"])
			fmt.Println("price:", itemMap["price"])
			fmt.Println("sold:", itemMap["historical_sold"])
			fmt.Println("---")
		}
	})

	// Start scraping
	err := c.Visit(url)
	if err != nil {
		fmt.Println("Error:", err)
	}
}