package otomoto

import (
	"github.com/gocolly/colly"
	"github.com/kviatkovsky/web-scraper/internal/model/car"
)

func Scrap(c *colly.Collector) []car.Car {
	otomotoCars := []car.Car{}
	c.OnHTML(`div[data-testid="search-results"] > div`, func(h *colly.HTMLElement) {
		otomotoCar := car.Car{}
		otomotoCar.Name = h.ChildText(`h2[data-sentry-element="Title"]`)
		otomotoCar.Price = h.ChildText(`h3[data-sentry-element="Price"]`)
		otomotoCar.Mileage = h.ChildText(`dd[data-parameter="mileage"]`)
		otomotoCar.Year = h.ChildText(`dd[data-parameter="year"]`)
		otomotoCar.Currency = h.ChildText(`p[data-sentry-element="PriceCurrency"]`)
		otomotoCar.Gearbox = h.ChildText(`dd[data-parameter="gearbox"]`)
		otomotoCar.Fuel = h.ChildText(`dd[data-parameter="fuel_type"]`)
		//otomotoCar.ImageUrl = h.ChildAttr("a", "href")
		//otomotoCar.Url = h.ChildAttr("a", "hrefy")
		otomotoCars = append(otomotoCars, otomotoCar)
	})
	c.Visit("https://www.otomoto.pl/osobowe/bmw/seria-5?search%5Bfilter_enum_fuel_type%5D=diesel&search%5Bfilter_enum_generation%5D=gen-g30-2017&search%5Bfilter_float_price%3Ato%5D=120000&search%5Border%5D=relevance_web")

	return otomotoCars
}
