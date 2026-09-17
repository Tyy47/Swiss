package main

import (
	"fmt"
	"swiss/build"
	"swiss/gen"
	"swiss/initialize"
	"swiss/shortcuts"
	
	"github.com/Tyy47/clibox/colorbin"
	"github.com/Tyy47/clibox/argbin"
)

// Creating the root object of the application
var root = argbin.Root{
	AppName: "swiss",
	AppVersion: "1.2",
	CommandList: make([]*argbin.Command, 0),
	Description: `
╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
help: Opens the help menu.
build <string>: Allows you to build program via swiss.
init <string>: Inits a project using Swiss.
gen: A variety of codes that can be generated via Swiss.
sc: Shortcuts that are multiple commands in one.`,
}

// helpCommand creates the "help" command for the root.
func helpCommand() *argbin.Command {
	return &argbin.Command{
		Name: "-h",
		AdditionalNames: []string{"--help"},
		Execute: func(ctx *argbin.Context) error {
			fmt.Println(root.GetDescription())
			return nil
		},
	}
}

// versionCommand creates the "version" command for the root.
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
			fmt.Printf("swiss: version %s\n", version)
			return nil
		},
	}
}

// helpFlag creates a generic flag for a command to print out the commands help menu.
func helpFlag() *argbin.Flag {
	return &argbin.Flag{
		Execute: func(ctx *argbin.Context) error {
			fmt.Println(ctx.Command.GetDescription())
			return nil
		},
	}
}


func main() {
	
	// Command storage to add to root.AddCommand
	commands := []*argbin.Command{
		helpCommand(),
		versionCommand(),
		build.BuildCommand(),
		build.RunCommand(),
		gen.GenerateCommand(),
		initialize.InitCommand(),
		shortcuts.ShortcutCommand(),
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

func init() {
	for _, cmd := range root.CommandList {
		cmd.AddFlag("-h", helpFlag())
		cmd.AddFlag("--help", helpFlag())

		for _, subCmd := range cmd.Subcommands {
			subCmd.AddFlag("-h", helpFlag())
			subCmd.AddFlag("--help", helpFlag())
		}
	}
}
