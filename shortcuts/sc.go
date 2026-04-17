package shortcuts

import "swiss/utils"

func GitCommitSC() {
	var commitMessage string
	if len(utils.AdditionalArguments) <= 0 {
		utils.Warning("Commit message is blank, fill in commit message to continue.")
		return
	} else {
		commitMessage = utils.AdditionalArguments[0]
	}

	utils.RunCommand("git", "add", ".")
	utils.RunCommand("git", "commit", "-m", commitMessage)
	
	utils.Success("Commit created.")
}

func GitPushSC() {
	var commitMessage string
	if len(utils.AdditionalArguments) <= 0 {
		utils.Warning("Commit message is blank, fill in commit message to continue.")
		return
	} else {
		commitMessage = utils.AdditionalArguments[0]
	}

	utils.RunCommand("git", "add", ".")
	utils.RunCommand("git", "commit", "-m", commitMessage)
	utils.RunCommand("git", "push")

	utils.Success("Commit pushed to repository.")
}
