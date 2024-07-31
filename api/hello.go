package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// ExchangeRate represents the structure of an exchange rate entry
type ExchangeRate struct {
	Currency   string
	BuyRate    string
	SellRate   string
}

// FetchExchangeRates fetches and parses the exchange rates from the Bank of Abyssinia website
func FetchExchangeRates(url string) ([]ExchangeRate, error) {
	// Make HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to get a valid response. Status Code: %d", resp.StatusCode)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the response body: %w", err)
	}

	var rates []ExchangeRate

	// Select the table using the CSS selectors
	doc.Find("div > div > div > div > div > div > div > div > div > div > div > div > div > div > table").Each(func(i int, table *goquery.Selection) {
		table.Find("tbody tr").Each(func(j int, row *goquery.Selection) {
			currency := row.Find("td:nth-child(1)").Text()
			buyRate := row.Find("td:nth-child(2)").Text()
			sellRate := row.Find("td:nth-child(3)").Text()
			
			rate := ExchangeRate{
				Currency: currency,
				BuyRate:  buyRate,
				SellRate: sellRate,
			}
			rates = append(rates, rate)
		})
	})

	return rates, nil
}

func main() {
	url := "https://www.bankofabyssinia.com"

	rates, err := FetchExchangeRates(url)
	if err != nil {
		log.Fatalf("Error fetching exchange rates: %v", err)
	}

	for _, rate := range rates {
		fmt.Printf("Currency: %s, Buy Rate: %s, Sell Rate: %s\n", rate.Currency, rate.BuyRate, rate.SellRate)
	}
}
