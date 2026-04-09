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

func (pt *ProjectTemplate) InitializeProject() error {
	if pt.InitTimes > 0 {
		utils.Warning(pt.Language + " project is already initialized.")
		return fmt.Errorf("Project has already been initialized")
	}

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
			return nil
		}

		return nil
	}

	command := exec.Command(pt.BuildTool, pt.Flags...)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		utils.Error(pt.Language + " project has failed to initalize.")
		utils.Note("Reason: " + err.Error())
		return err
	}

	return nil
}

func (pt *ProjectTemplate) FlagHandler() {
	args := utils.AdditionalArguments

	if len(args) >= 1 {
		for arg := range args {
			switch args[arg] {
			case "-g", "--git":
				gitInit()
			}
		}
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

func searchInitRegistry(Language string) bool {
	for project := range pRegistry.Projects {
		if Language == pRegistry.Projects[project].Language {
			return true
		}
	}

	return false
}

func ParseInput() {
	if len(utils.Arguments) < 3 {
		return
	}

	for arg := range utils.Arguments[2:] {
		for project := range len(pRegistry.Projects) {
			argument := strings.ToLower(utils.Arguments[arg])
			if argument == pRegistry.Projects[project].Language {
				if pRegistry.Projects[project].InitTimes > 0 {
					utils.Warning(pRegistry.Projects[project].Language + " has been initalized already.")
					return
				}
				if err := pRegistry.Projects[project].InitializeProject(); err != nil {
					log.Fatal(err)
					return
				} else {
					pRegistry.Projects[project].FlagHandler()
					pRegistry.Projects[project].InitTimes += 1
					utils.Success(pRegistry.Projects[project].Language + " project has been created.")
				}
			}
		}
	}
}

// Adds created project template objects to registry on launch
func init() {
	projectTemplateArray := []ProjectTemplate{
		golangProject(),
	}

	pRegistry.AddToRegistry(projectTemplateArray...)
}
