package utils

import (
	"fmt"
)

// Prints the main help menu for Swiss
func DisplayHelp() {
	var help_menu string = `
╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
help: Opens the help menu.
dict <string>: Shows system commands based on the shell argument provided.
build <string>: Allows you to build program via swiss.
init <string>: Inits a project using Swiss.
create <string>: Make a set of folders/files using Swiss.
net: A set of networking tools.
gen: A variety of codes that can be generated via Swiss.
sc: Shortcuts that are multiple commands in one.
	`
	fmt.Println(help_menu)
}

// Prints the help menu for the Command Dictionary module
func CommandHelp() {
	var help_menu string = `
╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
Command Dictionary Module - Contains a variety of sub commands that show the most used commands depending on the input.

-h --help: Opens the help menu
ps: Prints Powershell command dictionary.
bash: Prints Bash command dictionary.
git: Prints git command dictionary.
docker: Prints docker command dictionary.
`
	fmt.Println(help_menu)
}

// Prints the help menu for the Build module
func BuildHelp() {
	var help_menu string = `
╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
Build module - Builds or Runs a program based on the language inputted.

-h --help: Opens the help menu.
-l --list: Prints a list of available languages to build and run with their respective build tools available in Swiss.
build <string>: Builds a program based on the language you input.
run <string>: Runs a program based on the language you input.
`
	fmt.Println(help_menu)
}

// Prints the help menu for the Create module
func CreateHelp() {
	var help_menu string = `
╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
Create module - Creates folder or files based on your inputs.

-h --help: Opens the help menu.
create <file | folder> <names: string>: Makes folders/files via Swiss. Affix file or folder to specify what to create.
`
	fmt.Println(help_menu)
}

// Prints the help menu for the Init module
func InitHelp() {
	var help_menu string = `
╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
Init module - Initialize a project using Swiss.

-h --help: Opens the help menu.
-l --list: Prints a list of projects that can be initialized and if they are supported with additional arguments for names.
-g --git: Inits git alongside your project.
-j --jujutsu: Inits jj alongside your project.
init <string> [name: string]: Inits a project based on the given input. 
`
	fmt.Println(help_menu)
}

// Prints the help menu for the Networking module
func NetHelp() {
	var help_menu string = `
╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
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
	var help_menu string = `
╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
Gen module - Generate codes through Swiss.

-h --help: Opens the help menu.
uuid: Generates an 128 bit hexadecimal string.
secret [length : int]: Generates a hexadecimal string based on length provided, 16 characters long by default.
`
	fmt.Println(help_menu)
}

// Prints the help menu for the Shortcut module
func ShortcutHelp() {
	var help_menu string = `
╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
Shortcut module - Commands that are multiple commands into one.

-h --help: Opens the help menu.
commit <message : string>: Adds all changed files to commit with a message.
push [message : string]: Adds all files, commits changes with a message, then pushes to your repository.
sync: Fetch's all changes to the repository and prints a status message with changes to the repository.
`
	fmt.Println(help_menu)
}
