package utils

import (
	"fmt"
)

// Prints an error message to the terminal
func Error(statement string) {
	fmt.Println(ColorRed + "error" + ColorReset + ": " + statement)
}

// Prints a warning message to the terminal
func Warning(statement string) {
	fmt.Println(ColorYellow + "warning" + ColorReset + ": " + statement)
}

// Prints a success message to the terminal
func Success(statement string) {
	fmt.Println(ColorGreen + "success" + ColorReset + ": " + statement)
}

// Prints a note message to the terminal
func Note(statement string) {
	fmt.Println(ColorCyan + "note" + ColorReset + ": " + statement)
}

// Prints an "output" message to the terminal.
// This output message is used if a function produces something.
func Output(statement string) {
	fmt.Println(ColorWhite + "output" + ColorReset + ": " + statement)
}

// Prints a reason message. This function is mainly used in the Crash function to provide a crash reason to the user.
func Reason(statement string) {
	fmt.Println(ColorWhite + "reason" + ColorReset + ": " + statement)
}

// Function takes in an error and displays the error message in a formatted way.
func crashMessage(err error) {
	fmt.Println(ColorPurple + "crash reason" + ColorReset + ": " + err.Error())
}
