package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"runtime"

	"github.com/Tyy47/clibox/outbin"
	"github.com/Tyy47/clibox/argbin"
)

<<<<<<< HEAD
// Swiss version number
const VERSION_NUMBER string = "1.1a"

// Global Arguments
=======
>>>>>>> cleanup
var (
	

	// util errors
	ErrUnableToCreateFile = errors.New("unable to create file in current directory")
	ErrToolNotInstalled = errors.New("is not installed")
)

var Output = outbin.NewOutput(os.Stdout, os.Stderr)

// CheckFileExists takes in a list of file names and checks if they exist.
// Results are added to a map to iterate on, an error is returned if the stats of a file cannot be retrieved.
func CheckFileExists(files ...string) (map[string]bool, error) {

	fileMap := make(map[string]bool)

	for _, file := range files {

<<<<<<< HEAD
	comm.Stdin = os.Stdin
	comm.Stdout = os.Stdout
	comm.Stderr = os.Stderr
=======
		_, err := os.Stat(file)
>>>>>>> cleanup

		// if the file doesn't exist add it to the map with a value of false
		if os.IsNotExist(err) { fileMap[file] = false }

		// If the file exists it adds it to the file mape with a value of true
		if err == nil { fileMap[file] = true }

	}

	return fileMap, nil
}

// CheckFolderExists takes in a list of folder names and iterates over them to see if they exist.
// Results are added to a map to iterate upon and an error is returned if there is a problem
// gathering the folders.
func CheckFolderExists(folders ...string) (map[string]bool, error) {
	
	folderMap := make(map[string]bool)

	for _, folder := range folders {
	
		// Grab folder info
		_, err := os.Stat(folder)

		// If the folder exists, add it to the map with a value of true
		if err == nil { folderMap[folder] = true }

		// If the folder doesnt exist, add it to the map with a value of false
		if os.IsNotExist(err) { folderMap[folder] = false }
	}

	return folderMap, nil
}

// GetUsersName gathers the user's username from the user package and returns it.
func GetUsersName() string {
	user, err := user.Current()
	if err != nil {
		Error("Unable to get current user.")
		Crash(err)
	}
	
	return user.Username
}

// GetOperatingSystem gathers the users operating system from the runtime package and returns it as a string.
func GetOperatingSystem() string {
	// Returns Windows, Linux, or Darwin ( Apple )
	return runtime.GOOS
}

<<<<<<< HEAD
// Takes in a file name and runs the CheckFileExists function to check if it exists. 
// If so, it returns a warning statement stating that the file exists. 
// If it doesn't exist, the function will create the file.
// If the muted argument is toggled to false, it'll print a statement saying that the file was created. 
func MakeFile(file string, muted bool) {
	if CheckFileExists(file) {
		if !muted {
			Warning(file + " file exists.")
		}
		return
=======

// MakeFile takes in a collection of files and creates them in the current directory.
// Returns an error if unable to create a file.
func MakeFile(files ...string) error {
	if data, err := CheckFileExists(files...); err != nil {
		return err
>>>>>>> cleanup
	} else {
		for k, v := range data {
			if !v {
				if err := os.WriteFile(k, []byte(""), 0o666); err != nil {
					return err
				}
			}
		}

	}
	return nil
}

// MakeFolder takes in a collection of folder names and creates them. 
// Returns an error if unable to create the folder.
func MakeFolder(folders ...string) error {
	dirInfo, err := CheckFolderExists(folders...)
	if err != nil {
		return err
	}

<<<<<<< HEAD
	if dirInfo {
		if !muted {
			Warning(folder + " folder exists.")
		}
		return
	} else {
		err := os.Mkdir(folder, 0755)
		if err != nil {
			Error(err.Error())
			return
=======
	for k, v := range dirInfo {
		if !v {
			if err := os.Mkdir(k, 0755); err != nil {
				return err
			}
>>>>>>> cleanup
		}
	}

	return nil
}

// MoveFileToFolder takes two paths of files and moves the file from the oldPath to the newPath.
// Returns an error if unable to move files.
func MoveFileToFolder(oldPath string, newPath string) error {
	if err := os.Rename(oldPath, newPath); err != nil {
		return err
	}

	return nil
}

// DoesToolExist takes in a string that is a tool and checks the version silently to check if it's installed.
// Will return an error stating the tool is not installed.
func DoesToolExist(tool string) error {
	command := exec.Command(tool, "-v")

	if err := command.Run(); err != nil {
		return fmt.Errorf("%s %w.", tool, ErrToolNotInstalled)
	}

	return nil
}

// helpFlag creates a generic flag for a command to print out the commands help menu.
func HelpFlag() *argbin.Flag {
	return &argbin.Flag{
		Execute: func(ctx *argbin.Context) error {
			help := ctx.Command.HelpMenu
			fmt.Println(help)
			return nil
		},
		Terminal: true,
	}
}


// ToggleOutputForCMD makes the cmd's output route to stdout and stderr.
func ToggleOutputForCMD(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
}
