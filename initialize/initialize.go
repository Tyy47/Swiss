package initialize

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"swiss/utils"

	"github.com/Tyy47/clibox/argbin"
	"github.com/Tyy47/clibox/outbin"
)

var (
	ErrUnknownProject = errors.New("unknown project name")
)

const initProjectList = `Languages:
Rust - Cargo: swiss init rust
Go - Go: swiss init go
C - Swiss: swiss init c
Zig - Zig: swiss init zig
Python - uv: swiss init python
Typescript - bun: swiss init ts`

var output = outbin.NewOutput(os.Stdout, os.Stderr)

// findProject reads through a map of projects. If a project is found,
// It'll return. If not, it will return an error.
func findProject(s string) (*project, error) {
	lower := strings.ToLower(s)

	proj, ok := initMap[Language(lower)] 
	if !ok {
		return nil, ErrUnknownProject
	}

	return &proj, nil
	
}

// createProjectDirectory creates the files and folders required for the project init
func createProjectDirectory(p *project) error {
	
	// Create folders based on the project request
	if p.Folders != nil {
		if err := utils.MakeFolder(p.Folders...); err != nil {
			return err
		}
	}

	// Create files based on the project request
	if p.Files != nil {
		if err := utils.MakeFile(p.Files...); err != nil {
			return err
		}
	}

	return nil
}

// createExecuteCommand builds the command to create the project and returns it as a
// pointer to and exec Cmd.
func createExecuteCommand(proj *project) *exec.Cmd {
	
	// Create the command
	cmd := exec.Command(proj.Tool, proj.Arguments...)

	return cmd
} 


// initialize is the main function on the init process for swiss.
// It grabs the project from findProject and feeds it into createProjectDirectory.
// It returns the project and err if the two functions fail.
func initialize(projectName string) (*project, error) {

	// Finds the project based on the users input 
	proj, err := findProject(projectName)
	if err != nil {
		return nil, err
	}

	// Creates the project folders & files.
	if err := createProjectDirectory(proj); err != nil {
		return nil, err
	}

	return proj, nil
}

// gitInit runs "git init" in the current directory
func gitInit() error {
	// Initing git command
	init := exec.Command("git", "init")

	// Runs git init and returns an error if unable to
	if err := init.Run(); err != nil {
		return err
	}

	return nil
}

// jjInit runs "jj git init" in the current directory
func jjInit() error {
	// Init jujutsu command
	init := exec.Command("jj", "git", "init")

	// Runs the init command and returns the error if unsuccessful
	if err := init.Run(); err != nil {
		return err
	}

	return nil
}

func InitCommand() *argbin.Command{
	return &argbin.Command{
		Name: "init",
		TakesValue: true,
		Execute: func(ctx *argbin.Context) error {
			silent, _ := ctx.Values["silent"].(bool)

			proj, err := initialize(ctx.ParsedValue)
			if err != nil {
				return err
			}
			
			cmd := createExecuteCommand(proj)

			if !silent {
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
			}

			if err := cmd.Run(); err != nil {
				return err
			}

			output.Successf("%s project has been created.", ctx.ParsedValue)
			return nil
		},
		Flags: argbin.Flags{
			"--silent": {
				Execute: func(ctx *argbin.Context) error {
					ctx.Values["silent"] = true
					return nil
				},
			},
			"-s": {
				Execute: func(ctx *argbin.Context) error {
					ctx.Values["silent"] = true
					return nil
				},
			},
			"-g": {
				Execute: func(ctx *argbin.Context) error {
					silent, _ := ctx.Values["silent"].(bool)

					if err := gitInit(); err != nil {
						output.Error("unable to init git")
						return nil
					}

					if !silent {
						output.Success("git initialized successfully")
					}
					
					return nil
				},
			},
			"--git": {
				Execute: func(ctx *argbin.Context) error {
					silent, _ := ctx.Values["silent"].(bool)

					if err := gitInit(); err != nil {
						output.Error("unable to init git")
						return nil
					}

					if !silent {
						output.Success("git initialized successfully")
					}
					
					return nil
				},
			},
			"-j": {
				Execute: func(ctx *argbin.Context) error {
					silent, _ := ctx.Values["silent"].(bool)

					if err := jjInit(); err != nil {
						output.Error("unable to init jujutsu")
						return nil
					}

					if !silent {
						output.Success("jujutsu initialized successfully")
					}
					
					return nil
				},
			},
			"--jujutsu": {
				Execute: func(ctx *argbin.Context) error {
					silent, _ := ctx.Values["silent"].(bool)

					if err := jjInit(); err != nil {
						output.Error("unable to init jujutsu")
						return nil
					}

					if !silent {
						output.Success("jujutsu initialized successfully")
					}
					
					return nil
				},
			},
		},
		Description: `
╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
Init module - Initialize a project using Swiss.

-h --help: Opens the help menu.
-l --list: Prints a list of projects that can be initialized and if they are supported with additional arguments for names.
-g --git: Inits git alongside your project.
-j --jujutsu: Inits jj alongside your project.
init <string> [name: string]: Inits a project based on the given input. 
init web: Inits a web based project using Bun & Vite with a selected framework. 
`,
	}
}
