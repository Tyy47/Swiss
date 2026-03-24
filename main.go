package main

import (
	"swiss/build"
	"swiss/command-dict"
	"swiss/create"
	"swiss/helps"
	"swiss/initialize"
	"swiss/utils"
)


func ArgParser() {
	// Grabs arguments via wrapper
	args := utils.GatherArgs()
	
	if len(args) > 1 {
		switch args[1] {
		case "help", "-h":
			helpHandler()
		case "version", "-v":
			utils.PrintVersionNumber()
		case "dict":
			dictHandler()
		case "install", "-i":
			build.SwissInstall()
		case "build":
			buildHandler()
		case "run":
			runHandler()
		case "init":
			initHandler()
		case "create":
			createHandler()
		}
	} else {
		helpHandler()
	}
}

func helpHandler() {
	helps.DisplayHelp()
}

func dictHandler() {
	args := utils.GatherArgs()
	if len(args) > 2 {
		switch args[2] {
		case "-h", "--help":
			helps.CommandHelp()
		case "bash":
			commanddict.PrintBashCommands()
		case "ps":
			commanddict.PrintPowershellCommands()
		default:
			helps.CommandHelp()
		}
	} else {
		helps.CommandHelp()
	} 
}

func buildHandler() {
	args := utils.GatherArgs()
	if len(args) > 2 {
		if args[1] == "build" {
			switch args[2] {
			case "-h", "--help":
				helps.BuildHelp()
			case "-l", "--list":
				build.PrintSupportedBuildLanguages()
			default:
				build.CheckLanguageAndBuild(args[2])
			}
		} else {
			helps.BuildHelp()
		}
	} else {
		helps.BuildHelp()
	}
}

func runHandler() {
	args := utils.GatherArgs()
	if len(args) > 2 {
		if args[1] == "run" {
			switch args[2] {
			case "-h", "--help":
				helps.BuildHelp()
			case "-l", "--list":
				build.PrintSupportedRunLanguages()
			default:
				build.CheckLanguageAndRun(args[2])
			}
		} else {
			helps.BuildHelp()
		}
	} else {
		helps.BuildHelp()
	}
}

func initHandler() {
	args := utils.GatherArgs()
	if len(args) > 2 {
		if args[1] == "init" {
			switch args[2] {
			case "-h", "--help":
				helps.InitHelp()
			case "-l", "--list":
				initialize.PrintInitProjectList()
			default:
				initialize.HandleInput(args[2])
			}
		} else {
			helps.InitHelp()
		}
	} else {
		helps.InitHelp()
	}
}

func createHandler() {
	args := utils.GatherArgs()
	if len(args) > 2 {
		if args[1] == "create" {
			switch args[2] {
			case "-h", "--help":
				helps.CreateHelp()
			case "file", "files":
				create.CreateFiles()
			case "folder", "folders":
				create.CreateFolders()
			default:
				helps.CreateHelp()
			}
		} else {
			helps.CreateHelp()
		}
	} else {
		helps.CreateHelp()
	}
}

func main() {
	ArgParser()
}
