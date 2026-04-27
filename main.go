package main

import (
	"swiss/build"
	commanddict "swiss/command-dict"
	"swiss/gen"
	"swiss/initialize"
	"swiss/network"
	"swiss/shortcuts"
	"swiss/utils"
)

// Command struct to store information about modules commands
type Command struct {
	Name          string
	Flags         []string
	Subcommands   map[string]func(args *[]string)
	Handler       func()
	HelpMenu      func()
	SingleRun     bool                 // Boolean statement to check if the program can run with no arguments like "swiss build"
	ShortHandFunc func(args *[]string) // Short hand function that runs if single run function is ran like "swiss build"
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
		Handler: utils.DisplayHelp,
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
		Flags:   []string{"-u"},
		Handler: build.UpdateSwiss,
	}
}

// Creates the "build" command and returns it
func buildCommand() Command {
	return Command{
		Name:     "build",
		HelpMenu: utils.BuildHelp,
		Subcommands: map[string]func(args *[]string){
			"-h":     func(args *[]string) { utils.BuildHelp() },
			"--help": func(args *[]string) { utils.BuildHelp() },
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
		HelpMenu: utils.BuildHelp,
		Subcommands: map[string]func(args *[]string){
			"-h":     func(args *[]string) { utils.BuildHelp() },
			"--help": func(args *[]string) { utils.BuildHelp() },
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
		HelpMenu: utils.CommandHelp,
		Subcommands: map[string]func(args *[]string){
			"-h":     func(args *[]string) { utils.CommandHelp() },
			"--help": func(args *[]string) { utils.CommandHelp() },
			"ps":     func(args *[]string) { commanddict.PrintPowershellCommands() },
			"bash":   func(args *[]string) { commanddict.PrintBashCommands() },
			"git":    func(args *[]string) { commanddict.PrintGitCommands() },
			"docker": func(args *[]string) { commanddict.PrintDockerCommands() },
		},
	}
}

// Creates the "init" command and returns it
func initCommand() Command {
	return Command{
		Name:     "init",
		HelpMenu: utils.InitHelp,
		Subcommands: map[string]func(args *[]string){
			"-h":     func(args *[]string) { utils.InitHelp() },
			"--help": func(args *[]string) { utils.InitHelp() },
			"-l":     func(args *[]string) { initialize.PrintInitProjectList() },
			"--list": func(args *[]string) { initialize.PrintInitProjectList() },
			"go":     func(args *[]string) { initialize.CreateProject() },
			"rust":   func(args *[]string) { initialize.CreateProject() },
			"c":      func(args *[]string) { initialize.CreateProject() },
			"html":   func(args *[]string) { initialize.CreateProject() },
			"zig":    func(args *[]string) { initialize.CreateProject() },
			"python": func(args *[]string) { initialize.CreateProject() },
			"web":    func(args *[]string) { initialize.CreateWebProject() },
		},
	}
}

// Creates the "net" command and returns it
func netCommand() Command {
	return Command{
		Name:     "net",
		HelpMenu: utils.NetHelp,
		Subcommands: map[string]func(args *[]string){
			"-h":      func(args *[]string) { utils.NetHelp() },
			"--help":  func(args *[]string) { utils.NetHelp() },
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
		HelpMenu: utils.GenHelp,
		Subcommands: map[string]func(args *[]string){
			"-h":     func(args *[]string) { utils.GenHelp() },
			"--help": func(args *[]string) { utils.GenHelp() },
			"uuid":   func(args *[]string) { gen.GenerateUUID() },
			"secret": func(args *[]string) { gen.GenerateSecret() },
		},
	}
}

// Creates the "shortcut" command and returns it
func shortcutCommand() Command {
	return Command{
		Name:     "sc",
		HelpMenu: utils.ShortcutHelp,
		Subcommands: map[string]func(args *[]string){
			"-h":     func(args *[]string) { utils.ShortcutHelp() },
			"--help": func(args *[]string) { utils.ShortcutHelp() },
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
		utils.DisplayHelp()
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
			for _, subArg := range args[i+1:] {
				// If the subcommand exists, it will execute the subcommand function
				if subFunc, ok := cmd.Subcommands[subArg]; ok {
					subFunc(&utils.Arguments)
				}
			}
			//return Arg parser attempt change
		}
	}
	utils.Warning(utils.Arguments[1] + " is not an available command.")
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
		generateCommand(),
		shortcutCommand(),
	}

	GlobalCommandDatabase.registerCommand(commandArray...)
}
