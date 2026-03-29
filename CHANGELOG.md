# Swiss 1.0.2


## Init Module

When initializing a C, HTML, or Go project. Git will initialize as well.
When initing a C project. main.c will be moved to the src/ directory when created.
Removed automatic git init when initializing a project via Swiss. To have git created alongside your project, use the -g or --git after the language name to create a .git repo.

## Command Dictionary Module

Added Docker commands to dictionary
Added Git commands to dictionary

## Create Module
Made create more dynamic when creating folders and files. Before you had to write a new line of swiss to either create files or folders, now you can write in one line, example below.

Before: `swiss create file "test1.md" "test2.md"; swiss create folder "foo" "bar"`
After: `swiss create file "test3.md" "test4.md" folder "foo" "bar"`

You can also string them together like so:
`swiss create file "test3.md" "test4.md" folder "foo" "bar" file "test5.md", "test6.md"`

## Misc

Add function to move files for structuring projects

## Bugs Fixes

Fixed issue where main.c file was not being generated when initing C project
