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
}

type Items struct {
	folders []string
	files []string
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


func (i *Items) createItems() {
	for file := range len(i.files) {
		utils.MakeFile(i.files[file], false)
	}

	for folder := range len(i.folders) {
		utils.MakeFolder(i.folders[folder], false)
	}
}

func (p *project) initialize() error {
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
}
