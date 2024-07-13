package main

import (
	// importing Colly
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"tc-web-scraper/utils"

	"github.com/gocolly/colly"
) 

type TradingCard struct { 
	url, image, name, price string 
}

 
func main() { 
	c := colly.NewCollector()

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"

	c.SetRequestTimeout(120 * time.Second)

	var cards []TradingCard

	// iterating over the list of HTML trading card elements 
	c.OnHTML("li.brwrvr__item-card", func(e *colly.HTMLElement) {
		card := TradingCard{}
	
		// scraping the data of interest 
		card.url = e.ChildAttr("a.bsig__title__wrapper", "href") 
		card.image = e.ChildAttr("img.brwrvr__item-card__image", "src") 
		card.name = e.ChildText("h3.textual-display.bsig__title__text") 
		card.price = e.ChildText("span.textual-display.bsig__price.bsig__price--displayprice")
	 
		cards = append(cards, card)
	})

	c.OnScraped(func(r *colly.Response) {
		// Get the desktop path
    desktopPath, err := utils.GetDesktopPath()
    if err != nil {
        log.Fatalf("Failed to get desktop path: %v", err)
    }

		timestamp := time.Now().Format("2006-01-02T15-04-05")

		filename := filepath.Join(desktopPath, fmt.Sprintf(`trading-cards-%s.csv`, timestamp))

		// opening the CSV file
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalln("Failed to create output CSV file", err)
		}
		defer file.Close()

		writer := csv.NewWriter(file)

		// writing the CSV headers
		headers := []string{
			"url",
			"image",
			"name",
			"price",
		}
		writer.Write(headers)

		// writing each card as a CSV row
		for _, card := range cards {
			// converting a TradingCard to an array of strings
			record := []string{
				card.url,
				card.image,
				card.name,
				card.price,
			}

			// adding a CSV record to the output file
			writer.Write(record)
		}
		defer writer.Flush()
	})

	err := c.Visit("https://www.ebay.com/b/Major-Leagues-MLB-Baseball-Sports-Trading-Cards-Accessories/212/bn_17106685?Graded=Yes&_pgn=2&mag=1&rt=nc")
	if err != nil {
		log.Fatal(err)
	}

}
