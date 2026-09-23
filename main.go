package main

import (
	"errors"
	"fmt"

	"github.com/Tyy47/clibox/argbin"
	"github.com/Tyy47/clibox/colorbin"
	"swiss/build"
	"swiss/gen"
	"swiss/help"
	"swiss/initialize"
	"swiss/shortcuts"
	"swiss/utils"
)

// Creating the root object of the application
var root = argbin.Root{
	AppName:     "swiss",
	AppVersion:  "1.2",
	CommandList: make([]*argbin.Command, 0),
	HelpMenu: `
╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
	
usage: swiss [command] [additional_arguments] <flags>
	
Commands:
	build: Builds a program that uses swiss made shortcuts.
	run: Runs a program that uses swiss made shortcuts.
	init: Initializes a programming based project in current folder.
	gen: Generates different codes that are most commonly used in development
	sc: Command shortcuts for various CLI utilities to make development faster

Flags:
	-h, --help: Displays the swiss help menu
	-v, --version: Displays the current swiss version number`,
}

<<<<<<< HEAD
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

// Looks through a map of the project registry and returns the lookup
func commandLookup() map[string]*Command {
	lookup := make(map[string]*Command)
	for i := range GlobalCommandDatabase.Registry {
		cmd := &GlobalCommandDatabase.Registry[i]
		lookup[cmd.Name] = cmd
		for _, flag := range cmd.Flags {
			lookup[flag] = cmd
		}
	}
	return lookup
}

// Creates the "help" command and returns it
func helpCommand() Command {
	return Command{
		Name:    "help",
		Flags:   []string{"-h"},
		Handler: help.DisplayHelp,
	}
}

// Creates the "version" command and returns it
func versionCommand() Command {
	return Command{
		Name:    "version",
		Flags:   []string{"-v"},
		Handler: utils.PrintVersionNumber,
	}
}

// Creates the "install" command and returns it
func swissInstallCommand() Command {
	return Command{
		Name:    "install",
		Flags:   []string{"-i"},
		Handler: build.SwissInstall,
	}
}

// Creates the "update" command and returns it
func swissUpdateCommand() Command {
	return Command{
		Name:    "update",
		Flags:   []string{"update"},
		Handler: func() { build.UpdateSwiss(&utils.Arguments) },
	}
}

// Creates the "build" command and returns it
func buildCommand() Command {
	return Command{
		Name:     "build",
		HelpMenu: help.BuildHelp,
		Subcommands: map[string]func(args *[]string){
			"-h":     func(args *[]string) { help.BuildHelp() },
			"--help": func(args *[]string) { help.BuildHelp() },
			"go":     func(args *[]string) { build.HandleBuildInput() },
			"rust":   func(args *[]string) { build.HandleBuildInput() },
			"c":      func(args *[]string) { build.HandleBuildInput() },
			"zig":    func(args *[]string) { build.HandleBuildInput() },
		},
		SingleRun:     true,
		ShortHandFunc: func(args *[]string) { build.BuildProject() },
	}
}

// Creates the "run" command and returns it
func runRunCommand() Command {
	return Command{
		Name:     "run",
		HelpMenu: help.BuildHelp,
		Subcommands: map[string]func(args *[]string){
			"-h":     func(args *[]string) { help.BuildHelp() },
			"--help": func(args *[]string) { help.BuildHelp() },
			"-l":     func(args *[]string) { build.PrintRunProgramList() },
			"--list": func(args *[]string) { build.PrintRunProgramList() },
			"go":     func(args *[]string) { build.HandleRunInput() },
			"rust":   func(args *[]string) { build.HandleRunInput() },
			"c":      func(args *[]string) { build.HandleRunInput() },
			"python": func(args *[]string) { build.HandleRunInput() },
		},
		ShortHandFunc: func(args *[]string) { build.RunProject() },
		SingleRun:     true,
	}
}

