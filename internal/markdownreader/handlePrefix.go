package markdownreader

import (
	"strings"
)

func handlePrefix(line string) string {

	// Handle big title
	if strings.HasPrefix(line, "# ") {
		returnLine := handleBigTitle(line)
		return returnLine
	}

	// Handle other titles
	if strings.HasPrefix(line, "## ") || strings.HasPrefix(line, "### ") || strings.HasPrefix(line, "#### ") || strings.HasPrefix(line, "##### ") || strings.HasPrefix(line, "###### ") {
		returnLine := handleTitle(line)
		return returnLine
	}

	// Handle bullet list
	if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "+ ") || strings.HasPrefix(line, "* ") {
		returnLine := handleBulletList(line)
		return returnLine
	}

	// Handle numbered list
	if strings.HasPrefix(line, "1") || strings.HasPrefix(line, "2") || strings.HasPrefix(line, "3") || strings.HasPrefix(line, "4") || strings.HasPrefix(line, "5") || strings.HasPrefix(line, "6") || strings.HasPrefix(line, "7") || strings.HasPrefix(line, "8") || strings.HasPrefix(line, "9") {
		returnLine := handleNumberedList(line)
		return returnLine
	}

	return line
}
