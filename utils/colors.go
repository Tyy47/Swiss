package utils

import (
	"errors"
)

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
type color struct{
	ValidColors []string // Stores all valid colors that can be used
	InvalidColorFound bool // Tracks if an invalid color was called
}

// Colors object to convert strings into colored text
var Colors = color{
	ValidColors: []string{
		"black",
		"red",
		"green",
		"yellow",
		"purple",
		"cyan",
		"white",
	},
	InvalidColorFound: false,
}

// Takes an input of a string to check if its a valid color that can be printed.
// If a valid color is found, it will return true, otherwise it will return false.
func (c *color) checkColor(ColorSelection string) error {
	// Creates a new crash error if the code has an invalid color
	var err error = errors.New("Swiss crash: invalid color selection.")

	// Checks if the input is a valid color selection
	switch ColorSelection {
		case "black", "red", "green", "yellow", "purple", "cyan", "white":
			return nil
		default:
			return err
	}
}

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

// Changes the color of the given string to black
func (c *color) Black(word string) string {
	return colorBlack + word + colorReset
}
