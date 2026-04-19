<h1 align="center">Swiss Documentation</h1>

## Commands

### Misc
A collection of misc commands

`swiss` or `swiss -h` will print out the help menu. \
`swiss -v` or `swiss version` will print out the version name. \
`swiss -v` or `swiss install` will place a Swiss executable in your local/bin. ( Linux command only ) \
`swiss -u` or `swiss update` will clone the repo into an install folder, then will prompt you for install choice. 

### Build
The Build module. Build/Compile or Run programs based on the language provided and that is available via Swiss.

`swiss build -h` or `swiss build --help` to print out the build help menu. \
`swiss build -l` or `--list` will print out a list of languages that can be built/compiled via Swiss.
`swiss build <language : string>` will build/compile the program based on the language provided and that is available via Swiss. \
`swiss run <language : string>` will run the program based on the language provided and that is available via Swiss. 

### Create
The Create module. Create files and or folders through Swiss on mass.

`swiss create -h` or `--help` will print out the help menu for the Create module \
`swiss create <file | folder> <names: string>` Creates folders or files based on the 3rd argument. Names for files must be provided in order to create them. Supports file extensions and creation of multiple files/folders at once

### Init
The Init module. Init a project using Swiss with the available project templates available.

`swiss init -h` or `--help` to print out the help menu for the Init module. \
`swiss init -l` or `--list` to print out a list of languages Swiss can initialize. \
`swiss init <language : string> -g` or `--git` will init a git repo while creating your project. \
`swiss init <language : string> -j` or `--jujutsu` will init a jujutsu repo while creating your project. \
`swiss init <language : string> [project name : string]` Inits a project based on the language provided and the project name for certain languages.

### Net
The Net module. Gather a variety of networking related information about a given domain.

`swiss net -h` or `swiss net --help` to print out the help menu for the Net module. \
`swiss connect <domain : string>` Connects to a domain or an IP address and prints out an http response code. \
`swiss addr <domain : string>` Prints out the IPv4 and v6 addresses of the given domain. \
`swiss ns <domain : string>` Prints out the name servers of the current domain. \
`swiss cname <domain : string>` Prints out the cname records for the given domain. \
`swiss txt <domain : string>` Prints out txt records for the given domain. \
`swiss mx <domain : string>` Prints out mx records for the given domain. \
`swiss gather <domain : string>` Compiles all information that the net module offers and outputs it to a file.

### Gen
The Gen module. Generate codes needed for development via Swiss.

`swiss gen -h` or `swiss net --help` to print out the help menu for the Gen module. \
`swiss gen uuid` generates a 128 bit number in the UUID format \
`swiss gen secret [length : int]` generates a secret code based on the input. Will generate a code thats 16 characters long with no length input. 

### Shortcut
The Shortcut module. A list of command shortcuts that combine most used commands into one to provide a faster developing experience.

`swiss sc commit <message : string>` to add all changes files to your commit and provide a message for your commit. Combined commands: git add . & git commit -m "your message here." \
`swiss sc push [message : string]` to add all your changed files to your commit and provide an optional message to your commit, then pushes all commits to your repository. \
`swiss sc sync` fetch's the repository and grabs any changes, then prints out the git status to see repo changes.
