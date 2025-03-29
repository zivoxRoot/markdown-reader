package markdownreader

import (
	"strings"
)

func handlePrefix(line string) string {

	// Handle title 3
	if strings.HasPrefix(line, "### ") {
		returnLine := handleTitle3(line)
		return returnLine
	}

	// Handle title 2
	if strings.HasPrefix(line, "## ") {
		returnLine := handleTitle2(line)
		return returnLine
	}

	// Handle title 1
	if strings.HasPrefix(line, "# ") {
		returnLine := handleTitle1(line)
		return returnLine
	}

	// Handle bullet list
	if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "+ ") || strings.HasPrefix(line, "* ") {
		returnLine := handleBulletList(line)
		return returnLine
	}

	return line
}

func handleTitle1(line string) string {
	var returnLine string

	returnLine = strings.TrimPrefix(line, "# ")
	// returnLine = "\033[38;5;69m" + returnLine + "\033[0m"
	returnLine = "\033[1;38;5;69m" + returnLine + "\033[0m"
	return returnLine
}

func handleTitle2(line string) string {
	var returnLine string

	returnLine = "TITLE 2 : " + line
	return returnLine
}

func handleTitle3(line string) string {
	var returnLine string

	returnLine = "TITLE 3 : " + line
	return returnLine
}

func handleBulletList(line string) string {
	var returnLine string

	if strings.HasPrefix(line, "-") {
		returnLine = strings.TrimPrefix(line, "-")
	}
	if strings.HasPrefix(line, "+") {
		returnLine = strings.TrimPrefix(line, "+")
	}
	if strings.HasPrefix(line, "*") {
		returnLine = strings.TrimPrefix(line, "*")
	}

	returnLine = "•" + returnLine
	return returnLine
}
