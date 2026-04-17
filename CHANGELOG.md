# Swiss 1.1.0

## New Modules:

### Shortcuts:
#### The new shortcut module for Swiss provides a set of commands that combine multiple commands into one to save typing time.
- Commit - The commit shortcut combines git add . & git commit -m "commit message here".
- Push - The push shortcut combines git add . & git commit -m "" & pushes your commits to the added repository.

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
- Added jujutsu startup to DOCS.md
- Made Readme look nicer
