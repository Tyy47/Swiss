package help

import (
	"fmt"
)

// Prints out the title section of Swiss.
func TitleCard() {
	title := `╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
	`
	// Having Print instead of Println makes it log to the terminal without a newline which works better for the extended strings
	fmt.Print(title)
}

// Prints the main help menu for Swiss
func DisplayHelp() {
	TitleCard()
	var help_menu string = `
usage: swiss [module_name] [additional_arguments]
	
Swiss Commands:

Swiss utility commands:
	help - Displays the main help menu for Swiss. Running swiss with a module name will display that modules help menu.
	install - Installs the newest version of Swiss available on GitHub.

Build Module:
	build - Builds a program that uses swiss made shortcuts.
	run - Runs a program that uses swiss made shortcuts.

Initialize Module:
	init - Initializes a programming based project in current folder.

Dictionary Module:
	dict: Shows system and utility commands for other programs and scripting languages.

Network Module:
	net - Provides a set of networking tools to check connections.

Generator Module: 
	gen - Generates different codes that are most commonly used in development

Shortcut Module:
	sc - Command shortcuts for various CLI utilities to make development faster
`
	fmt.Println(help_menu)
}

// Prints the help menu for the Command Dictionary module
func CommandHelp() {
	TitleCard()
	var help_menu string = `
usage: swiss dict <arguments>

Dictionary Commands - Each command prints out commonly used commands for each scripting language or cli utility:

Scripting Languages: 
	ps: Prints out a list PowerShell commands.
	bash: Prints out a list Bash commands.

Tools:
	git: Prints out a list Git commands.
	docker: Prints out a list Docker commands.

Flags:
	-h | --help: Prints out the dictionary help menu.
`
	fmt.Println(help_menu)
}

// Prints the help menu for the Build module
func BuildHelp() {
	TitleCard()
	var help_menu string = `
Build module - Builds or Runs a program based on the language inputted.

-h --help: Opens the help menu.
-l --list: Prints a list of available languages to build and run with their respective build tools available in Swiss.
build <string>: Builds a program based on the language you input.
run <string>: Runs a program based on the language you input.
`
	fmt.Println(help_menu)
}

// Prints the help menu for the Init module
func InitHelp() {
	TitleCard()
	var help_menu string = `
Init module - Initialize a project using Swiss.

-h --help: Opens the help menu.
-l --list: Prints a list of projects that can be initialized and if they are supported with additional arguments for names.
-g --git: Inits git alongside your project.
-j --jujutsu: Inits jj alongside your project.
init <string> [name: string]: Inits a project based on the given input. 
init web: Inits a web based project using Bun & Vite with a selected framework. 
`
	fmt.Println(help_menu)
}

// Prints the help menu for the Networking module
func NetHelp() {
	TitleCard()
	var help_menu string = `
Net module - A variety of networking tools.

-h --help: Opens the help menu.
connect <domain : string>: Prints out an http response code when connecting to the domain and port.
port <domain : string> <port : string> Attempts to connect to the domain and check if the port is opened or closed.
addr <domain : string>: Prints out the IPv4 and v6 addresses of the given domain.
ns <domain : string>: Prints out the name servers of the current domain.
cname <domain : string>: Prints out the cname records for the given domain.
txt <domain : string>: Prints out txt records for the given domain.
mx <domain : string>: Prints out mx records for the given domain.
gather <domain : string>: Compiles all information that the net module offers and outputs it to a file.
`
	fmt.Println(help_menu)
}

// Prints the help menu for the Generator module
func GenHelp() {
	TitleCard()
	var help_menu string = `
Gen module - Generate codes through Swiss.

-h --help: Opens the help menu.
uuid: Generates an 128 bit hexadecimal string.
secret [length : int]: Generates a hexadecimal string based on length provided, 16 characters long by default.
`
	fmt.Println(help_menu)
}

// Prints the help menu for the Shortcut module
func ShortcutHelp() {
	TitleCard()
	var help_menu string = `
Shortcut module - Commands that are multiple commands into one.

-h --help: Opens the help menu.
commit <message : string>: Adds all changed files to commit with a message.
push [message : string]: Adds all files, commits changes with a message, then pushes to your repository.
sync: Fetch's all changes to the repository and prints a status message with changes to the repository.
`
	fmt.Println(help_menu)
}
