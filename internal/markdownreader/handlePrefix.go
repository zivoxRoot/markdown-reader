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

	return line
}
