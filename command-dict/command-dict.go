package commanddict

import "fmt"


const commandListPowershell string = `
== A list of Powershell commands. Visit the official documentation for more info. ==

- dir: Shows files in current directory.`

const commandListBash string = `
== A list of Bash commands. Visit their man, info, or help pages to learn more. ==

- man, info, --help: Shows documentation on a command.
- ls: Shows files in current directory.
- cat: Reads a file and prints it out to the terminal.
- grep: Searches through an output to find the given argument.
- ip link: Shows info on all network interfaces.
- dmesg: Shows system logs.`

const commandListGit string = `
== A list of Git commands. Visit the official documentation to learn more or -h on the commands to view help in the terminal. ==

-- Start a project:
init: Initializes git in the current directory.
clone <repository_url: string> [directory : string]: Clones a repository to current directory or optional directory if provided.

-- Branches
branch <branch_name : string>: Creates a branch in the repo
switch <branch_name : string>: Switches the current branch to the one given.

-- Files
add <items : string>: Adds files or folders to repositories tracking.
rm <items : string>: Remove files or folders from repository tracking.
commit -m <message : string>: Adds a commit message to the repository.

-- Getting Changes
pull: Pulls changes to branch from the repository.
push: Pushes changes to branch to the repository.`

const commandListDocker string = ``

func PrintPowershellCommands() {
	fmt.Println(commandListPowershell)
}

func PrintBashCommands() {
	fmt.Println(commandListBash)
}

func PrintGitCommands() {
	fmt.Println(commandListGit)
}
