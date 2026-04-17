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

	utils.RunCommand(true, "git", "add", ".")
	utils.RunCommand(true, "git", "commit", "-m", commitMessage)
	
	utils.Success("Commit created.")
}

func GitPushSC() {
	commitMessage := utils.CheckArguments(utils.Arguments, 4, 3)
	if commitMessage == "" {
		utils.Warning("Commit message is blank, fill in commit message to continue.")
		return
	}

	utils.RunCommand(true, "git", "add", ".")
	utils.RunCommand(true, "git", "commit", "-m", commitMessage)
	utils.RunCommand(true, "git", "push")

	utils.Success("Commit pushed to repository.")
}
