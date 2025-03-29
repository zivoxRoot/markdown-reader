package markdownreader

import (
	"strings"
)

func handleQuotes(line string) string {
	var returnLine string

	returnLine = strings.TrimPrefix(line, "> ")
	return colors.Dim() + "│ " + returnLine + colors.Reset()
}
