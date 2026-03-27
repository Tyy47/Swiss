package initialize

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"swiss/messages"
	"swiss/utils"
)

const initProjectList = `
Rust
Go
C
HTML`

type project struct {
	Language  string
	Tool      string
	Arguments []string
	Folders   []string
	Files     []string
	GitInit   bool
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

func (p *projectRegistry) addToRegistry(newProject project) {
	p.projects = append(p.projects, newProject)
}

func (p *project) initialize() error {
	var err error
	if p.Arguments[0] == "manual" {
		p.manualInitialize(p.Language, p.Folders, p.Files)
		return err
	}

	command := exec.Command(p.Tool, p.Arguments...)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		messages.Error(p.Language + " project failed to initialize. Check output below for more details.")
		return err
	}

	if p.Tool == "go" && len(p.Arguments) >= 3 && p.Arguments[0] == "mod" {
		utils.MakeFile("main.go", false)
	}

	return nil
}

func (p *project) manualInitialize(language string, folders []string, files []string) {
	for file := range files {
		utils.MakeFile(files[file], false)
	}

	for folder := range folders {
		utils.MakeFolder(folders[folder], false)
	}

	// Switch case statement to grab language and check to see if the files need to moved anywhere after creation
	switch strings.ToLower(language) {
	case "c":
		utils.MoveFileToFolder("./main.c", "./src/main.c", true)
	default:
		return
	}

	// Success message is made in the HandleInput() function
}

// Inits git in current directory when called.
func gitInit() error {
	init := exec.Command("git", "init")

	if err := init.Run(); err != nil {
		messages.Error("Unable to init git")
		return err
	}

	// Make gitignore before adding files to repo so it can get added.
	utils.MakeFile(".gitignore", false)

	add := exec.Command("git", "add", ".")

	if err := add.Run(); err != nil {
		messages.Error("Unable to add files to git project")
		return err
	}

	mainBranch := exec.Command("git", "branch", "-M", "main")

	if err := mainBranch.Run(); err != nil {
		messages.Error("Unable to change main branch to 'main'.")
		return err
	}

	messages.Success("Git has been initialized.")

	return nil
}

func createRustProject() project {
	program := project{
		Language:  "rust",
		Tool:      "cargo",
		Arguments: []string{"init"},
		GitInit:   true,
	}

	return program
}

func createGoProject() project {
	args := utils.GatherArgs()
	if len(args) >= 4 {
		program := project{
			Language:  "go",
			Tool:      "go",
			Arguments: []string{"mod", "init", args[3]},
			GitInit:   true,
		}
		return program
	} else {
		program := project{
			Language:  "go",
			Tool:      "go",
			Arguments: []string{"mod", "init", "project"},
			GitInit:   true,
		}
		return program
	}
}

func createCProject() project {
	program := project{
		Language:  "c",
		Tool:      "clang",
		Arguments: []string{"manual"},
		Folders:   []string{"src"},
		Files:     []string{"TODO.md", "README.md", "main.c"},
		GitInit:   true,
	}

	return program
}

func createHTMLProject() project {
	program := project{
		Language:  "html",
		Tool:      "html",
		Arguments: []string{"manual"},
		Folders:   []string{},
		Files:     []string{"TODO.md", "index.html", "styles.css", "main.js"},
		GitInit:   true,
	}

	return program
}

func HandleInput(argument string) {
	argument = strings.ToLower(argument)
	for project := range len(registry.projects) {
		if argument == registry.projects[project].Language {
			if err := registry.projects[project].initialize(); err != nil {
				log.Fatal(err)
				return
			} else {
				switch registry.projects[project].GitInit {
				case true:
					gitInit()
					messages.Success(registry.projects[project].Language + " project has been created.")
					return
				case false:
					messages.Success(registry.projects[project].Language + " project has been created.")
					return
				}
			}
		}
	}
	messages.Error("Unable to find " + argument + " in registry list.")
}

func init() {
	registry.projects = append(registry.projects, createRustProject())
	registry.projects = append(registry.projects, createGoProject())
	registry.projects = append(registry.projects, createCProject())
	registry.projects = append(registry.projects, createHTMLProject())
}
