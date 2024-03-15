package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	// Prompt the user to enter the file paths
	fmt.Println("Enter the path of the first file:")
	var file1Path string
	fmt.Scanln(&file1Path)

	fmt.Println("Enter the path of the second file:")
	var file2Path string
	fmt.Scanln(&file2Path)

	// Open the first file
	file1, err := os.Open(file1Path)
	if err != nil {
		fmt.Println("Error opening first file:", err)
		return
	}
	defer file1.Close()

	// Open the second file
	file2, err := os.Open(file2Path)
	if err != nil {
		fmt.Println("Error opening second file:", err)
		return
	}
	defer file2.Close()

	// Create scanners for both files
	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)

	// Define lipgloss styles for coloring output
	deletedStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#ff6666"))
	addedStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#66cc66"))
	unchangedStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))

	// Read lines from both files and compare them
	lineNumber := 1
	for scanner1.Scan() || scanner2.Scan() {
		text1 := scanner1.Text()
		text2 := scanner2.Text()

		// If one file has ended before the other, print error and exit
		if text1 == "" && text2 != "" {
			fmt.Println("First file has fewer lines than second file.")
			return
		} else if text1 != "" && text2 == "" {
			fmt.Println("Second file has fewer lines than first file.")
			return
		}

		// Compare lines
		if text1 != text2 {
			fmt.Printf("Line %d:\n", lineNumber)
			fmt.Printf("%s %s\n", deletedStyle.Render("-"), text1)
			fmt.Printf("%s %s\n\n", addedStyle.Render("+"), text2)
		} else {
			fmt.Printf("%s %s\n", unchangedStyle.Render(" "), text1)
		}

		lineNumber++
	}

	// Print success message if no differences found
	fmt.Println("Files are identical.")
}
