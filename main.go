package main

import (
	"errors"
	"fmt"
	"swiss/build"
	"swiss/gen"
	"swiss/initialize"
	"swiss/shortcuts"
	"swiss/utils"

	"github.com/Tyy47/clibox/argbin"
	"github.com/Tyy47/clibox/colorbin"
)

// Creating the root object of the application
var root = argbin.Root{
	AppName: "swiss",
	AppVersion: "1.2",
	CommandList: make([]*argbin.Command, 0),
	HelpMenu: `
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
		Name: "help",
		AdditionalNames: []string{"--help", "-h"},
		Execute: func(ctx *argbin.Context) error {
			fmt.Println(root.HelpMenu)
			return nil
		},
	}
}

// versionCommand creates the "version" command for the root.
func versionCommand() *argbin.Command {
	return &argbin.Command{
		Name: "version",
		AdditionalNames: []string{"--version", "-v"},
		Execute: func(ctx *argbin.Context) error {
			// Color app version to green
			version := colorbin.Green(root.AppVersion).ToHighIntensityBold().String()
			
			// Profit
			fmt.Printf("swiss: version %s\n", version)
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
		if errors.Is(err, argbin.ErrMissingArguments) {
			fmt.Println(root.HelpMenu)
			return
		}

		utils.Output.Error(err)
	}
}
