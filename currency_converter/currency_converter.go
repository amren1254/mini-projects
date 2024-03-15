package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const apiURL = "https://api.exchangeratesapi.io/latest"

type ExchangeRates struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
	Date  string             `json:"date"`
}

func main() {
	// Fetch exchange rates from the API
	rates, err := fetchExchangeRates()
	if err != nil {
		fmt.Println("Error fetching exchange rates:", err)
		return
	}

	// Print available currencies
	fmt.Println("Available currencies:")
	for currency := range rates.Rates {
		fmt.Println(currency)
	}

	// Prompt user for input
	fmt.Println("Enter amount to convert (e.g., 100 USD to EUR):")
	var input string
	fmt.Scanln(&input)

	// Parse input
	parts := strings.Split(input, " ")
	if len(parts) != 4 {
		fmt.Println("Invalid input format. Please enter in the format: amount source_currency to target_currency")
		return
	}

	amountStr := parts[0]
	sourceCurrency := strings.ToUpper(parts[1])
	targetCurrency := strings.ToUpper(parts[3])

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount:", err)
		return
	}

	// Check if currencies are available
	if _, ok := rates.Rates[sourceCurrency]; !ok {
		fmt.Println("Source currency not found.")
		return
	}
	if _, ok := rates.Rates[targetCurrency]; !ok {
		fmt.Println("Target currency not found.")
		return
	}

	// Convert currency
	sourceRate := rates.Rates[sourceCurrency]
	targetRate := rates.Rates[targetCurrency]
	result := (amount / sourceRate) * targetRate

	// Print result
	fmt.Printf("%.2f %s = %.2f %s\n", amount, sourceCurrency, result, targetCurrency)
}

func fetchExchangeRates() (*ExchangeRates, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rates ExchangeRates
	err = json.Unmarshal(body, &rates)
	if err != nil {
		return nil, err
	}

	return &rates, nil
}
