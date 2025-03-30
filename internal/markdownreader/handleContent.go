package markdownreader

import (
	"strings"
)

func checkWordStart(word string, specialChar string, style string) string {

	word = strings.TrimPrefix(word, specialChar) // Remove the prefix

	// Check if the word also end the style
	if strings.HasSuffix(word, specialChar) {
		word = strings.TrimSuffix(word, specialChar) // Remove the suffix

		word = style + word + colors.Reset()
		return word
	}

	word = style + word
	return word
}

func wordStyleEnd(word string, specialChar string) string {
	word = strings.TrimSuffix(word, specialChar) // Remove the suffix

	word = word + colors.Reset()
	return word
}

func handleLine(line string) string {
	line = strings.TrimPrefix(line, " ")
	words := strings.Split(line, " ")
	var returnSlice []string

	// Loop through words and check if they have special chars
	for _, word := range words {

		/// Inline code ///
		// Check for inline code start
		if strings.HasPrefix(word, "`") {
			word = strings.TrimPrefix(word, "`") // Remove the ` prefix

			// Check if the word also end the inline code
			if strings.HasSuffix(word, "`") {
				word = strings.TrimSuffix(word, "`") // Remove the ` suffix

				word = colors.BgBlack() + " " + colors.Red() + word + " " + colors.Reset()
				returnSlice = append(returnSlice, word)
				continue
			}

			word = colors.BgBlack() + " " + colors.Red() + word
			returnSlice = append(returnSlice, word)
			continue
		}
		// Check for inline code stop
		if strings.HasSuffix(word, "`") {
			result := wordStyleEnd(word, "`")
			returnSlice = append(returnSlice, result)
			continue
		}

		/// Bold ///

		// Check for bold start
		if strings.HasPrefix(word, "**") {
			result := checkWordStart(word, "**", colors.Bold())
			returnSlice = append(returnSlice, result)
			continue
		}

		// Check for bold stop
		if strings.HasSuffix(word, "**") {
			result := wordStyleEnd(word, "**")
			returnSlice = append(returnSlice, result)
			continue
		}

		/// Italic ///

		// Check for italic start
		if strings.HasPrefix(word, "*") {
			result := checkWordStart(word, "*", colors.Italic())
			returnSlice = append(returnSlice, result)
			continue
		}

		// Check for italic stop
		if strings.HasSuffix(word, "*") {
			result := wordStyleEnd(word, "*")
			returnSlice = append(returnSlice, result)
			continue
		}

		/// Strikethrough ///

		// Check for strikethrough start
		if strings.HasPrefix(word, "~") {
			result := checkWordStart(word, "~", colors.Strikethrough())
			returnSlice = append(returnSlice, result)
			continue
		}

		// Check for strikethrough stop
		if strings.HasSuffix(word, "~") {
			result := wordStyleEnd(word, "~")
			returnSlice = append(returnSlice, result)
			continue
		}

		// Add the word to the return slice
		returnSlice = append(returnSlice, word)
	}

	// Convert the return slice to a string to return it
	returnString := strings.Join(returnSlice, " ")
	return returnString
}
