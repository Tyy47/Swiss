<<<<<<< HEAD
# Swiss 1.1a
=======
# Swiss 1.2

## Operation Cleanup:
I took a break from Swiss as the code base was getting unruly trying to manage all custom sections like built in arg parsing. During this break,
I developed clibox as a modular cli toolbox. These tools should cover a majority of swiss's architecture.

- Code is functionally complete
- Need to fill out code documentation
- Need to have commands route to stdout and stderr ( make wrapper function that takes in a command pointer and adds that command output to stdout ) 
- Remove panics from main.go and properly handle conditional errors 
>>>>>>> cleanup

- Add git undo command in shortcut module to undo commits and accept an optional arg to go back x amount of commits.
- Add Astro to web init

<<<<<<< HEAD
### Helps
- Create more descriptive help menus for each modules and indiviuial commands.
1. Each modules gets a better description of what each module achieves - DONE
2. Each command gets specific builtin documentation

### Dict 
1. Make command dictionaries look better when printed
=======
## Module testing checklist
- A checklist of uncomplete and completed modules that have gone through testing to make sure they're in working order

[x] Build
[x] Gen
[x] Init
[x] Shortcut

### Misc bugs found:
- Global help menu additions arent added 
- Git push shortcut command doesn't push; just adds & commits. 

## Init
- rewrite file & folder creation process. 

## Help Menu rewrite:
- Main: 
- Init:  
- Build:  
- Shortcut: 
- Gen: 

## Module Ideas:
- fmt: format text or code files using their respective tools.
- shorthand build & run: readd this lost functionality. ( Turn off TakesValue on run and add and create a shorthand function to search current directory and find the language.
>>>>>>> cleanup
