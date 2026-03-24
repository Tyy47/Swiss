<h1 align="center">Swiss Documentation</h1>

## Commands

### Misc
A collection of misc commands

`swiss` or `swiss -h` will print out the help menu. \
`swiss -v` or `swiss version` will print out the version name. \

### Build
The Build module. Build/Compile or Run programs based on the language provided and that is available via Swiss.

`swiss -h` or `swiss --help` to print out the build help menu. \
`swiss build -l` or `--list` will print out a list of languages that can be built/compiled via Swiss.
`swiss build <language : string>` will build/compile the program based on the language provided and that is available via Swiss. \
`swiss run <language : string>` will run the program based on the language provided and that is available via Swiss. \

### Create
The Create module. Create files and or folders through Swiss on mass.

`swiss create -h` or `--help` will print out the help menu for the Create module \
`swiss create <file | folder> <names: string>` Creates folders or files based on the 3rd argument. Names for files must be provided in order to create them. Supports file extensions and creation of multiple files/folders at once

### Init
The Init module. Init a project using Swiss with the available project templates available.

`swiss init -h` or `--help` to print out the help menu for the Init module \
`swiss init -l` or `--list` to print out a list of languages Swiss can initialize projects for \
`swiss init <language : string> [project name : string]` Inits a project based on the language provided and the project name for certain languages.
