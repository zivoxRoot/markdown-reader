package markdownreader

import (
	"fmt"
	"strings"
	"io/ioutil"
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

	// Check that the Reader has a valid File field
	trimedFileName := strings.Trim(reader.File, " ")
	if trimedFileName == "" {
		return nil, fmt.Errorf("the markdown reader needs a file property")
	}

	// Read the file
	content, err := ioutil.ReadFile(reader.File)
	if err != nil {
		return nil, fmt.Errorf("error reading file %v : %v", reader.File, err)
	}

	// Convert the file to a slice
	lines := strings.Split(string(content), "\n")
	
	var returnValue []string

	for _, line := range(lines) {
		line := handlePrefix(line)
		returnValue = append(returnValue, line)
	}

	return returnValue, nil
}
