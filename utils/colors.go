package utils

import (
	"errors"
	"fmt"
	"strings"
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
type color struct{}

// Colors object to convert strings into colored text
var Colors = color{}

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

// Prints the given text in the color provided.
// Valid Colors are: "black", "red", "green", "yellow", "purple", "cyan", "white" 
func (c *color) PrintColor(text string, colorSelection string) error {
	// Lowercases the string to make input unified
	loweredString := strings.ToLower(colorSelection)

	// Checks if the color is valid using the lowercased string, Crashes the program if an invalid color is inputted.
	if err := c.checkColor(loweredString); err != nil {
		CrashCheck(err)
		return err
	}
	
	// Prints the colored text based on the selections
	switch colorSelection {
	case "red":
		fmt.Println(c.Red(text))
	case "green":
		fmt.Println(c.Green(text))
	case "yellow":
		fmt.Println(c.Yellow(text))
	case "purple":
		fmt.Println(c.Purple(text))
	case "cyan":
		fmt.Println(c.Cyan(text))
	case "white":
		fmt.Println(c.White(text))
	case "black":
		fmt.Println(c.Black(text))
	default:
		fmt.Println("MISSING COLOR")
	}

	return nil
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
