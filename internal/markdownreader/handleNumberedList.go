package markdownreader

import (
	"strings"
	"strconv"

	"github.com/zivoxRoot/markdown-reader/internal/colors"
)

func handleNumberedList(line string) string {
	// Check if the fist word ends with a dot
	cleanLine := strings.TrimPrefix(line, " ")
	splittedLine := strings.Split(cleanLine, " ")
	firstWord := splittedLine[0]
	restOfLine := strings.TrimPrefix(cleanLine, firstWord)

	if strings.HasSuffix(firstWord, ".") {
		cleanFirstWord := strings.TrimSuffix(firstWord, ".")

		// Check if the first word is a number
		if _, err := strconv.Atoi(cleanFirstWord); err != nil {
			return line
		}

		return colors.Cyan() + firstWord + colors.Reset() + restOfLine
	}

	return line
}
