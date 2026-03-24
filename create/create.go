package create

import (
	"swiss/utils"
)

// Make a create folder function and a create file function.

func CreateFiles() {
	args := utils.GatherAdditionalArgs()

	for files := range len(args) {
		utils.MakeFile(args[files], false)
	}
}

func CreateFolders() {
	args := utils.GatherAdditionalArgs()

	for folders := range len(args) {
		utils.MakeFolder(args[folders], false)
	}
}
