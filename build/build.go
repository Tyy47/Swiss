package build

import (
	"errors"
	"fmt"
	"os/exec"
<<<<<<< HEAD
	"os/signal"
	"path/filepath"
=======
>>>>>>> cleanup
	"strings"
	"swiss/utils"

	"github.com/Tyy47/clibox/argbin"
)

<<<<<<< HEAD
const buildProgramList = `Rust: Cargo
C: Clang
Go: Go
Zig: Zig
Web: Bun/Vite`
=======
var (
>>>>>>> cleanup

	// generic errors

	ErrUnknownLanguage = errors.New("unknown language")

	// build object errors

	ErrNilLang      = errors.New("lang argument cannot be nil")
	ErrNilBuildTool = errors.New("build tool argument cannot be nil")
	ErrNilBuildArgs = errors.New("build args argument cannot be nil")
	ErrNilBuildFile = errors.New("build file argument cannot be nil")

	// program type errors

	ErrNilProgram = errors.New("program return cannot be nil")
	ErrNilBuild = errors.New("build in program type cannot be nil")
	ErrNilRun = errors.New("run in program type cannot be nil")
)

// Language safely adds language guard rails
type Language string

// Available languages for swiss to build & run
const (
	Go   Language = "go"
	Rust Language = "rust"
)

// build stores all the needed information to build a program
type build struct {
	Tool           string   // Tool stores the build tool of a language (e.g cargo, bun, etc)
	BuildArguments []string // BuildArguments stores optional flags for the build Tool
	BuildFile      string   // File thats associated with a certain language I.E main.go, Cargo.toml, etc.
}

// run stores all the needed information to run a program
type run struct {
	Tool           string   // Tool stores the run tool of a language (e.g cargo, bun, etc)
	BuildArguments []string // BuildArguments stores optional flags for the run Tool
	BuildFile      string   // File thats associated with a certain language I.E main.go, Cargo.toml, etc.
}

// program stores context between functions to get build and run objects for later function executions
type program struct {
	Build *build
	Run *run
}

// buildMap contains all of the available languages that can be built via swiss.
var buildMap = map[Language]build{
	Go: {
		Tool:           "go",
		BuildArguments: []string{"build"},
		BuildFile:      "go.mod",
	},

	Rust: {
		Tool:           "cargo",
		BuildArguments: []string{"build"},
		BuildFile:      "Cargo.toml",
	},
}

<<<<<<< HEAD
func (b *buildRegistry) addToBuildRegistry(newBuild ...build) {
	b.builds = append(b.builds, newBuild...)
}

