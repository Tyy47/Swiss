package build

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"swiss/utils"
)

const runProgramList = `
Rust: Cargo
Go: Go
Python: Python
Typescript: Bun
`

type run struct {
	Language  string
	Tool      string
	Arguments []string
	RunFile   string
	BuildRequired bool // boolean statement to check if the program needs to built before running
	BuildFunction func() run // Function to be ran if BuildRequired is toggled to true
}

type runRegistry struct {
	runs []run
}

var runStorage = runRegistry{
	runs: []run{},
}

func PrintRunProgramList() {
	utils.Note("Languages are listed with their build tools.")
	fmt.Println(runProgramList)
}

func (r *runRegistry) addToRunRegistry(newRun ...run) {
	r.runs = append(r.runs, newRun...)
}

func (r *run) initializeRun() error {
	if r.BuildRequired == true {
		r.BuildFunction()
	}

	command := exec.Command(r.Tool, r.Arguments...)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		return err
	}

	return nil
}

func runRustProject() run {
	return run{
		Language:  "rust",
		Tool:      "cargo",
		Arguments: []string{"run"},
		RunFile: "Cargo.toml",
	}
}

func runGoProject() run {
	return run{
		Language:  "go",
		Tool:      "go",
		Arguments: []string{"run", "main.go"},
		RunFile: "main.go",
	}
}

func runPythonProject() run {
	return run{
		Language:  "python",
		Tool:      "python",
		Arguments: []string{"main.py"},
		RunFile: "main.py",
	}
}

func runTypescriptProject() run {
	return run{
		Language: "typescript",
		Tool: "bun",
		Arguments: []string{"main.ts"},
		RunFile: "main.ts",
	}
}

func runCProject() run {
	return run{
		Language: "c",
		Tool: "./main",
		RunFile: "main.c",
		BuildRequired: true,
		BuildFunction: func() run {
			project := buildCProject()

			if err := project.Initialize(); err != nil {
				utils.Error("Unable to build c program.")
			}

			return run{}
		},
	}
}

func scanForRunFiles() (bool, run) {
	// Scan directory for all files.
	files, err := os.ReadDir(".")
	if err != nil {
		utils.Error("Unable to read files in current directory.")
		return false, run{}
	}

	for _, project := range runStorage.runs {
		for _, file := range files {
			if file.Name() == project.RunFile {
				return true, project
			}
		}
	}

	utils.Warning("Unable to find inputted language, check language list for runnable languages via Swiss.")
	return false, run{}
}

func RunProject() {
	// Grabs the bool and build struct from scanForRunFiles()
	result, project := scanForRunFiles()

	if !result {
		utils.Error("Unable to run project, check inputted language to see if it's in Swiss run list.")
		return
	}

	if err := project.initializeRun(); err != nil {
		utils.Error("Unable to run project.")
		utils.Reason(err.Error())
		return
	}
	utils.Success(project.Language + " project has been ran.")
}

func HandleRunInput() {
	if len(utils.Arguments) < 3 {
		return
	}

	argument := strings.ToLower(utils.Arguments[2])

	for run := range len(runStorage.runs) {
		if argument == runStorage.runs[run].Language {
			if err := runStorage.runs[run].initializeRun(); err != nil {
				utils.Error("Unable to run project.")
				utils.Reason(err.Error())
				return
			} else {
				utils.Success(runStorage.runs[run].Language + " project has been ran.")
				return
			}
		}
	}
	utils.Error("Unable to find " + argument + " in registry list.")
}

func init() {
	runArray := []run{
		runGoProject(),
		runRustProject(),
		runPythonProject(),
		runTypescriptProject(),
		runCProject(),
	}

	runStorage.addToRunRegistry(runArray...)
}
