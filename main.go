package main

import (
	"fmt"
	"swiss/build"
	"swiss/gen"
	"swiss/initialize"
	"swiss/network"
	"swiss/shortcuts"
	"swiss/utils"
	
	"github.com/Tyy47/clibox/colorbin"
	"github.com/Tyy47/clibox/argbin"
)

// Creating the root object of the application
var root = argbin.Root{
	AppName: "swiss",
	AppVersion: "1.2",
	Description: "the cli army knife",
	CommandList: make([]*argbin.Command, 0),
}

// Creates the "help" command and returns it
func helpCommand() *argbin.Command {
	return &argbin.Command{
		Name: "-h",
		AdditionalNames: []string{"--help"},
		Execute: func(ctx *argbin.Context) error {
			utils.DisplayHelp()
			return nil
		},
	}
}

// Creates the "version" command and returns it
func versionCommand() *argbin.Command {
	return &argbin.Command{
		Name: "-v",
		AdditionalNames: []string{"--version"},
		Execute: func(ctx *argbin.Context) error {
			// Gather app version
			version, err := root.GetAppVersion()
			if err != nil {
				return err
			}
			
			// Color app version to green
			version = colorbin.Green(version).ToHighIntensityBold().String()
			
			// Profit
			fmt.Printf("swiss: version %s", version)
			return nil
		},
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
			"ts":     func(args *[]string) { initialize.CreateProject() },
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


func main() {
	
	// Command storage to add to root.AddCommand
	commands := []*argbin.Command{
		helpCommand(),
		versionCommand(),
		build.BuildCommand(),
	}

	// Adds all commands to app
	if err := root.AddCommand(commands...); err != nil {
		panic(err)
	}

	
	// Starts the project using argbin
	if err := root.Run(); err != nil {
		panic(err)
	}
}
