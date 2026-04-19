package build

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"swiss/utils"
)

const buildProgramList = `Rust: Cargo
C: Clang
Go: Go
Zig: Zig`

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
	utils.Note("Languages are listed with their build tools.")
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

func buildZigProject() build {
	zigBuild := build{
		Language: "zig",
		Tool: "zig",
		Arguments: []string{"build"},
	}

	return zigBuild
}

func HandleBuildInput() {
	if len(utils.Arguments) < 3 {
		return
	}

	argument := utils.Arguments[2]
	argument = strings.ToLower(argument)

	for build := range len(registry.builds) {
		if argument == registry.builds[build].Language {
			if err := registry.builds[build].initialize(); err != nil {
				utils.Crash(err)
				return
			} else {
				utils.Success(registry.builds[build].Language + " project has been compiled.")
				return
			}
		}
	}
	utils.Error("Unable to find " + argument + " in registry list.")
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
			utils.Error("Swiss install failed. Check output above.")
			return
		}

		utils.Success("Swiss installed successfully! Use 'swiss' in the terminal to gain access to the program.")
	} else {
		utils.Warning("Swiss install is not supported for " + system + ".")
		return
	}
}

func UpdateSwiss() {
	// Make directory to clone into
	utils.MakeFolder("swiss_install", true)
	
	// Clone the repository
	clone := exec.Command("git", "clone", "https://github.com/Tyy47/Swiss.git", "swiss_install/")

	if err := clone.Run(); err != nil {
		utils.Error("Unable to clone Swiss repo. Install manually or create a bug report on the repository.")
		utils.Crash(err)
		return
	}

	// Change directory into cloned repo
	if err := os.Chdir("swiss_install"); err != nil {
		utils.Error("Unable to change directory into swiss_install. Exiting.")
		utils.Crash(err)
		return
	}

	// Prompt the user to either go install or move to local/bin
	utils.Note("Select the number associated with the option in order to continue.")
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
				utils.Error("Unable to install Swiss using Go Install.")
				utils.Crash(err)
				break
			}

			utils.Success("Swiss successfully installed!")
		case "2":
			// Move to bin here
			SwissInstall()
		default:
			// Not a correct option
			utils.Warning("Incorrect option, try again.")
			continue
		}
		break
	}

	utils.Note("Cleaning up install files...")
	if err := os.Chdir(".."); err != nil {
		utils.Error("Unable to change directory.")
		utils.Crash(err)
		return
	}


	if err := os.RemoveAll("swiss_install"); err != nil {
		utils.Error("Unable to remove install files.")
		utils.Crash(err)
		return
	}
	utils.Success("Install files cleaned up and Swiss is installed!")
}

func init() {
	registry.builds = append(registry.builds, buildGoProject())
	registry.builds = append(registry.builds, buildRustProject())
	registry.builds = append(registry.builds, buildCProject())
	registry.builds = append(registry.builds, buildZigProject())
}
