# Swiss 1.0.2


## Init Module

When initializing a C, HTML, or Go project. Git will initialize as well.
When initing a C project. main.c will be moved to the src/ directory when created.
Removed automatic git init when initializing a project via Swiss. To have git created alongside your project, use the -g or --git after the language name to create a .git repo.

## Command Dictionary

Added Docker commands to dictionary
Added Git commands to dictionary

## Misc

Add function to move files for structuring projects

## Bugs Fixes

Fixed issue where main.c file was not being generated when initing C project
