package help

import (
	"fmt"
)

// Prints out the title section of Swiss.
func titleCard() {
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
	titleCard()
	var help_menu string = `
usage: swiss [module_name] [additional_arguments]
	
Swiss Commands:

Swiss utility commands:
	help - Displays the main help menu for Swiss. Running swiss with a module name will display that modules help menu.
	install - Installs swiss when ran inside of a cloned version of the Swiss repository.
	update - Updates swiss by cloning the repository, you'll have two methods of install when prompted.

Swiss Flags:
	-u, --unstable: Clones the unstable branch of the Swiss repo and installs it.

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
	titleCard()
	var help_menu string = `
usage: swiss dict <arguments> [flags]

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
	titleCard()
	var help_menu string = `
usage: swiss build <argument> [flags] | swiss run <argument> [flags]

Build Commands - Builds or Runs a program using it's main tool. Check the list of languages to see the tools required.

Commands:
	build: Compiles a program based on the language argument provided
	run: Runs a program based on the language argument provided

Flags:
	-h --help: Opens the help menu.
	-l --list: Prints a list of available languages to build and run with their respective build tools available in Swiss.
`
	fmt.Println(help_menu)
}

// Prints the help menu for the Init module
func InitHelp() {
	titleCard()
	var help_menu string = `
usage: swiss init <argument> [flags]

Init Commands - Initialize a project using Swiss.

Commands:
	init: Inits a project based on the given input. 
	init web: Inits a web based project using Bun & Vite with a selected framework. 

Flags:
	-h --help: Opens the help menu.
	-l --list: Prints a list of projects that can be initialized and if they are supported with additional arguments for names.
	-g --git: Inits git alongside your project.
	-j --jujutsu: Inits jj alongside your project.
`
	fmt.Println(help_menu)
}

// Prints the help menu for the Networking module
func NetHelp() {
	titleCard()
	var help_menu string = `
usage: swiss net <arguments> [flags]

Net Commands - A variety of networking tools.

Commands:
	connect: Prints out an http response code when connecting to the domain and port.
	port: Attempts to connect to the domain and check if the port is opened or closed.
	addr: Prints out the IPv4 and v6 addresses of the given domain.
	ns: Prints out the name servers of the current domain.
	cname: Prints out the cname records for the given domain.
	txt: Prints out txt records for the given domain.
	mx: Prints out mx records for the given domain.
	gather: Compiles all information that the net module offers and outputs it to a file.

Flags:
	-h --help: Opens the help menu.
`
	fmt.Println(help_menu)
}

// Prints the help menu for the Generator module
func GenHelp() {
	titleCard()
	var help_menu string = `
usage: swiss gen <arguments> [flags]

Gen module - Generate codes through Swiss.

Commands:
	uuid: Generates an 128 bit hexadecimal string.
	secret [length : int]: Generates a hexadecimal string based on length provided, 16 characters long by default.

Flags:
	-h --help: Opens the help menu.
`
	fmt.Println(help_menu)
}

// Prints the help menu for the Shortcut module
func ShortcutHelp() {
	titleCard()
	var help_menu string = `
usage: swiss sc <arguments> [flags]

Shortcut module - Commands that are multiple commands into one.

Commands:
	commit <message : string>: Adds all changed files to commit with a message.
	push [message : string]: Adds all files, commits changes with a message, then pushes to your repository.
	sync: Fetch's all changes to the repository and prints a status message with changes to the repository.

Flags:
	-h --help: Opens the help menu.
`
	fmt.Println(help_menu)
}
