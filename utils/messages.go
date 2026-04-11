package utils

import (
	"fmt"
)

func Error(statement string) {
	fmt.Println( ColorRed + "error" + ColorReset + ": " + statement)
}

func Warning(statement string) {
	fmt.Println(ColorYellow + "warning" + ColorReset + ": " + statement)
}

func Success(statement string) {
	fmt.Println(ColorGreen + "success" + ColorReset + ": " + statement)
}

func Note(statement string) {
	fmt.Println(ColorCyan + "note" + ColorReset + ": " + statement)
}

// Function takes in an error and displays the error message in a formatted way.
func CrashMessage(err error) {
	fmt.Println(ColorPurple + "crash reason" + ColorReset + ": " + err.Error())
}
