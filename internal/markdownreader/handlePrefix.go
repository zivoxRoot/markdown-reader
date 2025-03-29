package markdownreader

import (
	"strings"
)

func handlePrefix(line string) string {

	// Handle title 1
	if strings.HasPrefix(line, "# ") {
		returnLine := handleTitle1(line)
		return returnLine
	}

	// Handle title 2
	if strings.HasPrefix(line, "## ") {
		returnLine := handleTitle2(line)
		return returnLine
	}

	// Handle title 3
	if strings.HasPrefix(line, "### ") {
		returnLine := handleTitle3(line)
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
