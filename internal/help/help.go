package help

import (
	"fmt"
)

// ShowHelp shows the default help message
func ShowHelp() {
	ShowVersion()
	fmt.Println("https://github.com/zivoxRoot/markdown-reader")
	fmt.Println("")
	fmt.Println("A markdown reader for the terminal")
	fmt.Println("")
	fmt.Println("USAGE:")
	fmt.Println("  md [file]")
	fmt.Println("")
	fmt.Println("OPTIONS:")
	fmt.Println("  -h, --help                show this message and exit")
	fmt.Println("  -v, --version             show the version of markdown-reader")
}

// ShowVersion shows the current markdown-reader's version
func ShowVersion() {
	fmt.Println("markdown reader 0.0.0")
}
