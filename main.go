package main

import (
	"swiss/utils"
)

type Command struct {
	Name        []string
	Subcommands []Subcommand
	Handler     func()
}

type Subcommand struct {
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
		Name:    []string{"help", "-h"},
		Handler: utils.DisplayHelp,
	}

	return help
}

func versionCommand() Command {
	version := Command{
		Name:    []string{"version", "-v"},
		Handler: utils.PrintVersionNumber,
	}

	return version
}

func main() {
}

// Registers command into registry on program startup
func init() {
	registerCommand(helpCommand())
}
