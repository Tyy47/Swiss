package build

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
	"swiss/messages"
	"swiss/utils"
)

// Lists of languages
var supportedLanguages = [...]string{"c", "rust", "go"}
var buildLangsList = [...]string{"Rust: Cargo", "C: Clang", "Go: Go", "Typescript: tsc"}
var runLangsList = [...]string{"Rust: Cargo", "Go: Go"}


func PrintSupportedBuildLanguages() {
	messages.Note("Listings are laid out with their languages and build tools.")
	for i := range len(buildLangsList) {
		fmt.Println(buildLangsList[i])
	}
}

func PrintSupportedRunLanguages() {
	messages.Note("Listings are laid out with their languages and run tools.")
	for i := range len(runLangsList) {
		fmt.Println(runLangsList[i])
	}
}

func CheckLanguageAndBuild(language string) {
	lowercasedLanguage := strings.ToLower(language)
	// Iterates over supported languages and checks if input is supported
	var languageCheck bool = slices.Contains(supportedLanguages[:], lowercasedLanguage)
	if languageCheck {
		switch lowercasedLanguage {
		case "rust":
			rustBuild()
		case "go":
			goBuild(false)
		case "c":
			cBuild()
		default:
			messages.Warning("Unable to find language in supported languages list. Check the build list to view supported languages via Swiss.")
		}
	} else {
		// Prints an error and returns if inputted language is not in supported list.
		messages.Error("Inputted language is not supported via Swiss. Check the build list via swiss build -l")
		return
	}
}


func CheckLanguageAndRun(language string) {
	lowercasedLanguage := strings.ToLower(language)
	// Iterates over supported languages and checks if input is supported
	var languageCheck bool = slices.Contains(supportedLanguages[:], lowercasedLanguage)
	if languageCheck {
		switch lowercasedLanguage {
		case "rust":
			rustRun()
		case "go":
			goRun()
		default:
			messages.Warning("Unable to find language in supported languages list. Check the build list to view supported languages via Swiss.")
		}
	} else {
		// Prints an error and returns if inputted language is not in supported list.
		messages.Error("Inputted language is not supported via Swiss. Check the build list via swiss build -l")
		return
	}
}

func SwissInstall() {
	goBuild(true)
	command := exec.Command("mv", "swiss", "/home/" + utils.GetUsersName() + "/.local/bin/")

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr


	if err := command.Run(); err != nil {
		messages.Error("Swiss install failed. Check output above.")
		return
	}
	
	messages.Success("Swiss installed successfully! Use 'swiss' in the terminal to gain access to the program.")
}

// Build related functions
func rustBuild() {
	command := exec.Command("cargo", "build", "--release")

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		messages.Error("Cargo build failed. Check output above.")
		return
	}
	
	messages.Success("Rust program successfully built! Check target/release to view the binary for your program.")
}

func goBuild(muted bool) {
	command := exec.Command("go", "build")

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		messages.Error("Go build failed. Check output above.")
		return
	}

	if muted {
		return
	} else {
		messages.Success("Go program successfully built! Check local folder to view the binary for your program.")
	}
}

func cBuild() {
	command := exec.Command("clang", "main.c", "-Wall", "-Wextra", "-Wpedantic", "-Werror", "-g", "-o", "main")

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		messages.Error("Clang build failed. Check output above.")
		return
	}
	
	messages.Success("Clang program successfully built! Check local folder to view the binary for your program.")
}

// Run related functions

func rustRun() {
	command := exec.Command("cargo", "run")

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		messages.Error("Cargo run failed. Check output above.")
		return
	}
	
	messages.Success("Rust program successfully ran!")
}

func goRun() {
	command := exec.Command("go", "run", "main.go")

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		messages.Error("Go run failed. Check output above.")
		return
	}
	
	messages.Success("Go program successfully ran!")
}
