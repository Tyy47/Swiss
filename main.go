package main

import (
	"slices"

	"swiss/build"
	commanddict "swiss/command-dict"
	"swiss/create"
	"swiss/gen"
	"swiss/initialize"
	"swiss/network"
	"swiss/shortcuts"
	"swiss/utils"
)

// Command struct to store information about modules commands
type Command struct {
	Name        string
	Flags       []string
	Subcommands map[string]func(args *[]string)
	Handler     func()
	HelpMenu    func()
}

// Command storage struct
type CommandDB struct {
	Registry []Command
}

// Command storage
var GlobalCommandDatabase = CommandDB{}

// Add command to command storage
func (c *CommandDB) registerCommand(command ...Command) {
	c.Registry = append(c.Registry, command...)
}

// Registering commands //

func helpCommand() Command {
	return Command{
		Name:    "help",
		Flags:   []string{"-h"},
		Handler: utils.DisplayHelp,
	}
}

func versionCommand() Command {
	return Command{
		Name:    "version",
		Flags:   []string{"-v"},
		Handler: utils.PrintVersionNumber,
	}
}

func swissInstallCommand() Command {
	return Command{
		Name:    "install",
		Flags:   []string{"-i"},
		Handler: build.SwissInstall,
	}
}

func swissUpdateCommand() Command {
	return Command{
		Name:    "update",
		Flags:   []string{"-u"},
		Handler: build.UpdateSwiss,
	}
}

func buildCommand() Command {
	return Command{
		Name:     "build",
		HelpMenu: utils.BuildHelp,
		Subcommands: map[string]func(args *[]string) {
			"-h": func(args *[]string) { utils.BuildHelp() },
			"--help": func(args *[]string) { utils.BuildHelp() },
			"go": func(args *[]string) { build.HandleBuildInput() },
			"rust": func(args *[]string) { build.HandleBuildInput() },
			"c": func(args *[]string) { build.HandleBuildInput() },
			"zig": func(args *[]string) { build.HandleBuildInput() },
		},
	}
}

func runRunCommand() Command {
	return Command{
		Name:     "run",
		HelpMenu: utils.BuildHelp,
		Subcommands: map[string]func(args *[]string) {
			"-h": func(args *[]string) { utils.BuildHelp() },
			"--help": func(args *[]string) { utils.BuildHelp() },
			"-l": func(args *[]string) { build.PrintRunProgramList() },
			"--list": func(args *[]string) { build.PrintRunProgramList() },
			"go": func(args *[]string) { build.HandleRunInput() },
			"rust": func(args *[]string) { build.HandleRunInput() },
			"c": func(args *[]string) { build.HandleRunInput() },
			"python": func(args *[]string) { build.HandleRunInput() },
		},
	}
}

func dictionaryCommand() Command {
	return Command{
		Name:     "dict",
		HelpMenu: utils.CommandHelp,
		Subcommands: map[string]func(args *[]string){
			"-h": func(args *[]string) { utils.CommandHelp() },
			"--help": func(args *[]string) { utils.CommandHelp() },
			"ps": func(args *[]string) { commanddict.PrintPowershellCommands() },
			"bash": func(args *[]string) { commanddict.PrintBashCommands() },
			"git": func(args *[]string) { commanddict.PrintGitCommands() },
			"docker": func(args *[]string) { commanddict.PrintDockerCommands() },
		},
	}
}

func initCommand() Command {
	return Command{
		Name:     "init",
		HelpMenu: utils.InitHelp,
		Subcommands: map[string]func(args *[]string){
			"-h": func(args *[]string) { utils.InitHelp() },
			"--help": func(args *[]string) { utils.InitHelp() },
			"-l": func(args *[]string) { initialize.PrintInitProjectList() },
			"--list": func(args *[]string) { initialize.PrintInitProjectList() },
			"go": func(args *[]string) { initialize.CreateProject() },
			"rust": func(args *[]string) { initialize.CreateProject() },
			"c": func(args *[]string) { initialize.CreateProject() },
			"html": func(args *[]string) { initialize.CreateProject() },
			"zig": func(args *[]string) { initialize.CreateProject() },
			"python": func(args *[]string) { initialize.CreateProject() },
			"web": func(args *[]string) { initialize.CreateWebProject() },
		},
	}
}

func createCommand() Command {
	return Command{
		Name:     "create",
		HelpMenu: utils.CreateHelp,
		Subcommands: map[string]func(args *[]string){
			"-h": func(args *[]string) { utils.CreateHelp() },
			"--help": func(args *[]string) { utils.CreateHelp() },
			"create": func(args *[]string) { create.CreateItems() },
		},
	}
}

func netCommand() Command {
	return Command{
		Name:     "net",
		HelpMenu: utils.NetHelp,
		Subcommands: map[string]func(args *[]string){
			"-h": func(args *[]string) { utils.NetHelp() },
			"--help": func(args *[]string) { utils.NetHelp() },
			"connect": func(args *[]string) { network.Connection() },
			"port": func(args *[]string) { network.GetPortStatus() },
			"addr": func(args *[]string) { network.GetAddresses() },
			"ns": func(args *[]string) { network.GetNameServer() },
			"cname": func(args *[]string) { network.GetCNameRecords() },
			"txt": func(args *[]string) { network.GetTXTRecords() },
			"mx": func(args *[]string) { network.GetMXRecords() },
			"gather": func(args *[]string) { network.GatherData() },
		},
	}
}

func generateCommand() Command {
	return Command{
		Name:     "gen",
		HelpMenu: utils.GenHelp,
		Subcommands: map[string]func(args *[]string){
			"-h": func(args *[]string) { utils.GenHelp() },
			"--help": func(args *[]string) { utils.GenHelp() },
			"uuid": func(args *[]string) { gen.GenerateUUID() },
			"secret": func(args *[]string) { gen.GenerateSecret() },
		},
	}
}

func shortcutCommand() Command {
	shortcut := Command{
		Name:     "sc",
		HelpMenu: utils.ShortcutHelp,
		Subcommands: map[string]func(args *[]string){
			"-h": func(args *[]string) { utils.ShortcutHelp() },
			"--help": func(args *[]string) { utils.ShortcutHelp() },
			"commit": func(args *[]string) { shortcuts.GitCommitSC() },
			"push": func(args *[]string) { shortcuts.GitPushSC() },
			"sync": func(args *[]string) { shortcuts.GitSyncSC() },
		},
	}

	return shortcut
}

// Find and run command in registry
func runCommand() {
	if len(utils.Arguments) < 2 {
		utils.DisplayHelp()
		return
	}

	for _, command := range GlobalCommandDatabase.Registry {
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
	if len(utils.Arguments) <= 2 {
		command.HelpMenu()
		return
	}

	for _, arg := range utils.Arguments[1:] {
		for _, sub := range command.Subcommands {
			if arg == sub.Name {
				for _, argument := range utils.Arguments {
					if handler, ok := sub.Flags[argument]; ok {
						handler()
					}
				}
			}
		}
	}
}

func main() {
	runCommand()
}

// Registers command into registry on program startup
func init() {
	// Register Commands
	commandArray := []Command{
		helpCommand(),
		versionCommand(),
		swissInstallCommand(),
		swissUpdateCommand(),
		buildCommand(),
		runRunCommand(),
		dictionaryCommand(),
		initCommand(),
		netCommand(),
		createCommand(),
		generateCommand(),
		shortcutCommand(),
	}

	GlobalCommandDatabase.registerCommand(commandArray...)
}
