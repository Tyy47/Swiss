# Swiss 1.2

## Operation Cleanup:
I took a break from Swiss as the code base was getting unruly trying to manage all custom sections like built in arg parsing. During this break,
I developed clibox as a modular cli toolbox. These tools should cover a majority of swiss's architecture.

- Code is functionally complete
- Need to fill out code documentation
- Need to have commands route to stdout and stderr ( make wrapper function that takes in a command pointer and adds that command output to stdout )
- Remove panics from main.go and properly handle conditional errors 