// Creates the "dictionary" command and returns it
func dictionaryCommand() Command {
	return Command{
		Name:     "dict",
		HelpMenu: help.CommandHelp,
		Subcommands: map[string]func(args *[]string){
			"-h":     func(args *[]string) { help.CommandHelp() },
			"--help": func(args *[]string) { help.CommandHelp() },
			"ps":     func(args *[]string) { commanddict.PrintPowershellCommands() },
			"bash":   func(args *[]string) { commanddict.PrintBashCommands() },
			"git":    func(args *[]string) { commanddict.PrintGitCommands() },
			"docker": func(args *[]string) { commanddict.PrintDockerCommands() },
=======
// helpCommand creates the "help" command for the root.
func helpCommand() *argbin.Command {
	return &argbin.Command{
		Name:            "help",
		AdditionalNames: []string{"--help", "-h"},
		Execute: func(ctx *argbin.Context) error {
			fmt.Println(root.HelpMenu)
			return nil
>>>>>>> cleanup
		},
	}
}

<<<<<<< HEAD
// Creates the "init" command and returns it
func initCommand() Command {
	return Command{
		Name:     "init",
		HelpMenu: help.InitHelp,
		Subcommands: map[string]func(args *[]string){
			"-h":     func(args *[]string) { help.InitHelp() },
			"--help": func(args *[]string) { help.InitHelp() },
			"-l":     func(args *[]string) { initialize.PrintInitProjectList() },
			"--list": func(args *[]string) { initialize.PrintInitProjectList() },
			"go":     func(args *[]string) { initialize.CreateProject() },
			"rust":   func(args *[]string) { initialize.CreateProject() },
			"c":      func(args *[]string) { initialize.CreateProject() },
			"html":   func(args *[]string) { initialize.CreateProject() },
			"zig":    func(args *[]string) { initialize.CreateProject() },
			"python": func(args *[]string) { initialize.CreateProject() },
			"ts":     func(args *[]string) { initialize.CreateProject() },
			"web":    func(args *[]string) { initialize.CreateWebProject() },
=======
// versionCommand creates the "version" command for the root.
func versionCommand() *argbin.Command {
	return &argbin.Command{
		Name:            "version",
		AdditionalNames: []string{"--version", "-v"},
		Execute: func(ctx *argbin.Context) error {
			// Color app version to green
			version := colorbin.Green(root.AppVersion).ToHighIntensityBold().String()

			// Profit
			fmt.Printf("swiss: version %s\n", version)
			return nil
>>>>>>> cleanup
		},
	}
}

<<<<<<< HEAD
// Creates the "net" command and returns it
func netCommand() Command {
	return Command{
		Name:     "net",
		HelpMenu: help.NetHelp,
		Subcommands: map[string]func(args *[]string){
			"-h":      func(args *[]string) { help.NetHelp() },
			"--help":  func(args *[]string) { help.NetHelp() },
			"connect": func(args *[]string) { network.Connection() },
			"port":    func(args *[]string) { network.GetPortStatus() },
			"addr":    func(args *[]string) { network.GetAddresses() },
			"ns":      func(args *[]string) { network.GetNameServer() },
			"cname":   func(args *[]string) { network.GetCNameRecords() },
			"txt":     func(args *[]string) { network.GetTXTRecords() },
			"mx":      func(args *[]string) { network.GetMXRecords() },
			"gather":  func(args *[]string) { network.GatherData() },
		},
	}
}

// Creates the "generate" command and returns it
func generateCommand() Command {
	return Command{
		Name:     "gen",
		HelpMenu: help.GenHelp,
		Subcommands: map[string]func(args *[]string){
			"-h":     func(args *[]string) { help.GenHelp() },
			"--help": func(args *[]string) { help.GenHelp() },
			"uuid":   func(args *[]string) { gen.GenerateUUID() },
			"secret": func(args *[]string) { gen.GenerateSecret() },
		},
	}
}

// Creates the "shortcut" command and returns it
func shortcutCommand() Command {
	return Command{
		Name:     "sc",
		HelpMenu: help.ShortcutHelp,
		Subcommands: map[string]func(args *[]string){
			"-h":     func(args *[]string) { help.ShortcutHelp() },
			"--help": func(args *[]string) { help.ShortcutHelp() },
			"commit": func(args *[]string) { shortcuts.GitCommitSC() },
			"push":   func(args *[]string) { shortcuts.GitPushSC() },
			"sync":   func(args *[]string) { shortcuts.GitSyncSC() },
		},
	}
}

// Find and run command in registry
func runCommand() {
	// Checks if the length of the users given arguments are less then two, if so, displays the main swiss help menu.
	if len(utils.Arguments) < 2 {
		help.DisplayHelp()
		return
	}

	// Grabs arguments past swiss
	args := utils.Arguments[1:]
	// Creates the lookup map from the commandLookup function
	lookup := commandLookup()

	// Loops through the arguments with an integer place for each argument
	for i, arg := range args {
		// Loops through the lookup map for commands to see if the exist
		if cmd, ok := lookup[arg]; ok {
			// If there is a function made in Handler, it will run it.
			if cmd.Handler != nil {
				cmd.Handler()
			}

			// Statement to check if a command has SingleRun functionality. If it does, it runs the shorthand function.
			// Else, if the length of the arguments is less then or equal to two and it has a help menu, it will display a help menu.
			if cmd.SingleRun && cmd.ShortHandFunc != nil && len(utils.Arguments) == 2 {
				cmd.ShortHandFunc(&utils.Arguments)
			} else if len(utils.Arguments) <= 2 && cmd.HelpMenu != nil {
				cmd.HelpMenu()
			}

			// Loops through all the valid subcommands
			subArgs := args[i+1:]
			matched := false
			for _, subArg := range subArgs {
				// If the subcommand exists, it will execute the subcommand function
				if subFunc, ok := cmd.Subcommands[subArg]; ok {
					subFunc(&utils.Arguments)
					matched = true
				}
			}
			// If a subcommand was provided but none matched, give the user feedback.
			if len(subArgs) > 0 && !matched && len(cmd.Subcommands) > 0 {
				utils.Warning(subArgs[0] + " is not a valid subcommand for " + cmd.Name + ".")
				if cmd.HelpMenu != nil {
					cmd.HelpMenu()
				}
			}
			return
		}
	}
	utils.Warning(utils.Arguments[1] + " is not an available command.")
}

func main() {
	runCommand()
}

// Registers command into registry on program startup
func init() {
	// Register Commands into an array
	commandArray := []Command{
=======
func main() {
	// Command storage to add to root.AddCommand
	commands := []*argbin.Command{
>>>>>>> cleanup
		helpCommand(),
		versionCommand(),
		build.BuildCommand(),
		build.RunCommand(),
		build.SwissInstall(),
		gen.GenerateCommand(),
		initialize.InitCommand(),
		shortcuts.ShortcutCommand(),
	}
<<<<<<< HEAD
	
	// Registers commands one by one by unpacking command array
	GlobalCommandDatabase.registerCommand(commandArray...)
=======

	// Adds all commands to app
	if err := root.AddCommand(commands...); err != nil {
		// Panic for potential programmer errors
		panic(err)
	}

	// Starts the project using argbin
	if err := root.Run(); err != nil {
		if errors.Is(err, argbin.ErrMissingArguments) {
			fmt.Println(root.HelpMenu)
			return
		}

		utils.Output.Error(err)
	}
>>>>>>> cleanup
}
