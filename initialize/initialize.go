package initialize

import (
	"fmt"
	"log"
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
	if proj.Name != "" {
		os.Chdir("./" + proj.Name)
		defer os.Chdir("..")
	} else {
		// Make gitignore before adding files to repo so it can get added.
		utils.MakeFile(".gitignore", false)
		utils.MakeFile("TODO.md", false)
		utils.MakeFile("README.md", false)
	}

	init := exec.Command("git", "init")

	if err := init.Run(); err != nil {
		utils.Error("Unable to init git")
		return err
	}

	add := exec.Command("git", "add", ".")

	if err := add.Run(); err != nil {
		utils.Error("Unable to add files to git project")
		return err
	}

	mainBranch := exec.Command("git", "branch", "-M", "main")

	if err := mainBranch.Run(); err != nil {
		utils.Error("Unable to change main branch to 'main'.")
		return err
	}

	utils.Success("Git has been initialized.")

	return nil
}

// Inits jujutsu in current directory when called.
func jjInit() error {
	init := exec.Command("jj", "git", "init")

	if err := init.Run(); err != nil {
		utils.Error("Unable to init JJ")
		return err
	}

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

func getWebProject() project {
	programName := utils.GetUserInput("Enter your project name: ", "my-app")

	program := project{
		Name: programName,
		Language:   "web",
		Tool:       "bun",
		Arguments:  []string{"create", "vite", programName, "--template"},
		ManualInit: false,
	}

	for _, argument := range utils.AdditionalArguments {
		switch argument {
		case "react":
			program.Arguments = append(program.Arguments, "react-ts")
		case "sv", "svelte":
			program.Arguments = append(program.Arguments, "svelte-ts")
		case "angular":
			program.Arguments = append(program.Arguments, "angular-ts")
		case "vue":
			program.Arguments = append(program.Arguments, "vue-ts")
		default:
			program.Arguments = append(program.Arguments, "vanilla-ts")
		}
	}

	return program
}

func CreateWebProject() {
	project := getWebProject()

	if err := project.initialize(); err != nil {
		utils.Crash(err)
		return
	}

	flagHandler(&utils.AdditionalArguments, project)
	utils.Success("web project has been created.")
}

func CreateProject() {
	if len(utils.Arguments) < 3 {
		return
	}

	argument := strings.ToLower(utils.Arguments[2])
	for project := range len(registry.projects) {
		if argument == registry.projects[project].Language {
			if err := registry.projects[project].initialize(); err != nil {
				log.Fatal(err)
				return
			} else {
				flagHandler(&utils.AdditionalArguments, registry.projects[project])
				utils.Success(registry.projects[project].Language + " project has been created.")
				return

			}
		}
	}
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
