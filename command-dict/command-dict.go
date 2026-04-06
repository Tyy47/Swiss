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

-- Changes
pull: Pulls changes to branch from the repository.
push: Pushes changes to branch to the repository.
stash: Put away, grab, or clear changes you've stashed`

const commandListDocker string = `
== A list of Docker commands. Visit the official documentation to learn more or -h on the commands to view help in the terminal. ==

-- General Commands
--help: Shows help menu for docker and subcommands
info: Shows system-wide docker information

-- Images
build -t <image_name : string>: Build an image from a dockerfile ( . --no-cache at the end to build without a cache. ).
images: Prints a list of local images.
rmi <image_name : string>: Deletes an image.
image prune: Removes all unused images.

-- Containers
run --name <container_name : string> <image_name : string>: Create and run a container from an image with a custom name.
run -p <host_port : int>:<container_port : int> <image_name : string>: Runs a container with and publish a container's port/s to the host.
run -d <image_name : string>: Run a container in the background.
start | stop <container_name | container_id : string>: Start or stop a container.
rm <container_name : string>: Remove a stopped container.
exec -it <container_name : string> sh: Open a shell inside a running container
logs -f <container_name : string>: Fetch and follows logs of the given container.
inspect <container_name | container_id : string>: Inspects a given container.
ps: Lists all running containers
ps --all: Lists all active and inactive containers.
container stats: Views resource usages of containers.

-- Docker Hub
login -u <username : string>: Login into Docker Hub
push <username : string>/<image_name : string>: Publish an image to Docker Hub
search <image_name : string>: Searches Docker Hub for image.
pull <image_name : string>: Pulls image from Docker Hub.`

func PrintPowershellCommands() {
	fmt.Println(commandListPowershell)
}

func PrintBashCommands() {
	fmt.Println(commandListBash)
}

func PrintGitCommands() {
	fmt.Println(commandListGit)
}

func PrintDockerCommands() {
	fmt.Println(commandListDocker)
}
