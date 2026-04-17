package shortcuts

import "swiss/utils"

func GitCommitSC() {
	// Gathers a commit message
	var commitMessage string
	if len(utils.AdditionalArguments) <= 0 {
		utils.Warning("Commit message is blank, fill in commit message to continue.")
		return
	} else {
		commitMessage = utils.AdditionalArguments[0]
	}
	
	// Adds all files to the commit with a message.
	utils.RunCommand("git", "add", ".")
	utils.RunCommand("git", "commit", "-m", commitMessage)
	
	// Message stating that the commit was created.
	utils.Success("Commit created.")
}

func GitPushSC() {
	// Adds all changed files to commit
	utils.RunCommand("git", "add", ".")

	// Gathers a commit message if one is available.
	var commitMessage string
	if len(utils.AdditionalArguments) > 0 {
		commitMessage = utils.AdditionalArguments[0]
		utils.RunCommand("git", "commit", "-m", commitMessage)
	}

	// Pushes changes to repository
	utils.RunCommand("git", "push")

	// Message stating that changes were pushed
	utils.Success("Commit pushed to repository.")
}
