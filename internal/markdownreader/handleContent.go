package markdownreader

import (
	"strings"
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

				word = colorer.Bold() + word + colorer.Reset()
				returnSlice = append(returnSlice, word)
				continue
			}

			word = colorer.Bold() + word
			returnSlice = append(returnSlice, word)
			continue
		}

		// Check for bold stop
		if strings.HasSuffix(word, "**") {
			word = strings.TrimSuffix(word, "**") // Remove the ** suffix

			word = word + colorer.Reset()
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

				word = colorer.Italic() + word + colorer.Reset()
				returnSlice = append(returnSlice, word)
				continue
			}

			word = colorer.Italic() + word
			returnSlice = append(returnSlice, word)
			continue
		}

		// Check for italic stop
		if strings.HasSuffix(word, "*") {
			word = strings.TrimSuffix(word, "*") // Remove the * suffix

			word = word + colorer.Reset()
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

				word = colorer.Strikethrough() + word + colorer.Reset()
				returnSlice = append(returnSlice, word)
				continue
			}

			word = colorer.Strikethrough() + word
			returnSlice = append(returnSlice, word)
			continue
		}

		// Check for strikethrough stop
		if strings.HasSuffix(word, "~") {
			word = strings.TrimSuffix(word, "~") // Remove the ~ suffix

			word = word + colorer.Reset()
			returnSlice = append(returnSlice, word)
			continue
		}

		/// Inline code ///

		// Check for inline code start
		if strings.HasPrefix(word, "`") {
			word = strings.TrimPrefix(word, "`") // Remove the ` prefix

			// Check if the word also end the inline code
			if strings.HasSuffix(word, "`") {
				word = strings.TrimSuffix(word, "`") // Remove the ` suffix

				word = colorer.BgBlack() + " " + colorer.Red() + word + " " + colorer.Reset()
				returnSlice = append(returnSlice, word)
				continue
			}

			word = colorer.BgBlack() + " " + colorer.Red() + word
			returnSlice = append(returnSlice, word)
			continue
		}

		// Check for inline code stop
		if strings.HasSuffix(word, "`") {
			word = strings.TrimSuffix(word, "`") // Remove the ` suffix

			word = word + " " + colorer.Reset()
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