func (b *build) Initialize() error {
	command := exec.Command(b.Tool, b.Arguments...)
=======
// runMap contains all of the available languages that can be ran via swiss
var runMap = map[Language]run{
	Go: {
		Tool:           "go",
		BuildArguments: []string{"run", "main.go"},
		BuildFile:      "go.mod",
	},

	Rust: {
		Tool:           "cargo",
		BuildArguments: []string{"run"},
		BuildFile:      "Cargo.toml",
	},
}

// getLanguage takes a language input and returns a build object thats related to that language.
// If no languages are found, ErrUnknownLanguage is returned.
func getLanguage(lang string) (*program, error) {
	// Nil check for lang
	if lang == "" {
		return nil, ErrNilLang
	}
>>>>>>> cleanup

	// Lowercase the input
	fixed := strings.ToLower(lang)

	// Gather the build object
	buildObject, ok := buildMap[Language(fixed)]
	if !ok {
		return nil, ErrUnknownLanguage
	}

	// Gather the run object
	runObject, ok := runMap[Language(fixed)]
	if !ok {
		return nil, ErrUnknownLanguage
	}

	// Create a program and assign the build and run objects to store for context
	langs := program{
		Build: &buildObject,
		Run: &runObject,
	}

	// Return the build object
	return &langs, nil
}

// buildLanguage handles the execution of the build of a program.
func buildLanguage(p *program) error {
	// Nil check program
	if p == nil {
		return ErrNilProgram
	}

	// Nil check for build
	if p.Build == nil {
		return ErrNilBuild
	}

	// Assigning build to b
	b := p.Build

	// Takes in the tool and arguments to create a command
	cmd := exec.Command(b.Tool, b.BuildArguments...)

	// Assigns cmd's output to stdout and stderr
	utils.ToggleOutputForCMD(cmd)

	// Runs the command and returns the error if it fails
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

<<<<<<< HEAD
func buildRustProject() build {
	return build{
		Language:  "rust",
		Tool:      "cargo",
		Arguments: []string{"build", "--release"},
		BuildFile: "Cargo.toml",
	}
}

func buildGoProject() build {
	return build{
		Language:  "go",
		Tool:      "go",
		Arguments: []string{"build"},
		BuildFile: "main.go",
	}
}

func buildCProject() build {
	return build{
		Language:  "c",
		Tool:      "clang",
		Arguments: []string{"main.c", "-Wall", "-Wextra", "-Wpedantic", "-Werror", "-g", "-o", "main"},
		BuildFile: "main.c",
	}
}

func buildZigProject() build {
	return build{
		Language:  "zig",
		Tool:      "zig",
		Arguments: []string{"build"},
		BuildFile: "main.zig",
	}
}

func buildWebProject() build {
	return build{
		Language: "web",
		Tool: "bun",
		Arguments: []string{"run", "build"},
		BuildFile: "package.json",
	}
}

func scanForBuildFiles() (bool, build) {
	// Scan directory for all files.
	files, err := os.ReadDir(".")
	if err != nil {
		utils.Error("Unable to read files in current directory.")
		return false, build{}
	}

	for _, project := range registry.builds {
		for _, file := range files {
			if file.Name() == project.BuildFile {
				return true, project
			}
		}
	}

	utils.Warning("Unable to find inputted language, check language list for buildable languages via Swiss.")
	return false, build{}
}

func BuildProject() {
	// Grabs the bool and build struct from scanForFiles()
	result, project := scanForBuildFiles()

	if !result {
		utils.Error("Unable to build project, check inputted language to see if it's in Swiss build list.")
		return
	}

	if err := project.Initialize(); err != nil {
		utils.Error("Unable to compile project.")
		utils.Reason(err.Error())
		return
	}
	utils.Success(project.Language + " project has been compiled.")
}

func HandleBuildInput() {
	if len(utils.Arguments) < 3 {
		return
	}

	argument := utils.Arguments[2]
	argument = strings.ToLower(argument)

	for build := range len(registry.builds) {
		if argument == registry.builds[build].Language {
			if err := registry.builds[build].Initialize(); err != nil {
				utils.Error("Unable to compile project.")
				utils.Reason(err.Error())
				return
			} else {
				utils.Success(registry.builds[build].Language + " project has been compiled.")
				return
			}
		}
	}
	utils.Error("Unable to find " + argument + " in registry list.")
}

func SwissInstall() {
	system := utils.GetOperatingSystem()
	if system == "linux" {
		// Find the Go build configuration explicitly rather than relying on registry order.
		var goBuild *build
		for i := range registry.builds {
			if registry.builds[i].Language == "go" {
				goBuild = &registry.builds[i]
				break
			}
		}
		if goBuild == nil {
			utils.Error("Go build configuration not found in registry.")
			return
		}
		if err := goBuild.Initialize(); err != nil {
			utils.Error("Unable to build Swiss for install.")
			utils.Crash(err)
			return
		}

		command := exec.Command("mv", "swiss", "/home/"+utils.GetUsersName()+"/.local/bin/")

		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		if err := command.Run(); err != nil {
			utils.Error("Swiss install failed. Check output above.")
			return
		}

		utils.Success("Swiss installed successfully! Use 'swiss' in the terminal to gain access to the program.")
	} else {
		utils.Warning("Swiss install is not supported for " + system + ".")
		return
	}
}

func UpdateSwiss(args *[]string) {
	// Unstable check
	installSwissUnstable := false

	for _, toggle := range *args {
		if toggle == "-u" || toggle == "--update" {
			installSwissUnstable = true
			break
		}
	}

	// Make directory to clone into
	utils.MakeFolder("swiss_install", true)

	// Resolve an absolute path so cleanup is unaffected by later os.Chdir calls.
	installPath, err := filepath.Abs("swiss_install")
	if err != nil {
		utils.Error("Unable to resolve install path.")
		utils.Crash(err)
		return
	}

	// Cleans up swiss_install directory if the update process is interrupted.
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		utils.Reason("\nSwiss install interrupted, cancelling and cleaning up install files.")
		if err := os.RemoveAll(installPath); err != nil {
			utils.Error("Unable to clean up install files: " + err.Error())
		}
		os.Exit(1)
	}()

	// Clone the repository
	clone := exec.Command("git", "clone", "https://github.com/Tyy47/Swiss.git", "swiss_install/")
	
	// If there is an error in cloning the repository, it will crash and return an error statement.
	if err := clone.Run(); err != nil {
		utils.Error("Unable to clone Swiss repo. Install manually or create a bug report on the repository.")
		utils.Crash(err)
		return
	}

	// Change directory into cloned repo
	if err := os.Chdir("swiss_install"); err != nil {
		utils.Error("Unable to change directory into swiss_install. Exiting.")
		utils.Crash(err)
		return
	}
	
	// If the unstable toggle is true, it will switch to the unstable branch and install the unstable version of Swiss.
	if installSwissUnstable {
		if err := utils.RunCommand("git", "switch", "unstable"); err != nil {
			utils.CrashCheck(err)
		}

		utils.Note("Installing Swiss unstable.")
	}

	// Prompt the user to either go install or move to local/bin
	utils.Note("Select the number associated with the option in order to continue.")
	fmt.Println("How would you like to install Swiss?")
	fmt.Println("1. Go Install\n2. Move to local/bin ( Linux only )")
	for {
		var userInput string
		fmt.Scanln(&userInput)

		switch userInput {
		case "1":
			// Go install here
			install := exec.Command("go", "install")

			if err := install.Run(); err != nil {
				utils.Error("Unable to install Swiss using Go Install.")
				utils.Crash(err)
				break
=======
// runLanguage handles the "running" of a program.
// It will take a run object from program and gather all the information required
// to run the program.
func runLanguage(p *program) error {
	// Nil check for program
	if p == nil {
		return ErrNilProgram
	}

	// Nil check for stored build
	if p.Build == nil {
		return ErrNilRun
	}

	// Assigning run to r
	r := p.Run

	// Takes in the tool and arguments to create a command
	cmd := exec.Command(r.Tool, r.BuildArguments...)

	// Assigns cmd's output to stdout and stderr
	utils.ToggleOutputForCMD(cmd)

	// Runs the command and returns the error if it fails
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// listFlag returns a flag for build and run to print the list of valid languages
// that swiss can build and run.
func listFlag() *argbin.Flag { 
	return &argbin.Flag{
		Execute: func(ctx *argbin.Context) error {
			fmt.Println("Buildable and runable languages:")

			// Loop over build map and print entries
			for k, v := range buildMap {
				fmt.Printf("%s: required tool = %s.\n", k, v.Tool)
>>>>>>> cleanup
			}

			return nil
		},
		Terminal: true,
	}
}

<<<<<<< HEAD
func init() {
	buildList := []build{
		buildGoProject(),
		buildRustProject(),
		buildCProject(),
		buildZigProject(),
		buildWebProject(),
	}

	registry.addToBuildRegistry(buildList...)
=======
// SwissInstall installs swiss for Linux users when ran inside of the cloned swiss repo.
func SwissInstall() *argbin.Command {
	return &argbin.Command{
		Name: "install",
		Execute: func(ctx *argbin.Context) error {
			// Check if the user is running linux
			if utils.GetOperatingSystem() != "linux" {
				utils.Output.Warning("unable to install swiss. unsupported operating system")
				return nil
			}
			
			// Install message
			utils.Output.Info("installing swiss.")

			// Get language
			lang, err := getLanguage("go")
			if err != nil {
				return err
			}

			// Build swiss
			if err := buildLanguage(lang); err != nil {
				return err
			}

			// Create destination path
			desPath := fmt.Sprintf("/home/%s/.local/bin/swiss", utils.GetUsersName())

			// Build the move command
			moveCmd := exec.Command("mv", "swiss", desPath)

			// Execute the move command
			if err := moveCmd.Run(); err != nil {
				return err
			}

			// Success message after install
			utils.Output.Success("swiss has been installed! make sure .local/bin is added to your path")

			return nil
		},
	}
}

// BuildCommand creates the "build" command for swiss.
// Command is added in main.go in main function.
func BuildCommand() *argbin.Command {
	return &argbin.Command{
		Name:        "build",
		TakesValue:  true,
		Execute: func(ctx *argbin.Context) error {
			utils.Output.Info("searching for language.")
			b, err := getLanguage(ctx.ParsedValue)
			if err != nil {
				return err
			}

			utils.Output.Info("building program.")
			if err := buildLanguage(b); err != nil {
				return err
			}

			utils.Output.Successf("%s program has been built", ctx.ParsedValue)
			return nil
		},
		Flags: argbin.Flags{
			"-h" : utils.HelpFlag(),
			"--help": utils.HelpFlag(),
			"-l": listFlag(),
			"--list": listFlag(),
		},
		HelpMenu: `
╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
Build & Run module - Builds or Runs a program based on the language inputted.

Commands: 
	build <string>: Builds a program based on the language you input.
	run <string>: Runs a program based on the language you input.
Flags:
	-h --help: Opens the help menu.
	-l --list: Prints a list of available languages to build and run with their respective build tools available in Swiss.`,
	}
}

// RunCommand create the "run" command for swiss.
// Command is added in main.go in main function.
func RunCommand() *argbin.Command {
	return &argbin.Command{
		Name: "run",
		TakesValue: true,
		Execute: func(ctx *argbin.Context) error {
			utils.Output.Info("searching for language.")
			r, err := getLanguage(ctx.ParsedValue)
			if err != nil {
				return err 
			}
			
			utils.Output.Info("running program.")
			if err := runLanguage(r); err != nil {
				return err
			}

			utils.Output.Successf("%s program has been ran.", ctx.ParsedValue)
			return nil
		},
		Flags: argbin.Flags{
			"-h" : utils.HelpFlag(),
			"--help": utils.HelpFlag(),
			"-l": listFlag(),
			"--list": listFlag(),
		},
		HelpMenu: `
╭───────────────────  Swiss  ────────────────────╮
│                                                │
│       The army knife of CLI applications       │
│                                                │
╰────────────────────────────────────────────────╯
Build & Run module - Builds or Runs a program based on the language inputted.

Commands: 
	build <string>: Builds a program based on the language you input.
	run <string>: Runs a program based on the language you input.
Flags:
	-h --help: Opens the help menu.
	-l --list: Prints a list of available languages to build and run with their respective build tools available in Swiss.`,
	}
>>>>>>> cleanup
}
