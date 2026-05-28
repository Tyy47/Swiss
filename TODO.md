# Swiss 1.1a

- Add git undo command in shortcut module to undo commits and accept an optional arg to go back x amount of commits.
- Add Astro to web init

### Helps
- Create more descriptive help menus for each modules and indiviual commands.
1. Each modules gets a better description of what each module achieves
2. Each command gets specific builtin documentation

- Make helps menus its own module and import it into main instead of having all the helps live in utils.
- Create structs to assign modules (sc, net, etc) and the Command (sync, addr, etc)
- Turn help into a more structured command rather then a help menu for all commands. Instead swiss help sc will print the help menu for sc and swiss help sc sync will pull up all the parameters and additional details of the provided command.
