package initialize

import (
	"errors"
	"os"
	"os/exec"
	"strings"

	"swiss/utils"

	"github.com/Tyy47/clibox/argbin"
	"github.com/Tyy47/clibox/inputbin"
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

<<<<<<< HEAD
// Handles additional flags that might be tossed into the init command when ran to execute additional functions.
func flagHandler(additionalArgs *[]string, proj project) {
	for _, arg := range *additionalArgs {
		switch arg {
		case "-g", "--git":
			gitInit(proj)
		case "-j", "--jujutsu":
			jjInit()
		}
=======
	if proj.NeedsProjectName {
	
		ops := inputbin.InputOptions{
			Question: "enter project name: ",
		}

		// Gather user input
		userInput, err := inputbin.Text(&ops)
		if err != nil {
			// If input is blank it will get current working directory name
			userInput, err = func() (string, error) {
				// Grab working directory
				wd, _ := os.Getwd()
				
				// Split the directory string
				split := strings.Split(wd, "/")

				// Grab last entry
				if len(split) == 1 {
					return split[0], nil
				}

				return split[len(split)-1], nil
			}()
			
			// Defaults the project name to "project" if name cannot be gotten elsewhere
			if err != nil {
				userInput = "project"
			}
		}
		
		// Adds project name to arguments list
		proj.Arguments = append(proj.Arguments, userInput)
>>>>>>> cleanup
	}
	
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

<<<<<<< HEAD
// Inits git in current directory when called.
func gitInit(proj project) error {
	// Changes directory into init'd project folder
	if proj.Name != "" {
		if err := os.Chdir("./" + proj.Name); err != nil {
			utils.Error("Unable to change to project directory: " + err.Error())
			return err
		}
		defer os.Chdir("..") // Changes back to original directory when function finishes
	} else {
		// Creates all the files that are usually in a repository
		utils.MakeFile(".gitignore", false)
		utils.MakeFile("TODO.md", false)
		utils.MakeFile("README.md", false)
	}

=======
// gitInit runs "git init" in the current directory
func gitInit() error {
>>>>>>> cleanup
	// Initing git command
	init := exec.Command("git", "init")

	// Runs git init and returns an error if unable to
	if err := init.Run(); err != nil {
		return err
	}

<<<<<<< HEAD
	// Command to add all files to the repository
	add := exec.Command("git", "add", ".")

	// Runs the add command and returns the error if unsuccessful
	if err := add.Run(); err != nil {
		utils.Error("Unable to add files to git project")
		utils.Reason(err.Error())
		return err
	}

	// Command to create an initial commit message.
	commit := exec.Command("git", "commit", "-m", "initial")

	if err := commit.Run(); err != nil {
		utils.Error("Unable to create an initial commit message")
		utils.Reason(err.Error())
		return err
	}

	// Renames the master branch to "main" command
	mainBranch := exec.Command("git", "branch", "-M", "main")

	// Runs the rename command and returns the err if unsuccessful
	if err := mainBranch.Run(); err != nil {
		utils.Error("Unable to change main branch to 'main'.")
		utils.Reason(err.Error())
		return err
	}

	// Adds the remote repository link if the user provides one.
	remoteURL := utils.GetUserInput("Enter remote repository url (leave blank to skip): ", "")
	if remoteURL == "" {
		utils.Note("No remote URL provided, skipping remote add and push.")
		utils.Success("Git has been initialized.")
		return nil
	}

	remoteAdd := exec.Command("git", "remote", "add", "origin", remoteURL)
	if err := remoteAdd.Run(); err != nil {
		utils.Error("Unable to add the remote repository url.")
		utils.Reason(err.Error())
		return err
	}

	// Pushes the initial commit to the remote repository
	utils.Success("Pushing to remote repository...")
	push := exec.Command("git", "push", "-u", "origin", "main")
	if err := push.Run(); err != nil {
		utils.Error("Unable to push initial commit to remote repository.")
		utils.Reason(err.Error())
		return err
	}

	// Success message when finished
	utils.Success("Git has been initialized.")

=======
>>>>>>> cleanup
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

<<<<<<< HEAD
func createRustProject() project {
	return project{
		Language:   "rust",
		Tool:       "cargo",
		Arguments:  []string{"init"},
		ManualInit: false,
	}
}

func createGoProject() project {
	args := utils.Arguments
	if len(args) >= 4 {
		return project{
			Language:   "go",
			Tool:       "go",
			Arguments:  []string{"mod", "init", args[3]},
			Files:      []string{"main.go"},
			ManualInit: false,
		}
	} else {
		return project{
			Language:   "go",
			Tool:       "go",
			Arguments:  []string{"mod", "init", "project"},
			Files:      []string{"main.go"},
			ManualInit: false,
		}
	}
}

func createCProject() project {
	return project{
		Language:   "c",
		Tool:       "clang",
		Folders:    []string{"src"},
		Files:      []string{"main.c"},
		ManualInit: true,
	}
}

func createHTMLProject() project {
	return project{
		Language:   "html",
		Tool:       "html",
		Folders:    []string{},
		Files:      []string{"TODO.md", "index.html", "styles.css", "main.js"},
		ManualInit: true,
	}
}

func createZigProject() project {
	return project{
		Language:   "zig",
		Tool:       "zig",
		Arguments:  []string{"init"},
		ManualInit: false,
	}
}

func createPythonProject() project {
	return project{
		Language:   "python",
		Tool:       "uv",
		Arguments:  []string{"init"},
		ManualInit: false,
	}
}

func createTSProject() project {
	return project{
		Language:   "ts",
		Tool:       "bun",
		Arguments:  []string{"init", "--yes"},
		ManualInit: false,
	}
}

// Creates a web based project. Can be modified by users arguments for differing frameworks
func getWebProject() (project, string) {
	programName := utils.GetUserInput("Enter your project name: ", "my-app")
	var framework string // Allows to be more specific in success messages.

	program := project{
		Name:       programName,
		Language:   "web",
		Tool:       "bun",
		Arguments:  []string{"create", "vite", programName, "--template"},
		ManualInit: false,
	}

	var webFramework string
	if len(utils.AdditionalArguments) > 0 {
		webFramework = utils.AdditionalArguments[0]
	}
	switch webFramework {
	case "react":
		program.Arguments = append(program.Arguments, "react-ts")
		framework = "react"
	case "sv", "svelte":
		program.Arguments = append(program.Arguments, "svelte-ts")
		framework = "svelte"
	case "angular":
		program.Arguments = append(program.Arguments, "angular-ts")
		framework = "angular"
	case "vue":
		program.Arguments = append(program.Arguments, "vue-ts")
		framework = "vue"
	default:
		program.Arguments = append(program.Arguments, "vanilla-ts")
		framework = "web"
	}

	return program, framework
}

func CreateWebProject() {
	// Grabs the project that is the result of the getWebProject function
	project, framework := getWebProject()

	// Attempts to initialize the web project, if it fails, it will crash.
	if err := project.initialize(); err != nil {
		utils.CrashCheck(err)
		return
	}

	// Executes the flag handler for any additional arguments that are providied
	flagHandler(&utils.AdditionalArguments, project)

	// Checks if the framework contains anything other then web, then changes the success message.
	if framework != "web" {
		utils.Success(framework + " project has been created.")
		return
	}

	utils.Success("web project has been created.")
}

func CreateProject() {
	// Length check for later code
	if len(utils.Arguments) < 3 {
		return
	}

	// Grabs the name of the project the user init's
	argument := strings.ToLower(utils.Arguments[2])
	// Loops through projects in the registry
	for project := range len(registry.projects) {
		// Checks if an argument is in the products registry
		if argument == registry.projects[project].Language {
			// Runs the initialize method of the project if the project has been found, if it's an error, it checks if it's a crash.
			if err := registry.projects[project].initialize(); err != nil {
				utils.CrashCheck(err)
				return
			} else {
				// Executes all the additional flags if any are provided by the user.
				flagHandler(&utils.AdditionalArguments, registry.projects[project])
				utils.Success(registry.projects[project].Language + " project has been created.")
				return
=======
func InitCommand() *argbin.Command{
	return &argbin.Command{
		Name: "init",
		TakesValue: true,
		Execute: func(ctx *argbin.Context) error {
			// Toggle for silent output
			silent, _ := ctx.Values["silent"].(bool)
			
			utils.Output.Infof("initializing %s project", ctx.ParsedValue)
>>>>>>> cleanup

			// Finds the project
			proj, err := initialize(ctx.ParsedValue)
			if err != nil {
				return err
			}
			
			cmd := createExecuteCommand(proj)
			
			// Toggles stdout & stderr based on silent flag
			if !silent {
				// Put cmds output to stdout and stderr
				utils.ToggleOutputForCMD(cmd)
			}
			
			// Execute init command
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
		HelpMenu: `
╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
Init module - Initialize a project using Swiss.

Commands:
	init <string> [name: string]: Inits a project based on the given input. 

Flags:
	-h --help: Opens the help menu.
	-l --list: Prints a list of projects that can be initialized and if they are supported with additional arguments for names.
	-g --git: Inits git alongside your project.
	-j --jujutsu: Inits jj alongside your project.
`,
	}
}
