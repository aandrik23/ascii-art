package funcs

import (
	"ascii/funcs"
	"bytes"
	"os"
	"strings"
	"testing"
)

// Capture output to compare with expected ASCII art
func captureOutput(f func()) string {
	// Create a pipe to capture standard output
	r, w, _ := os.Pipe()

	// Save the original stdout
	old := os.Stdout
	defer func() {
		os.Stdout = old // Restore original stdout
	}()
	os.Stdout = w // Redirect stdout to the pipe

	f() // Call the function that produces output

	// Close the write end and read from the read end
	w.Close()

	var buf bytes.Buffer
	buf.ReadFrom(r) // Read from the pipe into the buffer

	return buf.String() // Return the captured output as a string
}

func TestPrintAsciiArt(t *testing.T) {
	// Specify the banner file name
	styleBanner := "standard" // Change this to your desired banner style
	bannerFileName := "../banners/" + styleBanner + ".txt"

	// Read the banner file
	fileContent, err := os.ReadFile(bannerFileName)
	if err != nil {
		t.Fatalf("Failed to read banner file: %v", err)
	}

	// Replace any Windows line endings with Unix style
	fileContentString := strings.ReplaceAll(string(fileContent), "\r\n", "\n")
	bannerLines := strings.Split(fileContentString, "\n")

	// Define input for testing
	input := []string{"A"} // Change this to the character you want to test

	// Calculate expected output
	expectedOutput := strings.Join(bannerLines[(int(input[0][0])-32)*9+1:(int(input[0][0])-32)*9+9], "\n") + "\n"

	// Capture the output of the PrintAsciiArt function
	actualOutput := captureOutput(func() {
		funcs.PrintAsciiArt(input, bannerLines)
	})

	// Compare actual output with expected output
	if actualOutput != expectedOutput {
		t.Errorf("Expected:\n%s\nGot:\n%s", expectedOutput, actualOutput)
	}
}
