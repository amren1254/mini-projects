package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generatePassword(length int) string {
	charset := map[int]string{
		0: "abcdefghijklmnopqrstuvwxyz",
		1: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		2: "0123456789",
		3: "!@#$%^&*()_+{}[];:,.<>?/~`",
	}

	// Use a seed based on current time for better randomness
	rand.Seed(time.Now().UnixNano())

	var password string
	for i := 0; i < length; i++ {
		// Randomly select a character set
		charsetIndex := rand.Intn(len(charset))
		// Randomly select a character from the chosen character set
		charIndex := rand.Intn(len(charset[charsetIndex]))
		password += string(charset[charsetIndex][charIndex])
	}
	return password
}

func main() {
	length := 12 // Change the length of the password as needed
	password := generatePassword(length)
	fmt.Println("Generated Password:", password)
}
