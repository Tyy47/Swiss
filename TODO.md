# Swiss 1.1.0

- Do some module clean up, add documentation to functions, prep for 1.1.0 release.



## Single call build and run

1. Add else if clause for if command.Handler != nil.
2. Make an if clause checking if SingleRun is true AND if len(utils.Arguments) == 2; If so, run the handler which will be the scanning and building function.
3. After creating and testing this for build, implement it for run as well.


## Bugs

- Create module doesn't work - Need to figure out why. Got it to work but its duplicating it's output twice when running a command. - Might remove create module as system commands could cover this better.
