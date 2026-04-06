# Swiss 1.1.0

1.1.0:

### "Net" module:
-    Port scanner
-    IP Address Pinging
-    WHOIS Lookup
-    HTTP 'Probe': Similar to a curl request to grab http status code and http information.
-    Advanced Ping: Similar to the default systems ping but provides more information regarding to packets and other related network information.

### "Generator" module:
 -   UUID
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
Remove "return" on line 248 in main.go \
Remove copy.md \
Add zig support to build, run, and init \
Add port checking functionality to network module \
Add better documentation of functions and objects throughout the program. \
Start work on generator module \
Plan next module

Module Ideas:

Shortcuts? I.E `swiss sc git push "update: message"` will add all changed files to commit, add a message to the commit, then push.
