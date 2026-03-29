package create

import (
	"swiss/helps"
	"swiss/utils"
)

func CreateItems() {
	args := utils.GatherArgs()
	var fileToggle bool = false
	var folderToggle bool = false

	if len(args) > 2 {
		for argument := range args {
			if args[argument] == "file" {
				fileToggle = true
				folderToggle = false
				continue
			}

			if args[argument] == "folder" {
				fileToggle = false
				folderToggle = true
				continue
			}

			if folderToggle {
				utils.MakeFolder(args[argument], false)
			}

			if fileToggle {
				utils.MakeFile(args[argument], false)
			}
		}
	} else {
		helps.CreateHelp()
	}
}
