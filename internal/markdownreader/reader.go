package markdownreader

import (
	"fmt"
)

type Reader struct {
	File string
}

func NewReader(file string) *Reader {
	reader := Reader{
		File: file,
	}

	return &reader
}

func (reader *Reader) ProcessMarkdown() ([]string, error) {
	fmt.Println("Proccessing markdown for file : ", reader.File)

	return nil, nil
}
