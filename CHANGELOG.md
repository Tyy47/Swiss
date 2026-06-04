# Swiss 1.1a

## Version Number Change:
Changed version numbering to two numbers and letters. The numerarical increases will signify a large patch for Swiss while lettering will indicate smaller patches related to bug fixes or smaller feature changes.

## Initialize
- Extended gits init process to now add a remote repository link.

## Build:
- Added web support to build module by using bun & vite.

## Run:
- Added run support to Typescript CLI apps (Main file must be named main.ts)

## Shortcuts
- Added optional argument to the sync shortcut to also pull changes using -p or --pull.

## Misc:
- Changed text color implementation to easily change colors of text in code.
- Rewrote help menus to be more informative of each command and module.

## Bug Fixes:
- Running `swiss run` with no language listed would show the wrong list of supported languages.
- Running or building a project would show a success message even if the run or build actually failed.
- Installing Swiss could crash if no build configurations were loaded.
- A network error that occurred while saving data to a file would be silently ignored instead of reported.
- The network output file was not being properly closed after data was written to it.
- Initializing a project into a folder that didn't exist would cause git commands to run in the wrong place.
- The initial git commit created during project setup had quotes around the message, making it look like `"initial"` instead of `initial`.
- Running `swiss init web` without specifying a framework would crash instead of defaulting to vanilla.
- Running `swiss sync --pull` would show a success message even if the pull failed.
- Building or running a project with an error would crash the program instead of showing a helpful error message.
- `swiss install` relied on registry index order to find the Go build, which could use the wrong configuration if the registry order changed.
- Cancelling `swiss update` mid-process could fail to clean up temporary files if the working directory had changed.
- `swiss gen secret` would crash when given a non-numeric or negative length instead of showing a validation error.
- Initialization flags like `-g` and `-j` could not be combined — only the first flag would take effect.
- Running `swiss init` with git and leaving the remote URL blank would attempt to add an empty remote instead of skipping gracefully.
- Passing an unknown subcommand to any command was silently ignored with no feedback to the user.
- Network lookup failures for IP addresses, nameservers, CNAME, TXT, and MX records would crash the entire program instead of displaying an error.
- The network connection's close call was deferred before the connection was confirmed open and placed after code that already used the connection.
- `GetUsersName` used a different logging method than the rest of the project, producing inconsistent error output.
- Commands executed through `RunCommand` had their stderr output silently discarded instead of being shown to the user.
- `MakeFile` and `MakeFolder` showed "already exists" warnings even when the caller requested silent operation.
