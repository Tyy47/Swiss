package build

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"swiss/messages"
)

const runProgramList = `
Rust: Cargo
Go: Go
`

type run struct {
	Language string
	Tool string
	Arguments []string
}

type runRegistry struct {
	runs []run
}

var runStorage = runRegistry{
	runs: []run{},
}

func PrintRunProgramList() {
	messages.Note("Languages are listed with their build tools.")
	fmt.Println(buildProgramList)
}

func (r *runRegistry) addToRunRegistry(newRun run) {
	r.runs = append(r.runs, newRun)
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
		Language: "rust",
		Tool: "cargo",
		Arguments: []string{"run"},
	}

	return runRust
}

func runGoProject() run {
	goRun := run{
		Language: "go",
		Tool: "go",
		Arguments: []string{"run", "main.go"},
	}

	return goRun
}

func HandleRunInput(argument string) {
	argument = strings.ToLower(argument)
	for run := range len(runStorage.runs) {
		if argument == runStorage.runs[run].Language {
			if err := runStorage.runs[run].initializeRun(); err != nil {
				log.Fatal(err)
				return
			} else {
				messages.Success(runStorage.runs[run].Language + " project has been ran.")
				return
			}
		}
	}
	messages.Error("Unable to find " + argument + " in registry list.")
}

func init() {
	runStorage.runs = append(runStorage.runs, runRustProject())
	runStorage.runs = append(runStorage.runs, runGoProject())
}
