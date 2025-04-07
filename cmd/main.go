package main

import (
	"fmt"
	"os"

	"github.com/zivoxRoot/markdown-reader/internal/help"
	"github.com/zivoxRoot/markdown-reader/internal/markdownreader"
)

func main() {
	// Get the file to read
	args := os.Args
	if len(args) == 1 || len(args) > 2 {
		help.ShowHelp()
		os.Exit(0)
	}
	// Check if the first args is help or version
	if args[1] == "-h" || args[1] == "--help" {
		help.ShowHelp()
		os.Exit(0)
	}
	if args[1] == "-v" || args[1] == "--version" {
		help.ShowVersion()
		os.Exit(0)
	}

	file := args[1]

	// Initialize a new markdown reader instance.
	reader := markdownreader.NewReader(file)
	result, err := reader.ProcessMarkdown()

	// Check for errors.
	if err != nil {
		fmt.Println("Markdown reader :", err)
		os.Exit(1)
	}

	// Print the stylized markdown.
	for _, line := range result {
		fmt.Println("  " + line)
	}
}