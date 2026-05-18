# Swiss 1.1.1

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
