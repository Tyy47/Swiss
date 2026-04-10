package initialize

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"swiss/utils"
)

const initProjectList = `
Rust - Cargo
Go - Go
C - Swiss
HTML - Swiss
Zig - Zig
Vanilla TS Web App - Bun/Vite
Svelte Web App - Bun/Vite
React Web App - Bun/Vite`

type project struct {
	Language  string
	Tool      string
	Arguments []string
	Folders   []string
	Files     []string
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
	if p.ManualInit {
		for file := range p.Files {
			utils.MakeFile(p.Files[file], false)
		}

		for folder := range p.Folders {
			utils.MakeFolder(p.Folders[folder], false)
		}

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

	if p.Tool == "go" && len(p.Arguments) >= 3 && p.Arguments[0] == "mod" {
		utils.MakeFile("main.go", false)
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
		Language:  "rust",
		Tool:      "cargo",
		Arguments: []string{"init"},
		ManualInit: false,
	}

	return program
}

func createGoProject() project {
	args := utils.Arguments
	if len(args) >= 4 {
		program := project{
			Language:  "go",
			Tool:      "go",
			Arguments: []string{"mod", "init", args[3]},
			ManualInit: false,
		}
		return program
	} else {
		program := project{
			Language:  "go",
			Tool:      "go",
			Arguments: []string{"mod", "init", "project"},
			ManualInit: false,
		}
		return program
	}
}

func createCProject() project {
	program := project{
		Language:  "c",
		Tool:      "clang",
		Folders:   []string{"src"},
		Files:     []string{"main.c"},
		ManualInit: true,
	}

	return program
}

func createHTMLProject() project {
	program := project{
		Language:  "html",
		Tool:      "html",
		Folders:   []string{},
		Files:     []string{"TODO.md", "index.html", "styles.css", "main.js"},
		ManualInit: true,
	}

	return program
}

func createZigProject() project {
	program := project{
		Language:  "zig",
		Tool:      "zig",
		Arguments: []string{"init"},
		ManualInit: false,
	}

	return program
}

func createVanillaWebProject() project {
	var projectName string

	if len(utils.Arguments) < 4 {
		projectName = "my-app"
	} else {
		projectName = utils.Arguments[3]
	}

	program := project{
		Language:  "web",
		Tool:      "bun",
		Arguments: []string{"create", "vite", projectName, "--template", "vanilla-ts"},
		ManualInit: false,
	}

	return program
}

func createReactProject() project {
	var projectName string

	if len(utils.Arguments) < 4 {
		projectName = "my-app"
	} else {
		projectName = utils.Arguments[3]
	}

	program := project{
		Language:  "react",
		Tool:      "bun",
		Arguments: []string{"create", "vite", projectName, "--template", "react-ts"},
		ManualInit: false,
	}

	return program
}

func createSvelteProject() project {
	var projectName string

	if len(utils.Arguments) < 4 {
		projectName = "my-app"
	} else {
		projectName = utils.Arguments[3]
	}

	program := project{
		Language:  "svelte",
		Tool:      "bun",
		Arguments: []string{"create", "vite", projectName, "--template", "svelte-ts"},
		ManualInit: false,
	}

	return program

}

func HandleInput() {
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
		createVanillaWebProject(),
		createReactProject(),
		createSvelteProject(),
	}

	registerProjects(projectArray...)
}
