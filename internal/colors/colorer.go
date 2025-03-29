package colors

import (
	"os"
)

type Colorer struct {
	ColorMode string
}

func NewColorer() *Colorer {
	colorizer := Colorer{}

	colorSupport := getColorSupport()
	colorizer.ColorMode = colorSupport

	return &colorizer
}

func getColorSupport() string {
	term := os.Getenv("COLORTERM")
	return term
}
