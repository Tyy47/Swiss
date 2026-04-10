# Solution for Web Project Initialization Issue

## Problem Analysis

In `initialize/initialize.go`, the package `init()` function (lines 262-272) automatically registers all projects when the package is imported. This means:

1. When `initialize` package is imported in `main.go`
2. The `init()` function runs
3. `createWebProject()` is called and immediately prompts for user input via `utils.GetUserInput()`
4. This happens every time the program starts, even when not running `swiss init web`

## Current Flow

```
Program Start → Import initialize package → Run init() package function → createWebProject() called → User prompt appears even when not needed
```

## Proposed Solution

### Strategy: Defer Web Project Registration

Instead of auto-registering web projects at package initialization time, we'll:

1. Store web project creation separately (not in registry.projects)
2. Add conditional registration logic in the `init()` function
3. Only register web projects when specifically requested via command-line
4. Keep other projects (Go, Rust, C, HTML, Zig) auto-registered as before

### Implementation Steps

#### Step 1: Modify `initialize/initialize.go`

**Change the line that registers web project:**

Current (line 269):
```go
createWebProject(),
```

Change to:
```go
// Create web project creation function but don't register yet
func getWebProject() project {
	projectName := utils.GetUserInput("Enter your project name: ", "my-app")
	return project{
		Language:   "web",
		Tool:       "bun",
		Arguments:  []string{"create", "vite", projectName, "--template"},
		ManualInit: false,
	}
}
```

#### Step 2: Modify the `init()` function in `initialize/initialize.go`

Update lines 262-272 to conditionally register web project:

```go
func init() {
	// Auto-register standard projects
	projectArray := []project{
		createGoProject(),
		createRustProject(),
		createCProject(),
		createHTMLProject(),
		createZigProject(),
		// Don't include createWebProject() here
	}

	registerProjects(projectArray...)

	// Register web project only when needed
	// This will be called from a specific handler function
	// See Step 3 below for details
}
```

#### Step 3: Add Web-Specific Handler

Create a new function that's used for web command handlers:

```go
func webProjectInit() {
	project := getWebProject()
	
	if err := project.initialize(); err != nil {
		log.Fatal(err)
		return
	}
	
	flagHandler()
	utils.Success("Web project has been created.")
}
```

#### Step 4: Update `main.go` Command Registry

Modify the `initCommand()` function (lines 153-179) to handle web commands separately:

Current `initCommand()` has:
```go
{
	Name: "init",
	HelpMenu: utils.InitHelp,
	Subcommands: []Subcommand{
		{
			Name: "init",
			Flags: map[string]func(){
				"-h":     utils.InitHelp,
				"--help": utils.InitHelp,
				"-l":     initialize.PrintInitProjectList,
				"--list": initialize.PrintInitProjectList,
				"go":     initialize.HandleInput,
				"rust":     initialize.HandleInput,
				"c":     initialize.HandleInput,
				"html":     initialize.HandleInput,
				"web":     initialize.HandleInput,  // This triggers the issue
				"svelte":     initialize.HandleInput,
				"react":     initialize.HandleInput,
			},
		},
	},
}
```

Replace `initialize.HandleInput` for web, svelte, and react with the new handlers:

```go
{
	Name: "init",
	HelpMenu: utils.InitHelp,
	Subcommands: []Subcommand{
		{
			Name: "init",
			Flags: map[string]func(){
				"-h":     utils.InitHelp,
				"--help": utils.InitHelp,
				"-l":     initialize.PrintInitProjectList,
				"--list": initialize.PrintInitProjectList,
				"go":     initialize.HandleInput,
				"rust":     initialize.HandleInput,
				"c":     initialize.HandleInput,
				"html":     initialize.HandleInput,
				"zig":     initialize.HandleInput,
				"web":     initialize.webProjectInit,      // New function
				"svelte":   initialize.webProjectInit,      // New function
				"react":    initialize.webProjectInit,      // New function
			},
		},
	},
}
```

**Note:** You'll need to export the function (change `webProjectInit` to `WebProjectInit`) or use a different approach if there are encapsulation concerns.

### Alternative Approach: Export Web Project Function

Since you may want to keep `createWebProject` in the public API, export it from `initialize/initialize.go`:

```go
// Update createWebProject function signature
func GetWebProject() project {
	projectName := utils.GetUserInput("Enter your project name: ", "my-app")
	return project{
		Language:   "web",
		Tool:       "bun",
		Arguments:  []string{"create", "vite", projectName, "--template"},
		ManualInit: false,
	}
}
```

Then in `initialize/initialize.go`:
```go
func init() {
	projectArray := []project{
		createGoProject(),
		createRustProject(),
		createCProject(),
		createHTMLProject(),
		createZigProject(),
		// Don't register web project here
	}

	registerProjects(projectArray...)

	// Web project will be registered later when command is executed
}

func HandleWebProject() {
	project := GetWebProject()
	if err := project.initialize(); err != nil {
		log.Fatal(err)
		return
	}
	
	flagHandler()
	utils.Success("Web project has been created.")
}
```

## How It Works After Changes

**Current flow causes issue:**
- `swiss` (without init) → Prompt appears ❌

**New flow prevents issue:**
- `swiss` (without init) → No prompt ✅
- `swiss init go` → Creates Go project ✅
- `swiss init web` → Creates web project ✅
- `swiss init react` → Creates React project ✅

## Benefits

1. **Lazy Registration**: Web project creation now happens only when explicitly requested
2. **Clearer Separation**: Standard projects are auto-registered; web projects are explicitly triggered
3. **User Experience**: No prompts appear when running other commands or no commands
4. **Backward Compatibility**: Other commands (go, rust, etc.) continue working exactly as before

## Testing

After implementing, test these scenarios:

1. Run `swiss` → Should show help, no prompts
2. Run `swiss init` → Show list, no prompts
3. Run `swiss init web` → Prompt for project name, create project
4. Run `swiss init react` → Prompt for project name, create React project
5. Run `swiss init go rust c html zig` → Create respective projects without prompts

## Files That Need Changes

1. `initialize/initialize.go`:
   - Lines 262-272: Modify `init()` function
   - Lines 214-238: Refactor `createWebProject()` either by exporting or splitting
   - Add new web-specific handler function

2. `main.go`:
   - Lines 153-179: Modify `initCommand()` to use new web handler

## Alternative Minimal Changes

If you prefer minimal changes, here's a simpler approach:

1. Remove `createWebProject()` from the `init()` function in `initialize.go`
2. Don't register web projects in `registry.projects`
3. Create a dedicated handler function `handleWeb()` that:
   - Calls `createWebProject()` before registration
   - Registers it temporarily for the initialization process
   - Doesn't add it back to the registry after

This would keep the existing `createWebProject()` function intact while preventing it from being called during package initialization.