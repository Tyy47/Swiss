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
)

// Swiss version number
const VERSION_NUMBER string = "1.1.0"

// Global Arguments
var (
	Arguments           = gatherArgs()
	AdditionalArguments = gatherAdditionalArgs()
)

// Prints the Swiss version number to the console
func PrintVersionNumber() {
	fmt.Println("Swiss version number: " + ColorGreen + VERSION_NUMBER + ColorReset)
}

// The function displays a message stating Swiss has crashed and a message stating the crash reason
func Crash(err error) {
	Error("Swiss has crashed! View output below to learn more. If error is related to swiss and not to a missing dependency, put in a request on GitHub.")
	crashMessage(err)
	os.Exit(1)
}

// Checks if err has a value other then nil, if it does, it runs the Crash function to safely exit Swiss.
func CrashCheck(err error) {
	if err != nil {
		Crash(err)
	}
}

// A function that takes an initial command arg and a packed string of other arguments.
// It executes the command and returns the result for manual error handling depending on the circumstance.
func RunCommand(command string, arguments ...string) error {
	comm := exec.Command(command, arguments...)

	err := comm.Run()
	return err
}

// Gathers argument via the os library
func gatherArgs() []string {
	args := os.Args
	return args
}

// Gathers arguments past the 3rd index, for example swiss init rust -g. Only -g is caught by this function as it is the third index.
func gatherAdditionalArgs() []string {
	if len(Arguments) < 3 {
		return []string{}
	}

	return Arguments[3:]
}

// Checks if arguments are a certain length, if so, it grabs the requested index and returns the value of args[index].
func CheckArguments(args []string, length int, index int) string {
	if len(args) <= length {
		return ""
	} else {
		return args[index]
	}
}

// Checks if a file exists, if it does it returns true, if not, returns false.
func CheckFileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return err == nil && info.Mode().IsRegular()
}

// Checks if a folder exists and returns a boolean value depending on the outcome and an error if checking fails. 
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

// Gathers the users username and returns it
func GetUsersName() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("Unable to get current user: %s", err)
	}

	return user.Username
}

// Gathers and returns the user operating system.
func GetOperatingSystem() string {
	// Returns Windows, Linux, or Darwin ( Apple )
	return runtime.GOOS
}

// Takes in a file name and runs the CheckFileExists function to check if it exists. 
// If so, it returns a warning statement stating that the file exists. 
// If it doesn't exist, the function will create the file.
// If the muted argument is toggled to false, it'll print a statement saying that the file was created. 
func MakeFile(file string, muted bool) {
	if CheckFileExists(file) {
		Warning(file + " file exists.")
		return
	} else {
		err := os.WriteFile(file, []byte(""), 0o666)
		if err != nil {
			Crash(err)
			return
		}
	}
	if !muted {
		Success(file + " file created.")
	}
}

// Takes in a folder name and runs the CheckFolderExists function to check if it exists. 
// If so, it returns a warning statement stating that the folder exists. 
// If it doesn't exist, the function will create the folder.
// If the muted argument is toggled to false, it'll print a statement saying that the folder was created. 
func MakeFolder(folder string, muted bool) {
	dirInfo, err := CheckFolderExists(folder)

	if err != nil {
		Error("Unable to create directory.")
		return
	}

	if dirInfo {
		Warning(folder + " folder exists.")
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

// Takes in two paths, an old path argument that holds the current path of the file you're trying to move.
// The new path is the location you're moving the file to.
// The muted argument allows you to toggle the moved message statement.
func MoveFileToFolder(oldPath string, newPath string, muted bool) {
	if err := os.Rename(oldPath, newPath); err != nil {
		Error("Unable to move file: " + err.Error())
		return
	}

	if !muted {
		Success("Successfully moved " + oldPath + " to " + newPath + ".")
	}
}

// Takes in a tool string and runs the version command on that tool.
// If it executes with no errors the function returns true.
// If the tool is not installed or path'd correctly, it will print out an error statement and return false.
func DoesToolExist(tool string) bool {
	command := exec.Command(tool, "-v")

	if err := command.Run(); err != nil {
		Error(tool + " is not installed or added to path.")
		return false
	}

	return true
}

// Prints the prompt provided and asks the user for an input. If an input is not provided or the scanner has failed. The fallback string will be returned instead.
func GetUserInput(prompt string, fallback string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)

	if scanner.Scan() {
		input := scanner.Text()
		return input
	}

	return fallback
}
