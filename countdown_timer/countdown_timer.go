package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	// Prompt the user to enter the duration in seconds
	fmt.Println("Enter the duration in seconds:")
	var durationInput string
	fmt.Scanln(&durationInput)

	// Convert duration input to an integer
	duration, err := strconv.Atoi(durationInput)
	if err != nil {
		fmt.Println("Invalid duration input.")
		return
	}

	// Set up a channel to listen for interrupt signals (e.g., Ctrl+C)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Create a timer
	timer := time.NewTimer(time.Duration(duration) * time.Second)

	// Loop to continuously check the timer
	for {
		select {
		case <-timer.C:
			// Timer has expired
			fmt.Println()
			fmt.Println("Time's up!")
			return
		case <-interrupt:
			// Interrupt signal received, stop the timer
			timer.Stop()
			fmt.Println("\nTimer stopped.")
			return
		default:
			// Print remaining time
			fmt.Printf("\rTime remaining: %d seconds", duration)
			time.Sleep(1 * time.Second)
			duration--
		}
	}
}
