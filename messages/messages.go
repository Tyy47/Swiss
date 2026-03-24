package messages

import (
	"fmt"
	"swiss/colors"
)

func Error(statement string) {
	fmt.Println( colors.ColorRed + "error" + colors.ColorReset + ": " + statement)
}

func Warning(statement string) {
	fmt.Println(colors.ColorYellow + "warning" + colors.ColorReset + ": " + statement)
}

func Success(statement string) {
	fmt.Println(colors.ColorGreen + "success" + colors.ColorReset + ": " + statement)
}

func Note(statement string) {
	fmt.Println(colors.ColorCyan + "note" + colors.ColorReset + ": " + statement)
}
