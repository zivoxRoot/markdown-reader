package markdownreader

import (
	"strconv"
	"strings"
)

// handleNumberedList colors a numbered list element.
func handleNumberedList(line string) string {
	line = strings.TrimPrefix(line, " ")
	words := strings.Split(line, " ")
	firstWord := words[0]
	restOfLine := strings.TrimPrefix(line, firstWord)

	// Check if the fist word ends with a dot.
	if strings.HasSuffix(firstWord, ".") {
		cleanFirstWord := strings.TrimSuffix(firstWord, ".")

		// Check if the first word is a number.
		if _, err := strconv.Atoi(cleanFirstWord); err != nil {
			return line
		}

		return colors.Reset() + colors.Cyan() + firstWord + colors.Reset() + restOfLine
	}

	return line
}
