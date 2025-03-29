package markdownreader

import (
	"strings"

	"github.com/zivoxRoot/markdown-reader/internal/colors"
)

func handleContent(line string) string {
	line = strings.TrimPrefix(line, " ")
	words := strings.Split(line, " ")
	var returnSlice []string

	// Loop through words and check if they have special chars
	for _, word := range words {

		// Start bold
		if strings.HasPrefix(word, "**") {
			word = strings.TrimPrefix(word, "**") // Remove the ** prefix

			// Check if the word also end the bold
			if strings.HasSuffix(word, "**") {
				word = strings.TrimSuffix(word, "**") // Remove the ** suffix

				returnSlice = append(returnSlice, colors.Blue())
				returnSlice = append(returnSlice, word)
				returnSlice = append(returnSlice, colors.Reset())
				continue
			}

			returnSlice = append(returnSlice, colors.Blue())
			returnSlice = append(returnSlice, word)
			continue
		}

		// Stop the bold
		if strings.HasSuffix(word, "**") {
			word = strings.TrimSuffix(word, "**") // Remove the ** suffix

			returnSlice = append(returnSlice, word)
			returnSlice = append(returnSlice, colors.Reset())
			continue
		}

		// Add the word to the return slice
		returnSlice = append(returnSlice, word)
	}

	// Convert the return slice to a string to return it
	returnString := strings.Join(returnSlice, " ")
	return returnString
}
