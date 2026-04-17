# Swiss Shortcut Module Ideas

## Git Workflow Chains

### Basic Git Chaining
```bash
swiss sc git push "feat: commit message"    # git add .; git commit -m "..."; git push
swiss sc git sync "hotfix:"                  # git add .; git commit -m "..."; git push; git pull
swiss sc git commit "update"                 # git add .; git commit -m "update"
```

### Undo & Recover
```bash
swiss sc git undo                            # git reset --soft HEAD~1
swiss sc git revert "commit-msg"             # git revert --no-edit
swiss sc git status                          # git status with branch info
```

### Branch Management
```bash
swiss sc git checkout "feature-branch"       # git checkout -b feature-branch
swiss sc git switch "release"                # git switch -c release
swiss sc git update                          # git fetch upstream && git merge upstream/main
```

### Advanced Git Operations
```bash
swiss sc git clean                           # git clean -fd
swiss sc git sync-all                        # git push origin main && git push origin --all
swiss sc git cherry-pick "commit-hash"      # git cherry-pick commit-hash
```

---

## Docker Quick Actions

### Container Management
```bash
swiss sc docker logs "container-name"        # docker logs -t container-name | tail -f
swiss sc docker clean                        # docker container prune -f
swiss sc docker restart "container"          # docker restart container-name
swiss sc docker exec "shell"                 # docker exec -it container-name /bin/bash
```

### Image Management
```bash
swiss sc docker images "filter"              # docker images | grep filter
swiss sc docker build "context"              # docker build -t image-name .
swiss sc docker push "image-name"            # docker push image-name
swiss sc docker clean-images                 # docker image prune -f
```

### Docker Compose
```bash
swiss sc docker up "service"                 # docker-compose up -d service
swiss sc docker stop "service"               # docker-compose stop service
swiss sc docker restart-compose              # docker-compose restart
```

---

## Development Efficiency

### Build & Run Chains
```bash
swiss sc dev cycle                           # build → test → format → lint
swiss sc dev quickrun                        # build && run
swiss sc dev test-run                        # go test ./... && echo "tests passed"
swiss sc dev build-fast                      # build with minimal dependencies
```

### Code Quality
```bash
swiss sc quality check                       # gofmt -s -w . && go vet ./...
swiss sc review "branch"                     # git diff main...branch
swiss sc style fmt                           # run format and lint tools
```

### Deployment
```bash
swiss sc deploy "production"                 # build && commit && push && notify
swiss sc deploy "staging"                    # build && push && docker push
swiss sc dev watch                           # run in dev mode with auto-reload
```

---

## System Management

### System Information
```bash
swiss sys info                               # OS details, RAM, CPU usage
swiss sys health                             # check disk, memory, uptime
swiss sys ports                              # list open ports and services
swiss sys network                            # show network interfaces
```

### System Maintenance
```bash
swiss sys tidy                               # go mod tidy, cargo clean, npm cache clean
swiss sys backup "backup"                    # copy current directory to backup-folder with timestamp
swiss sys snapshot 'description'             # create system snapshot
swiss sys restore "snapshot-id"              # restore system from snapshot
```

### Resource Management
```bash
swiss sys memory                             # show memory usage
swiss sys disk                               # show disk usage by folder
swiss sys log monitor                        # monitor system logs for errors
swiss sys backup "backup"                    # create incremental backup
```

---

## Project Management

### Project Setup
```bash
swiss project setup                          # init repo + create folders
swiss project quickstart "type"              # initialize project with defaults
swiss project config "settings"              # load project configuration
```

### Source Control
```bash
swiss project release "v1.0.0"               # create tag and push
swiss project push                            # push all branches and tags
swiss project pull                            # fetch and rebase all remotes
swiss project status                          # show project health and sync status
```

### Team & Deployment
```bash
swiss project deploy "server"                # build, commit, push, ssh deploy
swiss project notify "on-merge"               # send notification on merge
swiss project backup                          # backup project configuration
swiss project sync                            # sync all project files
```

---

## File System Power

### Bulk Operations
```bash
swiss fs bulk "pattern"                      # find files matching pattern, then bulk action
swiss fs watch "file"                        # watch file for changes
swiss fs backup "backup"                     # copy current directory to backup-folder with timestamp
swiss fs sync "destination"                  # sync files to destination folder
```

### File Operations
```bash
swiss fs find "keyword"                      # grep files containing keyword
swiss fs search "extension"                  # find all files with specific extension
swiss fs tree "directory"                    # display directory tree structure
swiss fs diff "folder1" "folder2"            # compare two folders
```

### Directory Operations
```bash
swiss fs mkdir "nested"                      # create nested directories
swiss fs clean "temp"                        # remove temporary files
swiss fs organize                             # organize files by type
swiss fs backup "backup-name"                # recursive backup to external drive
```

---

## Quick Action Commands

### Development Tools
```bash
swiss dev test                              # run all tests
swiss dev fmt                              # format code
swiss dev lint                             # run linting
swiss dev typecheck                         # type check code
```

### Environment Management
```bash
swiss env activate "dev"                    # activate development environment
swiss env setup "production"                # setup production configuration
swiss env switch "env-name"                 # switch between environments
swiss env backup "env"                      # backup current environment
```

### Database Operations
```bash
swiss db seed                              # run database seeds
swiss db migrate "up"                       # run database migrations
swiss db rollback                           # rollback last migration
swiss db sync                              # sync database schema
```

---

## Setup & Installation

### Swiss Configuration
```bash
swiss config set "key" "value"              # set configuration value
swiss config get "key"                      # get configuration value
swiss config export "file"                  # export configuration
swiss config import "file"                  # import configuration
```

### Swiss Shell Integration
```bash
swiss shell integration bash                # add to bash profile
swiss shell integration zsh                # add to zsh profile
swiss shell integration fish               # add to fish profile
swiss shell completion                      # generate shell completion
```

### Swiss Updates & Maintenance
```bash
swiss update                               # update Swiss to latest version
swiss check                                # check for updates and available features
swiss backup                               # backup Swiss configuration
swiss restore "config-id"                   # restore previous configuration
```

---

## Usage Guidelines

### Command Structure
Each shortcut command follows this pattern:
```
swiss <module> <shortcut> [<arguments>]
```

### Best Practices
1. **Keep it simple** - chain only 3-5 operations at most
2. **Use clear names** - descriptive shortcut names
3. **Add comments** - document complex chains in config
4. **Error handling** - shortcuts should provide clear error messages
5. **Logging** - output what's happening for transparency

### Extension Ideas
- [ ] Configuration-based shortcuts
- [ ] Project-specific shortcut templates
- [ ] Shortcut chaining (run chain A after chain B)
- [ ] Shortcut aliases (create personal shortcuts)
- [ ] Shortcut logging and analytics
- [ ] Interactive shortcut selection

### Implementation Priority
1. Git chaining operations (most common use case)
2. Development workflow shortcuts
3. System monitoring tools
4. Project management utilities

---

*Created for Swiss CLI shortcut module development*
