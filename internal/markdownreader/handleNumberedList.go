package markdownreader

import (
	"strconv"
	"strings"

	"github.com/zivoxRoot/markdown-reader/internal/colors"
)

func handleNumberedList(line string) string {
	// Check if the fist word ends with a dot
	line = strings.TrimPrefix(line, " ")
	words := strings.Split(line, " ")
	firstWord := words[0]
	restOfLine := strings.TrimPrefix(line, firstWord)

	if strings.HasSuffix(firstWord, ".") {
		cleanFirstWord := strings.TrimSuffix(firstWord, ".")

		// Check if the first word is a number
		if _, err := strconv.Atoi(cleanFirstWord); err != nil {
			return line
		}

		return colorer.Cyan() + firstWord + colors.Reset() + restOfLine
	}

	return line
}
