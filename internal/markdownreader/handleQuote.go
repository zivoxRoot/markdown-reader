package markdownreader

import (
	"strings"
)

func handleQuotes(line string) string {
	var returnLine string

	returnLine = strings.TrimPrefix(line, "> ")
	return colorer.Dim() + "│ " + returnLine + colorer.Reset()
}
