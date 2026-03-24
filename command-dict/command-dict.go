package commanddict

import "fmt"


const commandListPowershell string = `
== A list of Powershell commands. Visit their documentation for more info. ==

- dir: Shows files in current directory.`;

const commandListBash string = `
== A list of Bash commands. Visit their man, info, or help pages to learn more. ==

- man, info, --help: Shows documentation on a command.
- ls: Shows files in current directory.
- cat: Reads a file and prints it out to the terminal.
- grep: Searches through an output to find the given argument.
- ip link: Shows info on all network interfaces.
- dmesg: Shows system logs.`;

func PrintPowershellCommands() {
	fmt.Println(commandListPowershell)
}

func PrintBashCommands() {
	fmt.Println(commandListBash)
}
