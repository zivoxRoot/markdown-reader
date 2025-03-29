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

		/// Bold ///

		// Check for bold start
		if strings.HasPrefix(word, "**") {
			word = strings.TrimPrefix(word, "**") // Remove the ** prefix

			// Check if the word also end the bold
			if strings.HasSuffix(word, "**") {
				word = strings.TrimSuffix(word, "**") // Remove the ** suffix

				word = colors.Blue() + word + colors.Reset()
				returnSlice = append(returnSlice, word)
				continue
			}

			word = colors.Blue() + word
			returnSlice = append(returnSlice, word)
			continue
		}

		// Check for bold stop
		if strings.HasSuffix(word, "**") {
			word = strings.TrimSuffix(word, "**") // Remove the ** suffix

			word = word + colors.Reset()
			returnSlice = append(returnSlice, word)
			continue
		}

		/// Italic ///

		// Check for italic start
		if strings.HasPrefix(word, "*") {
			word = strings.TrimPrefix(word, "*") // Remove the * prefix

			// Check if the word also end the italic
			if strings.HasSuffix(word, "*") {
				word = strings.TrimSuffix(word, "*") // Remove the * suffix

				word = colors.Italic() + word + colors.Reset()
				returnSlice = append(returnSlice, word)
				continue
			}

			word = colors.Italic() +word
			returnSlice = append(returnSlice, word)
			continue
		}

		// Check for italic stop
		if strings.HasSuffix(word, "*") {
			word = strings.TrimSuffix(word, "*") // Remove the * suffix

			word = word + colors.Reset()
			returnSlice = append(returnSlice, word)
			continue
		}

		/// Strikethrough ///

		// Check for strikethrough start
		if strings.HasPrefix(word, "~") {
			word = strings.TrimPrefix(word, "~") // Remove the ~ prefix

			// Check if the word also end the strikethrough
			if strings.HasSuffix(word, "~") {
				word = strings.TrimSuffix(word, "~") // Remove the ~ suffix

				word = colors.Strikethrough() + word + colors.Reset()
				returnSlice = append(returnSlice, word)
				continue
			}

			word = colors.Strikethrough() +word
			returnSlice = append(returnSlice, word)
			continue
		}

		// Check for strikethrough stop
		if strings.HasSuffix(word, "~") {
			word = strings.TrimSuffix(word, "~") // Remove the ~ suffix

			word = word + colors.Reset()
			returnSlice = append(returnSlice, word)
			continue
		}

		// Add the word to the return slice
		returnSlice = append(returnSlice, word)
	}

	// Convert the return slice to a string to return it
	returnString := strings.Join(returnSlice, " ")
	return returnString
}
