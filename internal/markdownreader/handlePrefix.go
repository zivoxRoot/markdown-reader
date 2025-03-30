package markdownreader

import (
	"strings"
)

// handlePrefix check the first characters of a line to see if it is a special markdown element.
func handlePrefix(line string) string {

	// Big title
	if strings.HasPrefix(line, "# ") {
		returnLine := handleBigTitle(line)
		return returnLine
	}

	// Other titles
	if strings.HasPrefix(line, "## ") || strings.HasPrefix(line, "### ") || strings.HasPrefix(line, "#### ") || strings.HasPrefix(line, "##### ") || strings.HasPrefix(line, "###### ") {
		returnLine := handleTitle(line)
		return returnLine
	}

	// Bullet list
	if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "+ ") || strings.HasPrefix(line, "* ") {
		returnLine := handleBulletList(line)
		return returnLine
	}

	// Numbered list
	if strings.HasPrefix(line, "1") || strings.HasPrefix(line, "2") || strings.HasPrefix(line, "3") || strings.HasPrefix(line, "4") || strings.HasPrefix(line, "5") || strings.HasPrefix(line, "6") || strings.HasPrefix(line, "7") || strings.HasPrefix(line, "8") || strings.HasPrefix(line, "9") {
		returnLine := handleNumberedList(line)
		return returnLine
	}

	// Quotes
	if strings.HasPrefix(line, "> ") {
		returnLine := handleQuotes(line)
		return returnLine
	}

	// Code block
	if strings.HasPrefix(line, "```") {
		returnLine := handleCodeBlock()
		return returnLine
	}

	return line
}
