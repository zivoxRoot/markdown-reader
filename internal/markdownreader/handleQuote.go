package markdownreader

import (
	"strings"
)

// handleQuotes remove the markdown syntax character and style the quote.
func handleQuotes(line string) string {
	var returnLine string

	returnLine = strings.TrimPrefix(line, "> ")
	return colors.Reset() + colors.Dim() + "│ " + returnLine + colors.Reset()
}
