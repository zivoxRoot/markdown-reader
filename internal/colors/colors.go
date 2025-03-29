package colors

// Special styles
func Reset() string {
	return "\033[0m"
}

func Bold() string {
	return "\033[1m"
}

func Dim() string {
	return "\033[2m"
}

func Italic() string {
	return "\033[3m"
}

func Underline() string {
	return "\033[4m"
}

// Colors
func Black() string {
	return "\033[30m"
}

func BrightBlack() string {
	return "\033[90m"
}

func Red() string {
	return "\033[31m"
}

func BrightRed() string {
	return "\033[91m"
}

func Green() string {
	return "\033[32m"
}

func BrightGreen() string {
	return "\033[92m"
}

func Yellow() string {
	return "\033[33m"
}

func BrightYellow() string {
	return "\033[93m"
}

func Blue() string {
	return "\033[34m"
}

func BrightBlue() string {
	return "\033[94m"
}

func Magenta() string {
	return "\033[35m"
}

func BrightMagenta() string {
	return "\033[95m"
}

func Cyan() string {
	return "\033[36m"
}

func BrightCyan() string {
	return "\033[96m"
}

func White() string {
	return "\033[37m"
}

func BrightWhite() string {
	return "\033[97m"
}

// Background colors
func BgBlack() string {
	return "\033[40m"
}

func BrightBgBlack() string {
	return "\033[100m"
}

func BgRed() string {
	return "\033[41m"
}

func BrightBgRed() string {
	return "\033[101m"
}

func BgGreen() string {
	return "\033[42m"
}

func BrightBgGreen() string {
	return "\033[102m"
}

func BgYellow() string {
	return "\033[43m"
}

func BrightBgYellow() string {
	return "\033[103m"
}

func BgBlue() string {
	return "\033[44m"
}

func BrightBgBlue() string {
	return "\033[104m"
}

func BgMagenta() string {
	return "\033[45m"
}

func BrightBgMagenta() string {
	return "\033[105m"
}

func BgCyan() string {
	return "\033[46m"
}

func BrightBgCyan() string {
	return "\033[106m"
}

func BgWhite() string {
	return "\033[47m"
}

func BrightBgWhite() string {
	return "\033[107m"
}
