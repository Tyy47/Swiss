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
type color struct{}

// Colors object to convert strings into colored text
var Colors = color{}

// Changes the color of the given string to red
func (c *color) Red(word string) string {
	return colorRed + word + colorReset
}

// Changes the color of the given string to green
func (c *color) Green(word string) string {
	return colorGreen + word + colorReset
}

// Changes the color of the given string to yellow
func (c *color) Yellow(word string) string {
	return colorYellow + word + colorReset
}

// Changes the color of the given string to purple
func (c *color) Purple(word string) string {
	return colorPurple + word + colorReset
}

// Changes the color of the given string to cyan
func (c *color) Cyan(word string) string {
	return colorCyan + word + colorReset
}

// Changes the color of the given string to white
func (c *color) White(word string) string {
	return colorWhite + word + colorReset
}
