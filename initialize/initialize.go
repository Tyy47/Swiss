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
Go`

type project struct {
	Language string
	Tool string
	Arguments []string
	Folders []string
	Files []string
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
		p.manualInitialize(p.Folders, p.Files)
		return err
	}

	command := exec.Command(p.Tool, p.Arguments...)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	
	if err := command.Run(); err != nil {
		return err
	}

	if p.Tool == "go" && len(p.Arguments) >= 3 && p.Arguments[0] == "mod" {
		utils.MakeFile("main.go", false)
	}

	return nil
}

func (p *project) manualInitialize(folders []string, files []string) {
	for file := range files {
		utils.MakeFile(files[file], false)
	}

	for folder := range folders {
		utils.MakeFolder(folders[folder], false)
	}
	// Success message is made in the HandleInput() function
}

func createRustProject() project {
	program := project{
		Language: "rust",
		Tool: "cargo",
		Arguments: []string{"init"},
	}

	return program
}

func createGoProject() project {
	args := utils.GatherArgs()
	if len(args) >= 4 {
		program := project{
			Language: "go",
			Tool: "go",
			Arguments: []string{"mod", "init", args[3]},
		}
		return program
	} else {
		program := project{
			Language: "go",
			Tool: "go",
			Arguments: []string{"mod", "init", "project"},
		}
		return program
	}
}

func createCProject() project {
	program := project{
		Language: "c",
		Tool: "clang",
		Arguments: []string{"manual"},
		Folders: []string{"src"},
		Files: []string{"TODO.md", "README.md"},
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
				messages.Success(registry.projects[project].Language + " project has been created.")
				return
			}
		}
	}
	messages.Error("Unable to find " + argument + " in registry list.")
}

func init() {
	registry.projects = append(registry.projects, createRustProject())
	registry.projects = append(registry.projects, createGoProject())
	registry.projects = append(registry.projects, createCProject())
}
