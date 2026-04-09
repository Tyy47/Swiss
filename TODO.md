# Swiss 1.1.0

1.1.0:

### "Net" module:
-    Port scanner
-    IP Address Pinging
-    WHOIS Lookup
-    HTTP 'Probe': Similar to a curl request to grab http status code and http information.
-    Advanced Ping: Similar to the default systems ping but provides more information regarding to packets and other related network information.

### "Generator" module:
 -   UUID: Random numbers and letters in a certain order 12 - 4 - 4 - 4 - 12
 -   Hex Codes ( Based on main color input or random otherwise )
 -   "Secret" codes, strings of letters, numbers and symbols created based on users number. 

Move modules over to new arg parsing system

BUILD - DONE \
RUN - DONE \
DICT - DONE \
INIT - DONE \
NET - DONE \
CREATE - DONE \
UPDATE - DONE

04/06/26:

Update Changelog.md - DONE \
Update DOCs.md with net module documentation - DONE \
Update version number - DONE \
Remove "return" on line 248 in main.go - DONE \
Remove copy.md - DONE
Add zig support to build, run, and init - DONE \

04/07/26:

Add port checking functionality to network module - DONE \
Add better documentation of functions and objects throughout the program. \
Start work on generator module \
Plan next module

Module Ideas:

Shortcuts? I.E `swiss sc git push "update: message"` will add all changed files to commit, add a message to the commit, then push. \
Web init? I.E `swiss init web -s or -r or -v` will init a project using vite and the flags will make a Svelte, React, or Vanilla web based project.

vanilla - bun create vite my-vanilla-app --template vanilla-ts
react - bun create vite my-react-app --template react-ts
svelte - bun create vite my-svelte-app --template svelte-ts

The above is done but i'm thinking of refactoring the init module. It wasn't built with the current arg parser in mind and needs to be rebuilt.

### Init 2.0:

Init struct:
- Name of language - Used to name the struct in order to locate it by matching against the users arguments.
- Build tool - Used to build the language when running a command.
- Init Times - Tracks if the language has been init'd in the same line of arguments, if its equal to one then don't execute the handler function and move to the next language.
- Flags - Contains the arguments that it needs in order to build.
- Additional Flags - Contains additional optional or needed flags in order to create a project. I.E swiss init web -S will init a svelte project and -R a react project
- Compiler handler function or Method - Each init'd project struct handles its own project creation.
- Manual init bool flag - Checks if a program needs to be manually init'd.
- Manual Files - An array of file names that will get created during manual init.
- Manual Folders - An array of folder names that will get created during manual init.

Then, when a language is called via arguments, it'll call its own handler function or method to execute its init process rather then the current buggy setup calling > 1 times when initing a project.
