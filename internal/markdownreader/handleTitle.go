package markdownreader

import (
	"strings"

	"github.com/zivoxRoot/markdown-reader/internal/colors"
)

func handleBigTitle(line string) string {
	var returnLine string

	returnLine = strings.TrimPrefix(line, "# ")
	returnLine = colorer.BrightBgBlue() + colors.Bold() + " " + returnLine + " " + colors.Reset() + "\n"
	return returnLine
}

func handleTitle(line string) string {
	var returnLine string

	returnLine = colorer.BrightCyan() + line + colors.Reset()
	return returnLine
}
