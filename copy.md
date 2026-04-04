package test

import (
	"swiss/build"
	"swiss/command-dict"
	"swiss/create"
	"swiss/initialize"
	"swiss/network"
	"swiss/utils"
)

type Command struct {
	Name        []string
	Subcommands []Subcommand
	Handler     func()
}

type Subcommand struct {
	Flags   []string
	Handler func()
}

type CommandRegistry struct {
	Registry []Command
}

// Command storage
var GlobalCommandRegistry = CommandRegistry{}

// Add command to command storage
func registerCommand(command Command) {
	GlobalCommandRegistry.Registry = append(GlobalCommandRegistry.Registry, command)
}

// Global Arguments
var (
	Arguments           = utils.GatherArgs()
	AdditionalArguments = utils.GatherAdditionalArgs()
)

// Registering commands //

func helpCommand() Command {
	help := Command{
		Name:    []string{"help", "-h"},
		Handler: utils.DisplayHelp,
	}

	return help
}

func versionCommand() Command {
	version := Command{
		Name:    []string{"version", "-v"},
		Handler: utils.PrintVersionNumber,
	}

	return version
}

func ArgParser() {
	// Grabs arguments via wrapper
	args := utils.GatherArgs()

	if len(args) > 1 {
		switch args[1] {
		case "help", "-h":
			utils.DisplayHelp()
		case "version", "-v":
			utils.PrintVersionNumber()
		case "dict":
			dictHandler()
		case "install", "-i":
			build.SwissInstall()
		case "update", "-u":
			build.UpdateSwiss()
		case "build":
			buildHandler()
		case "run":
			runHandler()
		case "init":
			initHandler()
		case "create":
			createHandler()
		case "net":
			netHandler()
		}
	} else {
		utils.DisplayHelp()
	}
}

func dictHandler() {
	args := utils.GatherArgs()
	if len(args) > 2 {
		switch args[2] {
		case "-h", "--help":
			utils.CommandHelp()
		case "bash":
			commanddict.PrintBashCommands()
		case "ps":
			commanddict.PrintPowershellCommands()
		case "git":
			commanddict.PrintGitCommands()
		case "dock", "d", "docker":
			commanddict.PrintDockerCommands()
		default:
			utils.CommandHelp()
		}
	} else {
		utils.CommandHelp()
	}
}

func buildHandler() {
	args := utils.GatherArgs()
	if len(args) > 2 {
		if args[1] == "build" {
			switch args[2] {
			case "-h", "--help":
				utils.BuildHelp()
			case "-l", "--list":
				build.PrintBuildProgramList()
			default:
				build.HandleBuildInput(args[2])
			}
		} else {
			utils.BuildHelp()
		}
	} else {
		utils.BuildHelp()
	}
}

func runHandler() {
	args := utils.GatherArgs()
	if len(args) > 2 {
		if args[1] == "run" {
			switch args[2] {
			case "-h", "--help":
				utils.BuildHelp()
			case "-l", "--list":
				build.PrintRunProgramList()
			default:
				build.HandleRunInput(args[2])
			}
		} else {
			utils.BuildHelp()
		}
	} else {
		utils.BuildHelp()
	}
}

func initHandler() {
	args := utils.GatherArgs()
	if len(args) > 2 {
		if args[1] == "init" {
			switch args[2] {
			case "-h", "--help":
				utils.InitHelp()
			case "-l", "--list":
				initialize.PrintInitProjectList()
			default:
				initialize.HandleInput(args[2])
			}
		} else {
			utils.InitHelp()
		}
	} else {
		utils.InitHelp()
	}
}

func createHandler() {
	create.CreateItems()
}

func netHandler() {
	args := utils.GatherArgs()
	if len(args) > 3 {
		if args[1] == "net" {
			switch args[2] {
			case "-h", "--help":
				utils.NetHelp()
			case "connect":
				network.Connection(args[3])
			case "addr":
				network.GetAddresses(args[3], false)
			case "ns":
				network.GetNameServer(args[3], false)
			case "cname":
				network.GetCNameRecords(args[3], false)
			case "txt":
				network.GetTXTRecords(args[3], false)
			case "mx":
				network.GetMXRecords(args[3], false)
			case "gather":
				network.GatherData(args[3])
			}
		} else {
			utils.NetHelp()
		}
	} else {
		utils.NetHelp()
	}
}

func main() {
	ArgParser()
}

// Registers command into registry on program startup
func init() {
	registerCommand(helpCommand())
}
