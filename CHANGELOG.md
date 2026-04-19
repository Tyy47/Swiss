# Swiss 1.1.0

## New Modules:

### Shortcuts:
#### The new shortcut module for Swiss provides a set of commands that combine multiple commands into one to save typing time.
- Commit - The commit shortcut combines git add . & git commit -m "commit message here".
- Push - The push shortcut combines git add . & git commit -m "" & pushes your commits to the added repository.
- Synhc - The sync shortcut combines git fetch & git status.

### Net:
#### The net module provides of set of networking tools to gather information from domains and IP addresses.
- Connect: Connect to a domain/IP address to get back an HTTP response code.
- Addr: Displays IPv4 and 6 addresses of a given domain.
- NS: Displays the nameservers of a given domain
- CNAME: Displays CNames records based off a given domain
- TXT: Displays TXT records of a domain
- MX: Displays MX records of a domain.
- Gather: Gathers data from all available net functions and compiles them into a file for later viewing.

### Gen:
#### The gen module provides functions to generate codes.
- UUID: Generates an 128 bit number in a UUID format
- Secret: Generates a randomized string based on a given length input, if no input is provided, the default length will be 16 characters.


## Existing Module updates:

### Init
- Added support for zig
- Added support for python using uv
- Added jujutsu vc support for starting a project
- Reworked web init. Now able to specify the framework when building the web project.

### Build
- Added support for zig

### Run
- Added support for python

### Misc
- Rebuilt arg parsing system to work on a registry storage system. This now allows for more flexible commands and better growth structurally.
- Net module documentation added to DOCS.md
- Gen module documentation added to DOCS.md
- Shortcut module documentation added to DOCS.md
- Added jujutsu startup to DOCS.md
- Made Readme look nicer
