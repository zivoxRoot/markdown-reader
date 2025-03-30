package markdownreader

import (
	"strings"
)

// TODO: Group the two functions

func checkWordStart(word string, specialChar string, style string, background string) string {

	word = strings.TrimPrefix(word, specialChar) // Remove the prefix

	// Check if the word also end the style
	if strings.HasSuffix(word, specialChar) {

		word = strings.TrimSuffix(word, specialChar) // Remove the suffix

		// Check if it has a background
		if background != "" {
			word = background + " " + style + word + " " + colors.Reset()
			return word
		} 

		word = style + word + colors.Reset()
		return word
	}

	// Check if it has a background
	if background != "" {
		word = background + " " + style + word
		return word
	}

	word = style + word
	return word
}

func checkWordEnd(word string, specialChar string, hasBackground bool) string {

	word = strings.TrimSuffix(word, specialChar) // Remove the suffix

	// Check if it has a background
	if hasBackground == true {
		return word + " " + colors.Reset()
	}

	return word + colors.Reset()
}

func handleLine(line string) string {
	line = strings.TrimPrefix(line, " ")
	words := strings.Split(line, " ")
	var returnSlice []string

	// Loop through words and check if they have special chars
	for _, word := range words {

		/// Inline code ///
		if strings.HasPrefix(word, "`") {
			result := checkWordStart(word, "`", colors.Red(), colors.BgBlack())
			returnSlice = append(returnSlice, result)
			continue
		}
		if strings.HasSuffix(word, "`") {
			result := checkWordEnd(word, "`", true)
			returnSlice = append(returnSlice, result)
			continue
		}

		/// Bold ///
		if strings.HasPrefix(word, "**") {
			result := checkWordStart(word, "**", colors.Bold(), "")
			returnSlice = append(returnSlice, result)
			continue
		}
		if strings.HasSuffix(word, "**") {
			result := checkWordEnd(word, "**", false)
			returnSlice = append(returnSlice, result)
			continue
		}

		/// Italic ///
		if strings.HasPrefix(word, "*") {
			result := checkWordStart(word, "*", colors.Italic(), "")
			returnSlice = append(returnSlice, result)
			continue
		}
		// Check for italic stop
		if strings.HasSuffix(word, "*") {
			result := checkWordEnd(word, "*", false)
			returnSlice = append(returnSlice, result)
			continue
		}

		/// Strikethrough ///
		if strings.HasPrefix(word, "~") {
			result := checkWordStart(word, "~", colors.Strikethrough(), "")
			returnSlice = append(returnSlice, result)
			continue
		}
		// Check for strikethrough stop
		if strings.HasSuffix(word, "~") {
			result := checkWordEnd(word, "~", false)
			returnSlice = append(returnSlice, result)
			continue
		}

		// Add the word to the return slice
		returnSlice = append(returnSlice, word)
	}

	// Convert the return slice to a string to return it
	return strings.Join(returnSlice, " ")
}
