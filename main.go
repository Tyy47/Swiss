package main

import (
	"slices"

	"swiss/build"
	"swiss/utils"
)

type Command struct {
	Name        string
	Flags       []string
	Subcommands []Subcommand
	Handler     func()
}

type Subcommand struct {
	Name    string
	Flags   []string
	Handler func()
}

type CommandRegistry struct {
	Registry []Command
}

// Command storage
var GlobalCommandRegistry = CommandRegistry{}

// Add command to command storage
func registerCommand(command Command) {
	GlobalCommandRegistry.Registry = append(GlobalCommandRegistry.Registry, command)
}

// Global Arguments
var (
	Arguments           = utils.GatherArgs()
	AdditionalArguments = utils.GatherAdditionalArgs()
)

// Registering commands //

func helpCommand() Command {
	help := Command{
		Name:    "help",
		Flags:   []string{"-h"},
		Handler: utils.DisplayHelp,
	}

	return help
}

func versionCommand() Command {
	version := Command{
		Name:    "version",
		Flags:   []string{"-v"},
		Handler: utils.PrintVersionNumber,
	}

	return version
}

func swissInstallCommand() Command {
	install := Command{
		Name:    "install",
		Flags:   []string{"-i"},
		Handler: build.SwissInstall,
	}

	return install
}

// Find and run command in registry
func runCommand() {
	if len(Arguments) > 1 {
		for command := range GlobalCommandRegistry.Registry {
			if len(Arguments) > 1 && Arguments[1] == GlobalCommandRegistry.Registry[command].Name || slices.Contains(GlobalCommandRegistry.Registry[command].Flags, Arguments[1]) {
				GlobalCommandRegistry.Registry[command].Handler()
				return
			} else if len(Arguments) < 1 {
				utils.DisplayHelp()
				return
			}
		}
		utils.Warning(Arguments[1] + " is not an available command.")
	} else {
		utils.DisplayHelp()
	}
}

func main() {
	runCommand()
}

// Registers command into registry on program startup
func init() {
	registerCommand(helpCommand())
	registerCommand(versionCommand())
	registerCommand(swissInstallCommand())
}
