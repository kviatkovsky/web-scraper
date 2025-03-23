package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/kviatkovsky/web-scraper/internal/scraper/otomoto"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.otomoto.pl"),
	)
	otomotoCars := otomoto.Scrap(c)
	fmt.Println(otomotoCars)
	fmt.Println("END")
}
