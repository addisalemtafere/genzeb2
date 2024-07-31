package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func hello() {
	// URL of the Bank of Abyssinia exchange rates page
	url := "https://www.bankofabyssinia.com/ExchangeRate"

	// Make HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != 200 {
		log.Fatalf("Failed to get a valid response. Status Code: %d", resp.StatusCode)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("Failed to parse the response body: %v", err)
	}

	// Select the table using the provided XPath converted to CSS selectors
	doc.Find("div > div > div > div > div > div > div > div > div > div > div > div > div > div > table").Each(func(i int, table *goquery.Selection) {
		table.Find("tbody tr").Each(func(j int, row *goquery.Selection) {
			currency := row.Find("td:nth-child(1)").Text()
			buyRate := row.Find("td:nth-child(2)").Text()
			sellRate := row.Find("td:nth-child(3)").Text()
			
			fmt.Printf("Currency: %s, Buy Rate: %s, Sell Rate: %s\n", currency, buyRate, sellRate)
		})
	})
}
