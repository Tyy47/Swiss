# Swiss 1.1.0

- Do some module clean up, add documentation to functions, prep for 1.1.0 release.

- When initing a web project using git, when it creates files it doesnt place them into the inited project folder. need to change dir project folder.
    1. The problem is it's initing git outside of the project folder. Need to find a way to get the project name when a web app project has started then pass it into git init to change directory before initing the repository.
