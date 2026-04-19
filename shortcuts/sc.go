package shortcuts

import "swiss/utils"

// Adds all files to a commit using git with a required message
func GitCommitSC() {
	// Gathers a commit message
	var commitMessage string
	if len(utils.AdditionalArguments) <= 0 {
		utils.Warning("Commit message is blank, fill in commit message to continue.")
		return
	} else {
		commitMessage = utils.AdditionalArguments[0]
	}

	// Checks if adding files to commit will cause an error
	if err := utils.RunCommand("git", "add", "."); err != nil {
		utils.Error("Unable to add files to commit, make sure there is changes to add.")
		return
	}
	
	// Checks if message can be added to commit
	if err := utils.RunCommand("git", "commit", "-m", commitMessage); err != nil {
		utils.Error("Unable to add files to commit, make sure there is changes to add.")
		return
	}

	// Message stating that the commit was created.
	utils.Success("Commit created.")
}

func GitPushSC() {
	// Gathers a commit message if one is available.
	var commitMessage string
	if len(utils.AdditionalArguments) > 0 {
		// Adds all changed files to commit
		if err := utils.RunCommand("git", "add", "."); err != nil {
			utils.Error("Unable to add files to commit, make sure there is changes to add.")
			return
		}
		
		// Assigns message to commitMessage then commits
		commitMessage = utils.AdditionalArguments[0]
		if err := utils.RunCommand("git", "commit", "-m", commitMessage); err != nil {
			utils.Error("Unable to add files to commit, make sure there is changes to add.")
			return
		}
	}

	// Pushes changes to repository
	if err := utils.RunCommand("git", "push"); err != nil {
		utils.Error("Unable to push changes to repository")
		return
	}

	// Message stating that changes were pushed
	utils.Success("Commit pushed to repository.")
}
