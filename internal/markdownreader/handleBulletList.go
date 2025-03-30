package markdownreader

import (
	"strings"
)

// handleBulletList removes the '-', '+', or '*' and output a stylized bullet point.
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

	returnLine = colors.Reset() + "•" + returnLine
	return returnLine
}
