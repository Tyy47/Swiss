# Swiss 1.1.0

- Make Build and Run dynamic.
1. Get Rid of argument parsing
2. Make build and run a singular command that runs a build function
3. Rebuild Handle input function to scan for files, detect a certain build file (main.go, cargo.toml, etc.) to build the program.
4. Do some bug testing to see if the arg parser gets screwy with the -h & -l flags when passed in.
