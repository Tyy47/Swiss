package initialize

import (
	"fmt"
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

type ProjectTemplate struct {
	Language        string // Language name of the project
	BuildTool       string // Build tool of the project
	InitTimes       int    // How many times the project has been init'd during the init process
	Flags           []string
	AdditionalFlags []string
	Handler         func()
	Files           []string
	Folders         []string
	ManualInit      bool
	ManualFiles     []string
	ManualFolders   []string
}

type ProjectRegistry struct {
	Projects []ProjectTemplate
}

func (p *ProjectRegistry) AddToRegistry(project ...ProjectTemplate) {
	p.Projects = append(p.Projects, project...)
}

var pRegistry = ProjectRegistry{
	Projects: []ProjectTemplate{},
}

func PrintInitProjectList() {
	fmt.Println(initProjectList)
}

func (pt *ProjectTemplate) InitializeProject() {
	if pt.ManualInit {
		// Put manual initing logic here then return
		for file := range pt.ManualFiles {
			utils.MakeFile(pt.ManualFiles[file], false)
		}

		for folder := range pt.ManualFolders {
			utils.MakeFolder(pt.ManualFolders[folder], false)
		}

		// Switch case statement to grab language and check to see if the files need to moved anywhere after creation
		switch strings.ToLower(pt.Language) {
		case "c":
			utils.MoveFileToFolder("./main.c", "./src/main.c", true)
		default:
			return
		}

		return
	}

	command := exec.Command(pt.BuildTool, pt.Flags...)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		utils.Error(pt.Language + " project has failed to initalize.")
		utils.Note("Reason: " + err.Error())
		return
	}

	return
}

func golangProject() ProjectTemplate {
	golang := ProjectTemplate{
		Language:   "go",
		BuildTool:  "go",
		InitTimes:  0,
		Flags:      []string{"build"},
		Files:      []string{"main.go"},
		ManualInit: false,
	}
	return golang
}

// Adds created project template objects to registry on launch
func init() {
	projectTemplateArray := []ProjectTemplate{
		golangProject(),
	}

	pRegistry.AddToRegistry(projectTemplateArray...)
}
