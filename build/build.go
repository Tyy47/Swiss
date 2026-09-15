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

	ErrNilLang      = errors.New("lang argument cannot be nil")
	ErrNilBuildTool = errors.New("build tool argument cannot be nil")
	ErrNilBuildArgs = errors.New("build args argument cannot be nil")
	ErrNilBuildFile = errors.New("build file argument cannot be nil")

	// program type errors

	ErrNilProgram = errors.New("program return cannot be nil")
	ErrNilBuild = errors.New("build in program type cannot be nil")
	ErrNilRun = errors.New("run in program type cannot be nil")
)

// Language safely adds language guard rails
type Language string

const (
	Go   Language = "go"
	Rust Language = "rust"
)

// build stores all the needed information to build a program
type build struct {
	Tool           string   // Tool stores the build tool of a language (e.g cargo, bun, etc)
	BuildArguments []string // BuildArguments stores optional flags for the build Tool
	BuildFile      string   // File thats associated with a certain language I.E main.go, Cargo.toml, etc.
}

// run stores all the needed information to run a program
type run struct {
	Tool           string   // Tool stores the run tool of a language (e.g cargo, bun, etc)
	BuildArguments []string // BuildArguments stores optional flags for the run Tool
	BuildFile      string   // File thats associated with a certain language I.E main.go, Cargo.toml, etc.
}

// program stores context between functions to get build and run objects for later function executions
type program struct {
	Build *build
	Run *run
}

// buildMap contains all of the available languages that can be built via swiss.
var buildMap = map[Language]build{
	Go: {
		Tool:           "go",
		BuildArguments: []string{"build"},
		BuildFile:      "go.mod",
	},

	Rust: {
		Tool:           "cargo",
		BuildArguments: []string{"build"},
		BuildFile:      "Cargo.toml",
	},
}

// runMap contains all of the available languages that can be ran via swiss
var runMap = map[Language]run{
	Go: {
		Tool:           "go",
		BuildArguments: []string{"build"},
		BuildFile:      "go.mod",
	},

	Rust: {
		Tool:           "cargo",
		BuildArguments: []string{"build"},
		BuildFile:      "Cargo.toml",
	},
}

// getLanguage takes a language input and returns a build object thats related to that language.
// If no languages are found, ErrUnknownLanguage is returned.
func getLanguage(lang string) (*program, error) {
	// Nil check for lang
	if lang == "" {
		return nil, ErrNilLang
	}

	// Lowercase the input
	fixed := strings.ToLower(lang)

	// Gather the build object
	buildObject, ok := buildMap[Language(fixed)]
	if !ok {
		return nil, ErrUnknownLanguage
	}

	// Gather the run object
	runObject, ok := runMap[Language(fixed)]
	if !ok {
		return nil, ErrUnknownLanguage
	}

	// Create a program and assign the build and run objects to store for context
	langs := program{
		Build: &buildObject,
		Run: &runObject,
	}

	// Return the build object
	return &langs, nil
}

// buildLanguage handles the execution of the build of a program.
func buildLanguage(p *program) error {
	// Nil check program
	if p == nil {
		return ErrNilProgram
	}

	// Nil check for build
	if p.Build == nil {
		return ErrNilBuild
	}

	// Assigning build to b
	b := p.Build

	// Takes in the tool and arguments to create a command
	cmd := exec.Command(b.Tool, b.BuildArguments...)

	// Runs the command and returns the error if it fails
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// runLanguage handles the running the program.
func runLanguage(p *program) error {
	// Nil check for program
	if p == nil {
		return ErrNilProgram
	}

	// Nil check for stored build
	if p.Build == nil {
		return ErrNilRun
	}

	// Assigning run to r
	r := p.Run

	// Takes in the tool and arguments to create a command
	cmd := exec.Command(r.Tool, r.BuildArguments...)

	// Runs the command and returns the error if it fails
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// BuildCommand creates the "build" command for swiss.
// Command is added in main.go in main function.
func BuildCommand() *argbin.Command {
	return &argbin.Command{
		Name:        "build",
		Description: "build tool for swiss",
		TakesValue:  true,
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

// RunCommand create the "run" command for swiss.
// Command is added in main.go in main function.
func RunCommand() *argbin.Command {
	return &argbin.Command{
		Name: "run",
		Description: "run tool for swiss",
		TakesValue: true,
		Execute: func(ctx *argbin.Context) error {
			r, err := getLanguage(ctx.ParsedValue)
			if err != nil {
				return err 
			}

			if err := buildLanguage(r); err != nil {
				return err
			}

			return nil
		},
	}
}
