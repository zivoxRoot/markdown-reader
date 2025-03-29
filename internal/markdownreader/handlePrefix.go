package markdownreader

import (
	"strings"
)

func handlePrefix(line string) string {

	// TODO: Convert to switch case

	// Handle title 3
	if strings.HasPrefix(line, "###") {
		myLine := "TITLE 3 : " + line
		return myLine
	}

	// Handle title 2
	if strings.HasPrefix(line, "##") {
		myLine := "TITLE 2 : " + line
		return myLine
	}

	// Handle title 1
	if strings.HasPrefix(line, "#") {
		myLine := "TITLE 1 : " + line
		return myLine
	}

	// Handle bullet list
	if strings.HasPrefix(line, "-") {
		myLine := "BULLET POINT : " + line
		return myLine
	}

	return line
}
