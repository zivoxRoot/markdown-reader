package markdownreader

import (
	"strings"
)

func handleBigTitle(line string) string {
	var returnLine string

	returnLine = strings.TrimPrefix(line, "# ")
	returnLine = colorer.BrightBgBlue() + colorer.Bold() + colorer.BrightWhite() + " " + returnLine + " " + colorer.Reset() + "\n"
	return returnLine
}

func handleTitle(line string) string {
	var returnLine string

	returnLine = colorer.BrightCyan() + line + colorer.Reset()
	return returnLine
}
