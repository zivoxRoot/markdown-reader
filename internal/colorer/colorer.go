package colorer

import (
	"os"
)

type Colorer struct {
	ColorMode string
}

func NewColorer() *Colorer {
	colorer := Colorer{}

	colorSupport := getColorSupport()
	colorer.ColorMode = colorSupport

	return &colorer
}

func getColorSupport() string {
	term := os.Getenv("COLORTERM")
	return term
}
