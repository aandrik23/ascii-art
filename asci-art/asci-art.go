package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if arguments are passed
	if len(os.Args) == 1 {
		fmt.Println("Usage: go run . [STRING] [BANNER] || Example: go run . \"test\" standard")
		return
	}

	// The input text to be rendered in ASCII art
	argStr := os.Args[1]
	var styleBanner string

	// Check if the user specified a custom banner
	if len(os.Args) == 3 {
		styleBanner = strings.ToLower(os.Args[2]) // Custom banner
	} else {
		styleBanner = "standard" // Default banner if no banner is provided
	}

	// Read the banner file (styleBanner)
	file, err := os.ReadFile(styleBanner + ".txt")
	if err != nil {
		fmt.Println(styleBanner + " banner does not exist.")
		return
	}

	// Split the banner file into lines for ASCII art mapping
	lines := strings.Split(string(file), "\n")
	sepArgs := strings.Split(argStr, "\\n")

	// Print the ASCII art to the console
	printAsciiArt(sepArgs, lines)
}

// Function to print ASCII art to the console
func printAsciiArt(sentences []string, textFile []string) {
	for i, word := range sentences {
		if word == "" {
			if i != 0 {
				fmt.Println() // Print a new line for blank words
			}
			continue
		}
		for h := 1; h < 9; h++ { // ASCII art character height is 8
			for i := 0; i < len(word); i++ {
				for lineIndex, line := range textFile {
					if lineIndex == (int(word[i])-32)*9+h { // Map the character to ASCII art lines
						fmt.Print(line) // Print the corresponding line for the character
					}
				}
			}
			fmt.Println() // New line after each line of ASCII art
		}
	}
}
