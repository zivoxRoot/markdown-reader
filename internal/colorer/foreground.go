package colorer

func (colorer *Colorer) Black() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;0;0;0m"
	case "256":
		return "\033[38;5;0m"
	default:
		return "\033[30m"
	}
}

func (colorer *Colorer) Red() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;255;0;0m"
	case "256":
		return "\033[38;5;196m"
	default:
		return "\033[31m"
	}
}

func (colorer *Colorer) Green() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;0;255;0m"
	case "256":
		return "\033[38;5;46m"
	default:
		return "\033[32m"
	}
}

func (colorer *Colorer) Yellow() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;255;255;0m"
	case "256":
		return "\033[38;5;226m"
	default:
		return "\033[33m"
	}
}

func (colorer *Colorer) Blue() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;0;0;255m"
	case "256":
		return "\033[38;5;21m"
	default:
		return "\033[34m"
	}
}

func (colorer *Colorer) Magenta() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;255;0;255m"
	case "256":
		return "\033[38;5;201m"
	default:
		return "\033[35m"
	}
}

func (colorer *Colorer) Cyan() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;0;255;255m"
	case "256":
		return "\033[38;5;51m"
	default:
		return "\033[36m"
	}
}

func (colorer *Colorer) White() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;255;255;255m"
	case "256":
		return "\033[38;5;15m"
	default:
		return "\033[37m"
	}
}

// Bright Foreground Colors
func (colorer *Colorer) BrightBlack() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;85;85;85m"
	case "256":
		return "\033[38;5;8m"
	default:
		return "\033[90m"
	}
}

func (colorer *Colorer) BrightRed() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;255;85;85m"
	case "256":
		return "\033[38;5;9m"
	default:
		return "\033[91m"
	}
}

func (colorer *Colorer) BrightGreen() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;85;255;85m"
	case "256":
		return "\033[38;5;10m"
	default:
		return "\033[92m"
	}
}

func (colorer *Colorer) BrightYellow() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;255;255;85m"
	case "256":
		return "\033[38;5;11m"
	default:
		return "\033[93m"
	}
}

func (colorer *Colorer) BrightBlue() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;85;85;255m"
	case "256":
		return "\033[38;5;12m"
	default:
		return "\033[94m"
	}
}

func (colorer *Colorer) BrightMagenta() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;255;85;255m"
	case "256":
		return "\033[38;5;13m"
	default:
		return "\033[95m"
	}
}

func (colorer *Colorer) BrightCyan() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;85;255;255m"
	case "256":
		return "\033[38;5;14m"
	default:
		return "\033[96m"
	}
}

func (colorer *Colorer) BrightWhite() string {
	switch colorer.ColorMode {
	case "truecolor":
		return "\033[38;2;240;240;240m"
	case "256":
		return "\033[38;5;231m"
	default:
		return "\033[97m"
	}
}
