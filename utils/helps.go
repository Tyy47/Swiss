package utils

import (
	"fmt"
)

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
	`
	fmt.Println(help_menu)
}

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
