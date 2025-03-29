package colors

func (colorer *Colorer) Reset() string {
	return "\033[0m"
}

func (colorer *Colorer) Bold() string {
	return "\033[1m"
}

func (colorer *Colorer) Strikethrough() string {
	return "\033[9m"
}

func (colorer *Colorer) Dim() string {
	return "\033[2m"
}

func (colorer *Colorer) Italic() string {
	return "\033[3m"
}

func (colorer *Colorer) Underline() string {
	return "\033[4m"
}
