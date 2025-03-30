package markdownreader

// handleCodeBlock removes the code block characters and colors the text.
func handleCodeBlock() string {
	return colors.Reset() + "" + colors.Dim()
}
