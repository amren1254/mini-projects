package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initialize random number generator with current time
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random number between 1 and 100
	target := rand.Intn(100) + 1

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I have selected a random number between 1 and 100.")
	fmt.Println("Try to guess the number!")

	var guess int
	for {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Error reading input. Please enter a valid number.")
			continue
		}

		// Check if the guess is correct
		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}
	}
}
