package main

import "fmt"

func main() {
	// Prompt the user to enter a year
	fmt.Printf("Enter a year:")
	var year int
	fmt.Scanln(&year)

	// Check if it's a leap year
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}
