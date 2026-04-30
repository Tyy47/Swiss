package utils

// Colors const array
const (
	// High intensity bold colors
	colorBlack  = "\033[1;90m"
	colorRed    = "\033[1;91m"
	colorGreen  = "\033[1;92m"
	colorYellow = "\033[1;93m"
	colorPurple = "\033[1;95m"
	colorCyan   = "\033[1;96m"
	colorWhite  = "\033[1;97m"
	colorReset  = "\033[0m"
)

// Create the color type and create the struct for use in other packages.
type color struct {}
var Colors = color{}

func (c *color) Red(word string) string {
	return colorRed + word + colorReset
}

func (c *color) Green(word string) string {
	return colorGreen + word + colorReset
}

func (c *color) Yellow(word string) string {
	return colorYellow + word + colorReset
}

func (c *color) Purple(word string) string {
	return colorPurple + word + colorReset
}

func (c *color) Cyan(word string) string {
	return colorCyan + word + colorReset
}

func (c *color) White(word string) string {
	return colorWhite + word + colorReset
}
