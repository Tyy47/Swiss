package utils

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"

	"github.com/Tyy47/clibox/outbin"
)

var (
	

	// util errors
	ErrUnableToCreateFile = errors.New("unable to create file in current directory")
	ErrToolNotInstalled = errors.New("is not installed")
)

var output = outbin.NewOutput(os.Stdout, os.Stderr)

// CheckFileExists takes in a list of file names and checks if they exist.
// Results are added to a map to iterate on, an error is returned if the stats of a file cannot be retrieved.
func CheckFileExists(files ...string) (map[string]bool, error) {
	
	fileMap := make(map[string]bool)

	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			return nil, err
		}

		if info.Mode().IsRegular() {
			fileMap[file] = true
		} else {
			fileMap[file] = false
		}

	}

	return fileMap, nil
}

// CheckFolderExists takes in a list of folder names and iterates over them to see if they exist.
// Results are added to a map to iterate upon and an error is returned if there is a problem
// gathering the folders.
func CheckFolderExists(folders ...string) (map[string]bool, error) {
	
	// Create the map
	folderMap := make(map[string]bool, 0)

	// Loop over all the folders given
	for _, folder := range folders {
		
		// Gather folder stats
		info, err := os.Stat(folder)
		if err != nil {
			return nil, err
		}

		// Directory check
		if info.IsDir() {
			// Folder is a directory
			folderMap[folder] = true
		} else {
			// Folder item is not a directory
			folderMap[folder] = false
		}
	}

	// Return the map
	return folderMap, nil
}

// GetUsersName gathers the user's username from the user package and returns it.
func GetUsersName() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to get current user: %s", err)
	}

	return user.Username
}

// GetOperatingSystem gathers the users operating system from the runtime package and returns it as a string.
func GetOperatingSystem() string {
	// Returns Windows, Linux, or Darwin ( Apple )
	return runtime.GOOS
}


// MakeFile takes in a collection of files and creates them in the current directory.
// Returns an error if unable to create a file.
func MakeFile(files ...string) error {
	if data, err := CheckFileExists(files...); err != nil {
		return err
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

	for k, v := range dirInfo {
		if !v {
			if err := os.Mkdir(k, 0755); err != nil {
				return err
			}
		}
	}

	return nil
}

// Takes in two paths, an old path argument that holds the current path of the file you're trying to move.
// The new path is the location you're moving the file to.
// The muted argument allows you to toggle the moved message statement.
func MoveFileToFolder(oldPath string, newPath string) error {
	if err := os.Rename(oldPath, newPath); err != nil {
		return err
	}

	return nil
}

// Takes in a tool string and runs the version command on that tool.
// If it executes with no errors the function returns true.
// If the tool is not installed or path'd correctly, it will print out an error statement and return false.
func DoesToolExist(tool string) error {
	command := exec.Command(tool, "-v")

	if err := command.Run(); err != nil {
		return fmt.Errorf("%s %w.", tool, ErrToolNotInstalled)
	}

	return nil
}
