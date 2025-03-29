package markdownreader

import (
	"strings"

	"github.com/zivoxRoot/markdown-reader/internal/colors"
)

func handleTitle1(line string) string {
	var returnLine string

	returnLine = strings.TrimPrefix(line, "# ")
	returnLine = colors.BrightBgBlue() + colors.Bold() + " " + returnLine + " " + colors.Reset() + "\n"
	return returnLine
}

func handleTitle2(line string) string {
	var returnLine string

	returnLine = colors.BrightCyan() + line + colors.Reset()
	return returnLine
}

func handleTitle3(line string) string {
	var returnLine string

	returnLine = colors.BrightCyan() + line + colors.Reset()
	return returnLine
}
