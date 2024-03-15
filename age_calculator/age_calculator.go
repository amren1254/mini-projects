package main

import (
	"fmt"
	"time"
)

func main() {

	var birthdateInput string
	// Prompt the user to enter their birthdate
	fmt.Println("Enter your birthdate (YYYY-MM-DD): ")
	fmt.Scanln(&birthdateInput)

	// Parse the user input into a time.Time object
	birthDate, err := time.Parse("2006-01-02", birthdateInput)
	if err != nil {
		fmt.Println("Error parsing birthdate:", err)
		return
	}
	// Get the current date
	currentDate := time.Now()

	// Calculate the difference
	difference := currentDate.Sub(birthDate)

	// Convert the difference to various units
	days := int(difference.Hours() / 24)
	weeks := days / 7
	months := int(difference.Hours() / (24 * 30))
	years := currentDate.Year() - birthDate.Year()
	hours := int(difference.Hours())
	minutes := int(difference.Minutes())
	seconds := int(difference.Seconds())

	// Output the results
	fmt.Println("Your Age:")
	fmt.Printf("Days    : %d\n", days)
	fmt.Printf("Weeks   : %d\n", weeks)
	fmt.Printf("Months  : %d\n", months)
	fmt.Printf("Years   : %d\n", years)
	fmt.Printf("Hours   : %d\n", hours)
	fmt.Printf("Minutes : %d\n", minutes)
	fmt.Printf("Seconds : %d\n", seconds)

}
