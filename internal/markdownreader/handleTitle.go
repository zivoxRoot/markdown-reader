package markdownreader

import (
	"strings"
)

func handleBigTitle(line string) string {
	var returnLine string

	returnLine = strings.TrimPrefix(line, "# ")
	returnLine = colors.BrightBgBlue() + colors.Bold() + colors.BrightWhite() + " " + returnLine + " " + colors.Reset() + "\n"
	return returnLine
}

func handleTitle(line string) string {
	var returnLine string

	returnLine = colors.BrightCyan() + line + colors.Reset()
	return returnLine
}
