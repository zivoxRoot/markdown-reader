package main

import (
	"fmt"
	"os"

	"github.com/zivoxRoot/markdown-reader/internal/markdownreader"
)

func main() {
	// Initialize a new markdown reader instance
	reader := markdownreader.NewReader("../myFile.md")
	result, err := reader.ProcessMarkdown()

	// Check for errors
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	// Print the stylized markdown
	for _, line := range result {
		fmt.Println("  " + line)
	}
}
