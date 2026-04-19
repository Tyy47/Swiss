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
	Subcommands []Subcommand
	Handler     func()
	HelpMenu    func()
}

// Subcommand struct to store information about a given commands flags and associated functions through a map
type Subcommand struct {
	Name  string
	Flags map[string]func()
}

// Command storage struct
type CommandRegistry struct {
	Registry []Command
}

// Command storage
var GlobalCommandRegistry = CommandRegistry{}

// Add command to command storage
func registerCommand(command ...Command) {
	GlobalCommandRegistry.Registry = append(GlobalCommandRegistry.Registry, command...)
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

func swissUpdateCommand() Command {
	update := Command{
		Name:    "update",
		Flags:   []string{"-u"},
		Handler: build.UpdateSwiss,
	}

	return update
}

func buildCommand() Command {
	build := Command{
		Name:     "build",
		HelpMenu: utils.BuildHelp,
		Subcommands: []Subcommand{
			{
				Name: "build",
				Flags: map[string]func(){
					"-h":     utils.BuildHelp,
					"--help": utils.BuildHelp,
					"-l":     build.PrintBuildProgramList,
					"--list": build.PrintBuildProgramList,
					"go":     build.HandleBuildInput,
					"rust":   build.HandleBuildInput,
					"c":      build.HandleBuildInput,
					"zig":    build.HandleBuildInput,
				},
			},
		},
	}

	return build
}

func runRunCommand() Command {
	run := Command{
		Name:     "run",
		HelpMenu: utils.BuildHelp,
		Subcommands: []Subcommand{
			{
				Name: "run",
				Flags: map[string]func(){
					"-h":     utils.BuildHelp,
					"--help": utils.BuildHelp,
					"-l":     build.PrintRunProgramList,
					"--list": build.PrintRunProgramList,
					"go":     build.HandleRunInput,
					"rust":   build.HandleRunInput,
					"c":      build.HandleRunInput,
					"python": build.HandleRunInput,
				},
			},
		},
	}

	return run
}

func dictionaryCommand() Command {
	dict := Command{
		Name:     "dict",
		HelpMenu: utils.CommandHelp,
		Subcommands: []Subcommand{
			{
				Name: "dict",
				Flags: map[string]func(){
					"-h":     utils.CommandHelp,
					"--help": utils.CommandHelp,
					"ps":     commanddict.PrintPowershellCommands,
					"bash":   commanddict.PrintBashCommands,
					"git":    commanddict.PrintGitCommands,
					"docker": commanddict.PrintDockerCommands,
				},
			},
		},
	}

	return dict
}

func initCommand() Command {
	init := Command{
		Name:     "init",
		HelpMenu: utils.InitHelp,
		Subcommands: []Subcommand{
			{
				Name: "init",
				Flags: map[string]func(){
					"-h":     utils.InitHelp,
					"--help": utils.InitHelp,
					"-l":     initialize.PrintInitProjectList,
					"--list": initialize.PrintInitProjectList,
					"go":     initialize.CreateProject,
					"rust":   initialize.CreateProject,
					"c":      initialize.CreateProject,
					"html":   initialize.CreateProject,
					"zig":    initialize.CreateProject,
					"python": initialize.CreateProject,
					"web":    initialize.CreateWebProject,
				},
			},
		},
	}

	return init
}

func createCommand() Command {
	create := Command{
		Name:     "create",
		HelpMenu: utils.CreateHelp,
		Subcommands: []Subcommand{
			{
				Name: "create",
				Flags: map[string]func(){
					"-h":     utils.CreateHelp,
					"--help": utils.CreateHelp,
					"create": create.CreateItems,
				},
			},
		},
	}
	return create
}

func netCommand() Command {
	net := Command{
		Name:     "net",
		HelpMenu: utils.NetHelp,
		Subcommands: []Subcommand{
			{
				Name: "net",
				Flags: map[string]func(){
					"-h":      utils.NetHelp,
					"--help":  utils.NetHelp,
					"connect": network.Connection,
					"port":    network.GetPortStatus,
					"addr":    network.GetAddresses,
					"ns":      network.GetNameServer,
					"cname":   network.GetMXRecords,
					"txt":     network.GetTXTRecords,
					"mx":      network.GetMXRecords,
					"gather":  network.GatherData,
				},
			},
		},
	}
	return net
}

func generateCommand() Command {
	gen := Command{
		Name:     "gen",
		HelpMenu: utils.GenHelp,
		Subcommands: []Subcommand{
			{
				Name: "gen",
				Flags: map[string]func(){
					"-h":     utils.GenHelp,
					"--help": utils.GenHelp,
					"uuid":   gen.GenerateUUID,
					"secret": gen.GenerateSecret,
				},
			},
		},
	}

	return gen
}

func shortcutCommand() Command {
	shortcut := Command{
		Name:     "sc",
		HelpMenu: utils.ShortcutHelp,
		Subcommands: []Subcommand{
			{
				Name: "sc",
				Flags: map[string]func(){
					"-h":     utils.ShortcutHelp,
					"commit": shortcuts.GitCommitSC,
					"push":   shortcuts.GitPushSC,
				},
			},
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

	registerCommand(commandArray...)
}
