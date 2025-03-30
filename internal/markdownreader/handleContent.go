package markdownreader

import (
	"strings"
)

// checkWordStart remove a markdown format sign at the beginning of a word and style it.
func checkWordStart(word string, specialChar string, style string, background string) string {

	word = strings.TrimPrefix(word, specialChar)

	// Check if the word also end the style
	if strings.HasSuffix(word, specialChar) {

		word = strings.TrimSuffix(word, specialChar)

		if background != "" {
			word = background + " " + style + word + " " + colors.Reset()
			return word
		}

		word = style + word + colors.Reset()
		return word
	}

	if background != "" {
		word = background + " " + style + word
		return word
	}

	word = style + word
	return word
}

// checkWordEnd removes a markdown format sign at the end of a word and style it.
func checkWordEnd(word string, specialChar string, hasBackground bool) string {

	word = strings.TrimSuffix(word, specialChar)

	if hasBackground == true {
		return word + " " + colors.Reset()
	}

	return word + colors.Reset()
}

// handleLine loops through the words of a line, checks if they have markdown format sign, then style them.
func handleLine(line string) string {
	line = strings.TrimPrefix(line, " ")
	words := strings.Split(line, " ")
	var returnSlice []string

	// Loop through words to check if they have special chars
	for _, word := range words {

		// Inline code
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

		// Bold
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

		// Italic
		if strings.HasPrefix(word, "*") {
			result := checkWordStart(word, "*", colors.Italic(), "")
			returnSlice = append(returnSlice, result)
			continue
		}
		if strings.HasSuffix(word, "*") {
			result := checkWordEnd(word, "*", false)
			returnSlice = append(returnSlice, result)
			continue
		}

		// Strikethrough
		if strings.HasPrefix(word, "~") {
			result := checkWordStart(word, "~", colors.Strikethrough(), "")
			returnSlice = append(returnSlice, result)
			continue
		}
		if strings.HasSuffix(word, "~") {
			result := checkWordEnd(word, "~", false)
			returnSlice = append(returnSlice, result)
			continue
		}

		returnSlice = append(returnSlice, word)
	}

	// Convert the return slice to a string to return it
	return strings.Join(returnSlice, " ")
}
