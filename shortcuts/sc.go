package shortcuts

import (
	"os"
	"os/exec"
	"swiss/utils"
	"time"

	"github.com/Tyy47/clibox/argbin"
	"github.com/Tyy47/clibox/outbin"
)

var output = outbin.NewOutput(os.Stdout, os.Stderr)

// gitAddCommand creates a exec.Command and run's "git add .". Returns an exec error if unable to add files.
func gitAddCommand() error {
	
	// Create the add command
	addCmd := exec.Command("git", "add", ".")

	// Run the add command
	if err := addCmd.Run(); err != nil {
		return err
	}

	return nil
}

// Adds all files to a commit using git with a required message
func gitCommit(commitMessage string) error { 

	// Run git add command
	if err := gitAddCommand(); err != nil {
		return err
	}

	// Create commit command
	commitCmd := exec.Command("git", "commit", "-m", commitMessage)

	// Run git commit command and return the potential error
	if err := commitCmd.Run(); err != nil {
		return err
	}
	
	return nil
}

func gitPush(commitMessage string) error {

	// Run the git add command
	if err := gitAddCommand(); err != nil {
		return err
	}

	// Run the git commit command with a commit message from the user
	if err := gitCommit(commitMessage); err != nil {
		return err
	}

	time.Sleep(2 * time.Second)

	// Create the push command
	pushCmd := exec.Command("git", "push")

	utils.ToggleOutputForCMD(pushCmd)

	// Execute the push command
	if err := pushCmd.Run(); err != nil {
		return err
	}

	return nil
}

func gitPull() error {

	// Create the pull command
	pullCmd := exec.Command("git", "pull")

	// Execute the pull command
	if err := pullCmd.Run(); err != nil {
		return err
	}

	return nil
}

func gitSync() error  {
	
	// Create the fetch command
	fetchCmd := exec.Command("git", "fetch")

	// Run the fetch command
	if err := fetchCmd.Run(); err != nil {
		return err
	}

	// Create the status command
	statusCmd := exec.Command("git", "status")

	utils.ToggleOutputForCMD(statusCmd)

	// Run the status command
	if err := statusCmd.Run(); err != nil {
		return err
	}

	return nil
}

// gitCommitCommand creates the "commit" command. commit is a compressed and shorthand form of commiting to a git repository.
func gitCommitCommand() *argbin.Command {
	return &argbin.Command{
		Name: "commit",
		HelpMenu: "commit shortcut",
		TakesValue: true,
		Execute: func(ctx *argbin.Context) error {
			
			// Run commit command
			if err := gitCommit(ctx.ParsedValue); err != nil {
				return err
			}
			
			// Success message
			output.Success("git commit created.")

			return nil
		},
	}
}

func gitPushCommand() *argbin.Command {
	return &argbin.Command{
		Name: "push",
		HelpMenu: "push shortcut",
		TakesValue: true,
		Execute: func(ctx *argbin.Context) error {
			
			// Runs the git add command
			if err := gitAddCommand(); err != nil {
				return err
			}

			// Runs git commit -m ctx.ParsedValue
			if err := gitCommit(ctx.ParsedValue); err != nil {
				return err
			}

			// Runs the git push command
			if err := gitPush(ctx.ParsedValue); err != nil {
				return err
			}

			return nil
		},
	}
}

func gitSyncCommand() *argbin.Command {
	return &argbin.Command{
		Name: "sync",
		HelpMenu: "sync shortcut",
		Execute: func(ctx *argbin.Context) error {
			
			// Runs git fetch & git status
			if err := gitSync(); err != nil {
				return err
			}

			pull, _ := ctx.Values["pull"].(bool)
			if pull {
				if err := gitPull(); err != nil {
					return err
				}
			}
			
			output.Success("Local repository updated.")
			return nil
		},
		Flags: argbin.Flags{
			"-p": {
				Execute: func(ctx *argbin.Context) error {
					ctx.Values["pull"] = true
					return nil
				},
			},
			"--pull": {
				Execute: func(ctx *argbin.Context) error {
					ctx.Values["pull"] = true
					return nil
				},
			},
		},
	}
}

func ShortcutCommand() *argbin.Command {
	return &argbin.Command{
		Name: "sc",
		AdditionalNames: []string{"shortcut"},
		Subcommands: []*argbin.Command{
			gitCommitCommand(),
			gitPushCommand(),
			gitSyncCommand(),
		},
		HelpMenu: `
╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
Shortcut module - Commands that are multiple commands into one.

Commands:
	commit <message : string>: Adds all changed files to commit with a message.
	push [message : string]: Adds all files, commits changes with a message, then pushes to your repository.
	sync: Fetch's all changes to the repository and prints a status message with changes to the repository.

Flags:
	-h --help: Opens the help menu.`,
	}
}
