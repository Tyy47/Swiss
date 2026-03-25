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

func init() {
	registry.builds = append(registry.builds, buildGoProject())
	registry.builds = append(registry.builds, buildRustProject())
	registry.builds = append(registry.builds, buildCProject())
}
