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
	}

	return runRust
}

func runGoProject() run {
	goRun := run{
		Language:  "go",
		Tool:      "go",
		Arguments: []string{"run", "main.go"},
	}

	return goRun
}

func runPythonProject() run {
	pythonRun := run{
		Language: "python",
		Tool: "python",
		Arguments: []string{"main.py"},
	}

	return pythonRun
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
