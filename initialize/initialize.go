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

type project struct {
	Language   string
	Tool       string
	Arguments  []string
	Folders    []string
	Files      []string
	ManualInit bool
}

type projectRegistry struct {
	projects []project
}

var registry = projectRegistry{
	projects: []project{},
}

func PrintInitProjectList() {
	fmt.Println(initProjectList)
}

func (p *project) initialize() error {
	for file := range p.Files {
		utils.MakeFile(p.Files[file], false)
	}

	for folder := range p.Folders {
		utils.MakeFolder(p.Folders[folder], false)
	}

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

	command := exec.Command(p.Tool, p.Arguments...)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		utils.Error(p.Language + " project failed to initialize. Check output below for more details.")
		return err
	}

	return nil
}

func registerProjects(project ...project) {
	registry.projects = append(registry.projects, project...)
}

func flagHandler() {
	args := utils.AdditionalArguments

	if len(args) >= 1 {
		for _, arg := range args {
			switch arg {
			case "-g", "--git":
				gitInit()
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
func gitInit() error {
	init := exec.Command("git", "init")

	if err := init.Run(); err != nil {
		utils.Error("Unable to init git")
		return err
	}

	// Make gitignore before adding files to repo so it can get added.
	utils.MakeFile(".gitignore", false)
	utils.MakeFile("TODO.md", false)
	utils.MakeFile("README.md", false)

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
	projectName := utils.GetUserInput("Enter your project name: ", "my-app")

	program := project{
		Language:   "web",
		Tool:       "bun",
		Arguments:  []string{"create", "vite", projectName, "--template"},
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

	flagHandler()
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
				flagHandler()
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
