package initialize

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"swiss/utils"
)

const initProjectList = `Languages:
Rust - Cargo: swiss init rust
Go - Go: swiss init go [module name here]
C - Swiss: swiss init c
HTML - Swiss: swiss init html
Zig - Zig: swiss init zig
Python - uv: swiss init python

Web:
Vanilla TS Web App - Bun/Vite: swiss init web
Svelte - Bun/Vite: swiss init web sv or svelte
React - Bun/Vite: swiss init web react
Angular - Bun/Vite: swiss init web angular
Vue - Bun/Vite: swiss init web vue`

// Project structure for creation
type project struct {
	Name       string // Captures the project name for use in other functions
	Language   string
	Tool       string
	Arguments  []string
	Folders    []string
	Files      []string
	ManualInit bool
}

// Projects storage type
type projectRegistry struct {
	projects []project
}

// Storage for containing valid initable projects
var registry = projectRegistry{
	projects: []project{},
}

// Prints a list of projects that can be init'd via Swiss commands
func PrintInitProjectList() {
	utils.Note("Languages are listed along side their build tools and the commands to init them via Swiss.\n")
	fmt.Println(initProjectList)
}

// Project method that starts the creation of a project
func (p *project) initialize() error {
	// Loops over files in projects structure and creates them if there is any
	for file := range p.Files {
		utils.MakeFile(p.Files[file], false)
	}

	// Loops over folders in projects structure and creates them if there is any
	for folder := range p.Folders {
		utils.MakeFolder(p.Folders[folder], false)
	}
	
	// Checks if a project has to be "manually" init'd. This means that the language of the project thats being initialized has a special setup.
	if p.ManualInit {
		// Switch case statement to grab language and check to see if the files need to moved anywhere after creation
		switch strings.ToLower(p.Language) {
		case "c":
			utils.MoveFileToFolder("./main.c", "./src/main.c", true)
		default:
			return nil
		}
		return nil
	}
	
	// Executes a command using arguments from the Project structure
	command := exec.Command(p.Tool, p.Arguments...)
	
	// Sets the commands standard out and error to the the terminal so it's viewable when something occurs
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	
	// If the command cannot execute, it'll print a failed to init statement and return the error
	if err := command.Run(); err != nil {
		utils.Error(p.Language + " project failed to initialize. Check output below for more details.")
		return err
	}
	
	return nil
}

// Adds projects to the project registry by unpacking a project array
func registerProjects(project ...project) {
	registry.projects = append(registry.projects, project...)
}

// Handles additional flags that might be tossed into the init command when ran to execute additional functions.
func flagHandler(additonalArgs *[]string, proj project) {
	if len(*additonalArgs) >= 1 {
		for _, arg := range *additonalArgs {
			switch arg {
			case "-g", "--git":
				gitInit(proj)
				return
			case "-j", "--jujutsu":
				jjInit()
				return
			}
		}
	} else {
		return
	}
}

// Inits git in current directory when called.
func gitInit(proj project) error {
	// Changes directory into init'd project folder
	if proj.Name != "" {
		os.Chdir("./" + proj.Name)
		defer os.Chdir("..") // Changes back to original directory when function finishes
	} else {
		// Creates all the files that are ususally in a repository
		utils.MakeFile(".gitignore", false)
		utils.MakeFile("TODO.md", false)
		utils.MakeFile("README.md", false)
	}
	
	// Initing git command
	init := exec.Command("git", "init")
	
	// Runs git init and returns an error if unable to 
	if err := init.Run(); err != nil {
		utils.Error("Unable to init git")
		utils.Reason(err.Error())
		return err
	}
	
	// Command to add all files to the repository
	add := exec.Command("git", "add", ".")
	
	// Runs the add command and returns the error if unsuccessful
	if err := add.Run(); err != nil {
		utils.Error("Unable to add files to git project")
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
	
	// Success message when finished
	utils.Success("Git has been initialized.")
	
	return nil
}

// Inits jujutsu in current directory when called.
func jjInit() error {
	// Init jujutsu command
	init := exec.Command("jj", "git", "init")
	
	// Runs the init command and returns the error if unsuccessful
	if err := init.Run(); err != nil {
		utils.Error("Unable to init JJ")
		utils.Reason(err.Error())
		return err
	}
	
	// Success message if jj has been init'd
	utils.Success("Jujutsu has been initialized.")

	return nil
}

func createRustProject() project {
	program := project{
		Language:   "rust",
		Tool:       "cargo",
		Arguments:  []string{"init"},
		ManualInit: false,
	}

	return program
}

func createGoProject() project {
	args := utils.Arguments
	if len(args) >= 4 {
		program := project{
			Language:   "go",
			Tool:       "go",
			Arguments:  []string{"mod", "init", args[3]},
			Files:      []string{"main.go"},
			ManualInit: false,
		}
		return program
	} else {
		program := project{
			Language:   "go",
			Tool:       "go",
			Arguments:  []string{"mod", "init", "project"},
			Files:      []string{"main.go"},
			ManualInit: false,
		}
		return program
	}
}

func createCProject() project {
	program := project{
		Language:   "c",
		Tool:       "clang",
		Folders:    []string{"src"},
		Files:      []string{"main.c"},
		ManualInit: true,
	}

	return program
}

func createHTMLProject() project {
	program := project{
		Language:   "html",
		Tool:       "html",
		Folders:    []string{},
		Files:      []string{"TODO.md", "index.html", "styles.css", "main.js"},
		ManualInit: true,
	}

	return program
}

func createZigProject() project {
	program := project{
		Language:   "zig",
		Tool:       "zig",
		Arguments:  []string{"init"},
		ManualInit: false,
	}

	return program
}

func createPythonProject() project {
	program := project{
		Language:   "python",
		Tool:       "uv",
		Arguments:  []string{"init"},
		ManualInit: false,
	}

	return program
}

// Creates a web based project. Can be modified by users arguments for differing frameworks
func getWebProject() (project, string) {
	programName := utils.GetUserInput("Enter your project name: ", "my-app")
	var framework string // Allows to be more specific in success messages.

	program := project{
		Name: programName,
		Language:   "web",
		Tool:       "bun",
		Arguments:  []string{"create", "vite", programName, "--template"},
		ManualInit: false,
	}

	switch utils.AdditionalArguments[0] {
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
			// Runs the initalize method of the project if the project has been found, if it's an error, it checks if it's a crash.
			if err := registry.projects[project].initialize(); err != nil {
				utils.CrashCheck(err)
				return
			} else {
				// Executes all the additional flags if any are provided by the user.
				flagHandler(&utils.AdditionalArguments, registry.projects[project])
				utils.Success(registry.projects[project].Language + " project has been created.")
				return

			}
		}
	}
	// Message to the user if a language is not found in the projects registry list.
	utils.Error("Unable to find " + argument + " in registry list.")
}

func init() {
	projectArray := []project{
		createGoProject(),
		createRustProject(),
		createCProject(),
		createHTMLProject(),
		createZigProject(),
		createPythonProject(),
	}

	registerProjects(projectArray...)
}
