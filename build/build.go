package build

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/Tyy47/clibox/argbin"
)

var (
	// generic errors
	ErrUnknownLanguage = errors.New("unknown language")

	// build object errors

	ErrNilLang = errors.New("lang argument cannot be nil")
	ErrNilBuildTool = errors.New("build tool argument cannot be nil")
	ErrNilBuildArgs = errors.New("build args argument cannot be nil")
	ErrNilBuildFile = errors.New("build file argument cannot be nil")
)

type Language string

const (
	Go Language = "go"
	Rust Language = "rust"
	Python Language = "python"
)

type build struct {
	Tool      string // Tool stores the build tool of a language (e.g cargo, bun, etc)
	BuildArguments []string // BuildArguments stores optional flags for the build Tool
	BuildFile string // File thats associated with a certain language I.E main.go, Cargo.toml, etc.
}

// buildMap contains all of the available languages that can be built via swiss
var buildMap = map[Language]build{
	Go : {
		Tool: "go",
		BuildArguments: []string{"build"},
		BuildFile: "go.mod",
	},

	Rust : {
		Tool: "cargo",
		BuildArguments: []string{"build"},
		BuildFile: "Cargo.toml",
	},
}

// getLanguage takes a language input and returns a build object thats related to that language.
// If no languages are found, ErrUnknownLanguage is returned.
func getLanguage(lang string) (*build, error) {
	fixed := strings.ToLower(lang)
	buildObject, ok := buildMap[Language(fixed)]
	if !ok {
		return nil, ErrUnknownLanguage
	}

	return &buildObject, nil
}

// buildLanguage handles the execution of the build. 
func buildLanguage(b *build) error {

	// Takes in the tool and arguments to create a command
	cmd := exec.Command(b.Tool, b.BuildArguments...)
	
	// Runs the command and returns the error if it fails
	if err := cmd.Run(); err != nil {
		return err
	}
	
	return nil
}

// BuildCommand creates the "build" command for swiss. 
//
// Command is added in main.go in main function.
func BuildCommand() *argbin.Command {
	return &argbin.Command{
		Name: "build",
		Description: "build tool for swiss",
		TakesValue: true,
		Execute: func(ctx *argbin.Context) error {
			
			b, err := getLanguage(ctx.ParsedValue)
			if err != nil {
				return err
			}

			if err := buildLanguage(b); err != nil {
				return err
			}

			return nil
		},
	}
}
