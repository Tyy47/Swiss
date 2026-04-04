package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
)

const VERSION_NUMBER string = "1.0.2"

// Global Arguments
var (
	Arguments           = gatherArgs()
	AdditionalArguments = gatherAdditionalArgs()
)


func PrintVersionNumber() {
	fmt.Println("Swiss version number: " + ColorGreen + VERSION_NUMBER + ColorReset)
}

// Checks if an error has a value, if so, returns an empty value, an error message and the fatal log.
func CrashCheck(err error) {
	if err != nil {
		Error("Swiss has crashed! View output below to learn more.")
		log.Fatal(err)
		return
	}
}

// Gathers argument via the os library
func gatherArgs() []string {
	args := os.Args; return args
}

func gatherAdditionalArgs() []string {
	if len(Arguments) < 3 {
		return []string{}
	}

	return Arguments[3:]
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
		Warning(file + " file exists.")
		return
	} else {
		err := os.WriteFile(file, []byte(""), 0o666)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	if !muted {
		Success(file + " file created.")
	}
}

func MakeFolder(folder string, muted bool) {
	dirInfo, err := CheckFolderExists(folder)

	if err != nil {
		Error("Unable to create directory.")
		return
	}

	if dirInfo {
		Warning(folder + "folder exists.")
		return
	} else {
		err := os.Mkdir(folder, 0755)
		if err != nil {
			Error(err.Error())
			return
		}
	}
	if !muted {
		Success("Successfully created " + folder + ".")
	}
}

func MoveFileToFolder(oldPath string, newPath string, muted bool) {
	if err := os.Rename(oldPath, newPath); err != nil {
		Error("Unable to move file: " + err.Error())
		return
	}

	if !muted {
		Success("Successfully moved " + oldPath + " to " + newPath + ".")
	}
}

func DoesToolExist(tool string) bool {
	command := exec.Command(tool, "-v")

	if err := command.Run(); err != nil {
		Error(tool + " is not installed or added to path.")
		return false
	}

	return true
}

