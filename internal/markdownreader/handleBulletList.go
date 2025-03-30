package markdownreader

import (
	"strings"
)

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
