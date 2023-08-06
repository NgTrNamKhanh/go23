package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("https://topdev.vn/viec-lam-it"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("div[id=scroll-it-jobs]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		fmt.Println(e.Text)
		fmt.Println("Hello")
	})
	c.Visit("https://topdev.vn/viec-lam-it/python-kt34")

}