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
`

type run struct {
	Language  string
	Tool      string
	Arguments []string
	RunFile   string
}

type runRegistry struct {
	runs []run
}

var runStorage = runRegistry{
	runs: []run{},
}

func PrintRunProgramList() {
	utils.Note("Languages are listed with their build tools.")
	fmt.Println(buildProgramList)
}

func (r *runRegistry) addToRunRegistry(newRun ...run) {
	r.runs = append(r.runs, newRun...)
}

func (r *run) initializeRun() error {
	command := exec.Command(r.Tool, r.Arguments...)

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		return err
	}

	return nil
}

func runRustProject() run {
	runRust := run{
		Language:  "rust",
		Tool:      "cargo",
		Arguments: []string{"run"},
		RunFile: "Cargo.toml",
	}

	return runRust
}

func runGoProject() run {
	goRun := run{
		Language:  "go",
		Tool:      "go",
		Arguments: []string{"run", "main.go"},
		RunFile: "main.go",
	}

	return goRun
}

func runPythonProject() run {
	pythonRun := run{
		Language:  "python",
		Tool:      "python",
		Arguments: []string{"main.py"},
		RunFile: "main.py",
	}

	return pythonRun
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

	utils.Warning("Unable to find inputted language, check language list for buildable languages via Swiss.")
	return false, run{}
}

func RunProject() {
	// Grabs the bool and build struct from scanForRunFiles()
	result, project := scanForRunFiles()

	if !result {
		utils.Error("Unable to run project, check inputted language to see if it's in Swiss run list.")
	}

	project.initializeRun()
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
				utils.Crash(err)
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
	}

	runStorage.addToRunRegistry(runArray...)
}
