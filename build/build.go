package build

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"swiss/messages"
	"swiss/utils"
)

const buildProgramList = `
Rust: Cargo
C: Clang
Go: Go`

type build struct {
	Language  string
	Tool      string
	Arguments []string
}

type buildRegistry struct {
	builds []build
}

var registry = buildRegistry{
	builds: []build{},
}

func PrintBuildProgramList() {
	messages.Note("Languages are listed with their build tools.")
	fmt.Println(buildProgramList)
}

func (b *buildRegistry) addToBuildRegistry(newBuild build) {
	b.builds = append(b.builds, newBuild)
}

func (b *build) initialize() error {
	command := exec.Command(b.Tool, b.Arguments...)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		return err
	}

	return nil
}

func buildRustProject() build {
	rustBuild := build{
		Language:  "rust",
		Tool:      "cargo",
		Arguments: []string{"build", "--release"},
	}

	return rustBuild
}

func buildGoProject() build {
	goBuild := build{
		Language:  "go",
		Tool:      "go",
		Arguments: []string{"build"},
	}

	return goBuild
}

func buildCProject() build {
	cBuild := build{
		Language:  "c",
		Tool:      "clang",
		Arguments: []string{"main.c", "-Wall", "-Wextra", "-Wpedantic", "-Werror", "-g", "-o", "main"},
	}

	return cBuild
}

func HandleBuildInput(argument string) {
	argument = strings.ToLower(argument)
	for build := range len(registry.builds) {
		if argument == registry.builds[build].Language {
			if err := registry.builds[build].initialize(); err != nil {
				log.Fatal(err)
				return
			} else {
				messages.Success(registry.builds[build].Language + " project has been compiled.")
				return
			}
		}
	}
	messages.Error("Unable to find " + argument + " in registry list.")
}

func SwissInstall() {
	system := utils.GetOperatingSystem()
	if system == "linux" {
		// Builds Go program
		registry.builds[0].initialize()

		command := exec.Command("mv", "swiss", "/home/"+utils.GetUsersName()+"/.local/bin/")

		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		if err := command.Run(); err != nil {
			messages.Error("Swiss install failed. Check output above.")
			return
		}

		messages.Success("Swiss installed successfully! Use 'swiss' in the terminal to gain access to the program.")
	} else {
		messages.Warning("Swiss install is not supported for " + system + ".")
		return
	}
}

func UpdateSwiss() {
	// Make directory to clone into
	utils.MakeFolder("swiss_install", true)
	
	// Clone the repository
	clone := exec.Command("git", "clone", "https://github.com/Tyy47/Swiss.git", "swiss_install/")

	if err := clone.Run(); err != nil {
		messages.Error("Unable to clone Swiss repo. Install manually or create a bug report on the repository.")
		fmt.Println(err)
		return
	}

	// Change directory into cloned repo
	if err := os.Chdir("swiss_install"); err != nil {
		messages.Error("Unable to change directory into swiss_install. Exiting.")
		fmt.Println(err)
		return
	}

	// Prompt the user to either go install or move to local/bin
	messages.Note("Select the number associated with the option in order to continue.")
	fmt.Println("How would you like to install Swiss?")
	fmt.Println("1. Go Install\n2. Move to local/bin ( Linux only )")
	for {
		var userInput string
		fmt.Scanln(&userInput)

		switch userInput {
		case "1":
			// Go install here
			install := exec.Command("go", "install")

			if err := install.Run(); err != nil {
				messages.Error("Unable to install Swiss using Go Install.")
				fmt.Println(err)
				break
			}

			messages.Success("Swiss successfully installed!")
			break

		case "2":
			// Move to bin here
			SwissInstall()
			break
		default:
			// Not a correct option
			messages.Warning("Incorrect option, try again.")
			continue
		}
		break
	}

	messages.Note("Cleaning up install files...")
	if err := os.RemoveAll("./swiss_install"); err != nil {
		messages.Error("Unable to remove install files.")
		fmt.Println(err)
		return
	}
	messages.Success("Install files cleaned up and Swiss is installed!")
}

func init() {
	registry.builds = append(registry.builds, buildGoProject())
	registry.builds = append(registry.builds, buildRustProject())
	registry.builds = append(registry.builds, buildCProject())
}
