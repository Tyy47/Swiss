package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"swiss/colors"
	"swiss/messages"
)

const VERSION_NUMBER string = "1.0.2"


func PrintVersionNumber() {
	fmt.Println("Swiss version number: " + colors.ColorGreen + VERSION_NUMBER + colors.ColorReset)
}

// Gathers argument via the os library
func GatherArgs() []string {
	args := os.Args; return args
}

func GatherAdditionalArgs() []string {
	var additional []string
	args := GatherArgs()

	for _, v := range args[3:] {
		additional = append(additional, v)
	}

	return additional
}

func CheckFileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return err  == nil && info.Mode().IsRegular()
}

func CheckFolderExists(folderName string) (bool, error) {
	info, err := os.Stat(folderName)
	if err == nil {
		return info.IsDir(), nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err

}

func GetUsersName() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to get current user: %s", err)
	}

	return user.Username
}

func GetOperatingSystem() string {
	// Returns Windows, Linux, or Darwin ( Apple )
	return runtime.GOOS
}

func MakeFile(file string, muted bool) {
	if CheckFileExists(file) {
		messages.Warning(file + " file exists.")
		return
	} else {
		err := os.WriteFile(file, []byte(""), 0o666)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	if !muted {
		messages.Success(file + " file created.")
	}
}

func MakeFolder(folder string, muted bool) {
	dirInfo, err := CheckFolderExists(folder)

	if err != nil {
		messages.Error("Unable to create directory.")
		return
	}

	if dirInfo {
		messages.Warning(folder + "folder exists.")
		return
	} else {
		err := os.Mkdir(folder, 0755)
		if err != nil {
			messages.Error(err.Error())
			return
		}
	}
	if !muted {
		messages.Success("Successfully created " + folder + ".")
	}
}

func MoveFileToFolder(oldPath string, newPath string, muted bool) {
	if err := os.Rename(oldPath, newPath); err != nil {
		messages.Error("Unable to move file: " + err.Error())
		return
	}

	if !muted {
		messages.Success("Successfully moved " + oldPath + " to " + newPath + ".")
	}
}

func DoesToolExist(tool string) bool {
	command := exec.Command(tool, "-v")

	if err := command.Run(); err != nil {
		messages.Error(tool + " is not installed or added to path.")
		return false
	}

	return true
}
