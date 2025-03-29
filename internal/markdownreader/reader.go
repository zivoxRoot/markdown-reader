package markdownreader

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/zivoxRoot/markdown-reader/internal/colorer"
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

// Initialize a new colorer
var colors = colorer.NewColorer()

func (reader *Reader) ProcessMarkdown() ([]string, error) {

	// Check that the Reader has a valid File property
	trimedFileName := strings.Trim(reader.File, " ")
	if trimedFileName == "" {
		return nil, fmt.Errorf("the markdown reader needs a file property")
	}

	// Read the file
	content, err := ioutil.ReadFile(reader.File)
	if err != nil {
		return nil, err
	}

	// Convert the file's content to a slice
	lines := strings.Split(string(content), "\n")

	var returnValue []string

	// Start the return with a newline
	returnValue = append(returnValue, "\n")

	for _, line := range lines {

		// Check the prefix of each line for title, bullet list...
		line := handlePrefix(line)

		// Check the content of each line for bold, italic...
		line = handleContent(line)

		returnValue = append(returnValue, line)
	}

	return returnValue, nil
}
