package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	// Open the JSON file
	file, err := os.Open("json_to_csv/input.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	// Decode JSON data
	var data []map[string]interface{}
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Get current timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")

	// Create filename with timestamp
	filename := fmt.Sprintf("output_%s.csv", timestamp)

	// Create the CSV file
	csvFile, err := os.Create("json_to_csv/" + filename)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Create a CSV writer
	csvWriter := csv.NewWriter(csvFile)

	// Extract headers from the first JSON object
	if len(data) == 0 {
		fmt.Println("No data found in JSON file.")
		return
	}
	var headers []string
	for key := range data[0] {
		headers = append(headers, key)
	}

	// Write header
	err = csvWriter.Write(headers)
	if err != nil {
		fmt.Println("Error writing CSV header:", err)
		return
	}

	// Write data rows
	for _, record := range data {
		var row []string
		for _, header := range headers {
			row = append(row, fmt.Sprintf("%v", record[header]))
		}
		err := csvWriter.Write(row)
		if err != nil {
			fmt.Println("Error writing CSV row:", err)
			return
		}
	}

	// Flush the CSV writer
	csvWriter.Flush()

	// Check for errors
	if err := csvWriter.Error(); err != nil {
		fmt.Println("Error flushing CSV writer:", err)
		return
	}

	fmt.Printf("Data flushed into %s successfully.\n", filename)
}
