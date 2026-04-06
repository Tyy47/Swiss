# Swiss 1.1.0

1.1.0: \
    1. "Net" module: \
        1.1: Port scanner \
        1.2: IP Address Pinging \
        1.3: WHOIS Lookup \
        1.4: HTTP 'Probe': Similar to a curl request to grab http status code and http information. \
        1.5: Advanced Ping: Similar to the default systems ping but provides more information regarding to packets and other related network information. \
    2. "Generator" module \
        2.1: UUID \
        2.2: Hex Codes ( Based on main color input or random otherwise ) \
        2.3: "Secret" codes, strings of letters, numbers and symbols created based on users number. 



Command struct
    -- Name string
    -- Subcommands []Subcommand


Subcommand struct 
    -- Name ( flag or argument )
    -- Handler func(x y or z)
    -- Map of flags and functions. Have the keys be the flags and the values be calls to the functions


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
Remove copy.md \
Add zig support to build, run, and init \
Add port checking functionality to network module \
Add better documentation of functions and objects throughout the program. \
Start work on generator module \
Plan next module

Module Ideas:
Shortcuts? I.E `swiss sc git push "update: message"` will add all changed files to commit, add a message to the commit, then push.
