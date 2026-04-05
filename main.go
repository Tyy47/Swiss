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
	HelpMenu    func()
}

type Subcommand struct {
	Name  string
	Flags map[string]func()
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

func buildCommand() Command {
	build := Command{
		Name:    "build",
		HelpMenu: utils.BuildHelp,
		Subcommands: []Subcommand{
			{
				Name: "build",
				Flags: map[string]func(){
					"-h": utils.BuildHelp,
					"-l": build.PrintBuildProgramList,
					"go": build.HandleBuildInput,
					"rust": build.HandleBuildInput,
					"c": build.HandleBuildInput,
				},
			},
		},
	}

	return build
}

// Find and run command in registry
func runCommand() {
	if len(utils.Arguments) < 2 {
		utils.DisplayHelp()
		return
	}

	for _, command := range GlobalCommandRegistry.Registry {
		if utils.Arguments[1] == command.Name || slices.Contains(command.Flags, utils.Arguments[1]) {
			if len(command.Subcommands) > 0 {
				runSubcommand(command)
				return
			}
			command.Handler()
			return
		}
	}
	utils.Warning(utils.Arguments[1] + " is not an available command.")
}

func runSubcommand(command Command) {
	if len(utils.Arguments) < 2 { 
		command.HelpMenu()
		return
	}

	for _, arg := range utils.Arguments[1:] {
		for _, sub := range command.Subcommands {
			if arg == sub.Name {
				for _, arg := range utils.Arguments {
					if handler, ok := sub.Flags[arg]; ok {
						handler()
						return
					}
				}
				// No flag found, printing commands help menu
				command.HelpMenu()
				return
			}
		}
	}
	if len(utils.Arguments) > 2 {
		utils.Warning(utils.Arguments[2] + " is not a valid subcommand.")
	} else {
		command.HelpMenu()
	}
}

func main() {
	runCommand()
}

// Registers command into registry on program startup
func init() {
	// Register Commands
	registerCommand(helpCommand())
	registerCommand(versionCommand())
	registerCommand(swissInstallCommand())
	registerCommand(buildCommand())
}
