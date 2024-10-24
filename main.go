package main

import (
	"ascii/funcs"
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
	file, err := os.ReadFile("banners/" + styleBanner + ".txt")
	if err != nil {
		fmt.Println(styleBanner + " banner does not exist.")
		return
	}

	fileContent := strings.ReplaceAll(string(file), "\r\n", "\n")
	lines := strings.Split(fileContent, "\n")

	sepArgs := strings.Split(argStr, "\\n")

	// Print the ASCII art to the console
	funcs.PrintAsciiArt(sepArgs, lines)
}
