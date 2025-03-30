package markdownreader

import (
	"strings"
)

// handleBigTitle styles markdown titles with one '#'.
func handleBigTitle(line string) string {
	var returnLine string

	returnLine = strings.TrimPrefix(line, "# ")
	returnLine = colors.Reset() + colors.BrightBgBlue() + colors.Bold() + colors.BrightWhite() + " " + returnLine + " " + colors.Reset() + "\n"
	return returnLine
}

// handleTitle styles markdown titles that start with multiple '#'.
func handleTitle(line string) string {
	var returnLine string

	returnLine = colors.Reset() + colors.BrightCyan() + line + colors.Reset()
	return returnLine
}
