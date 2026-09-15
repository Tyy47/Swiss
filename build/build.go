package build

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"swiss/utils"

	"github.com/Tyy47/clibox/argbin"
)

const buildProgramList = `Rust: Cargo
C: Clang
Go: Go
Zig: Zig`

type build struct {
	Language  string
	Tool      string
	Arguments []string
	BuildFile string // File thats associated with a certain language I.E main.go, Cargo.toml, etc.
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
		BuildFile: "Cargo.toml",
	}

	return rustBuild
}

func buildGoProject() build {
	goBuild := build{
		Language:  "go",
		Tool:      "go",
		Arguments: []string{"build"},
		BuildFile: "main.go",
	}

	return goBuild
}

func buildCProject() build {
	cBuild := build{
		Language:  "c",
		Tool:      "clang",
		Arguments: []string{"main.c", "-Wall", "-Wextra", "-Wpedantic", "-Werror", "-g", "-o", "main"},
		BuildFile: "main.c",
	}

	return cBuild
}

func buildZigProject() build {
	zigBuild := build{
		Language:  "zig",
		Tool:      "zig",
		Arguments: []string{"build"},
		BuildFile: "main.zig",
	}

	return zigBuild
}

func scanForBuildFiles() (bool, build) {
	// Scan directory for all files.
	files, err := os.ReadDir(".")
	if err != nil {
		utils.Error("Unable to read files in current directory.")
		return false, build{}
	}

	for _, project := range registry.builds {
		for _, file := range files {
			if file.Name() == project.BuildFile {
				return true, project
			}
		}
	}

	utils.Warning("Unable to find inputted language, check language list for buildable languages via Swiss.")
	return false, build{}
}

func BuildProject() {
	// Grabs the bool and build struct from scanForFiles()
	result, project := scanForBuildFiles()

	if !result {
		utils.Error("Unable to build project, check inputted language to see if it's in Swiss build list.")
		return
	}

	project.initialize()
	utils.Success(project.Language + " project has been compiled.")
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

func init() {
	registry.builds = append(registry.builds, buildGoProject())
	registry.builds = append(registry.builds, buildRustProject())
	registry.builds = append(registry.builds, buildCProject())
	registry.builds = append(registry.builds, buildZigProject())
}
