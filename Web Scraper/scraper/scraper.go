package scraper

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func scraper() {
	fname := "data.csv"
	file, err := os.Create(fname)

	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()
	c.OnHTML("table#customers", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			data := []string{
				el.ChildText("td:nth-child(1)"),
				el.ChildText("td:nth-child(2)"),
				el.ChildText("td:nth-child(3)"),
			}
			err := writer.Write(data)
			if err != nil {
				log.Printf("Error waiting to csv: %q", err)
			}
		})
		fmt.Println("Scrapping Complete")
	})
	c.Visit("https://www.w3schools.com/html/html_tables.asp")
}
