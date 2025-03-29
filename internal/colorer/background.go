package colorer

func (colorer *Colorer) BgBlack() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;0;0;0m"
	case "256":
		return "\033[48;5;0m"
	default:
		return "\033[40m"
	}
}

func (colorer *Colorer) BgRed() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;255;0;0m"
	case "256":
		return "\033[48;5;196m"
	default:
		return "\033[41m"
	}
}

func (colorer *Colorer) BgGreen() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;0;255;0m"
	case "256":
		return "\033[48;5;46m"
	default:
		return "\033[42m"
	}
}

func (colorer *Colorer) BgYellow() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;255;255;0m"
	case "256":
		return "\033[48;5;226m"
	default:
		return "\033[43m"
	}
}

func (colorer *Colorer) BgBlue() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;0;0;255m"
	case "256":
		return "\033[48;5;21m"
	default:
		return "\033[44m"
	}
}

func (colorer *Colorer) BgMagenta() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;255;0;255m"
	case "256":
		return "\033[48;5;201m"
	default:
		return "\033[45m"
	}
}

func (colorer *Colorer) BgCyan() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;0;255;255m"
	case "256":
		return "\033[48;5;51m"
	default:
		return "\033[46m"
	}
}

func (colorer *Colorer) BgWhite() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;255;255;255m"
	case "256":
		return "\033[48;5;15m"
	default:
		return "\033[47m"
	}
}

// Bright Background Colors
func (colorer *Colorer) BrightBgBlack() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;85;85;85m"
	case "256":
		return "\033[48;5;8m"
	default:
		return "\033[100m"
	}
}

func (colorer *Colorer) BrightBgRed() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;255;85;85m"
	case "256":
		return "\033[48;5;9m"
	default:
		return "\033[101m"
	}
}

func (colorer *Colorer) BrightBgGreen() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;85;255;85m"
	case "256":
		return "\033[48;5;10m"
	default:
		return "\033[102m"
	}
}

func (colorer *Colorer) BrightBgYellow() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;255;255;85m"
	case "256":
		return "\033[48;5;11m"
	default:
		return "\033[103m"
	}
}

func (colorer *Colorer) BrightBgBlue() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;85;85;255m"
	case "256":
		return "\033[48;5;12m"
	default:
		return "\033[104m"
	}
}

func (colorer *Colorer) BrightBgMagenta() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;255;85;255m"
	case "256":
		return "\033[48;5;13m"
	default:
		return "\033[105m"
	}
}

func (colorer *Colorer) BrightBgCyan() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;85;255;255m"
	case "256":
		return "\033[48;5;14m"
	default:
		return "\033[106m"
	}
}

func (colorer *Colorer) BrightBgWhite() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[48;2;240;240;240m"
	case "256":
		return "\033[48;5;231m"
	default:
		return "\033[107m"
	}
}
