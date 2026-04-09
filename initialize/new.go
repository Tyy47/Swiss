package initialize


const initProjectList = `
Rust - Cargo
Go - Go
C - Swiss
HTML - Swiss
Zig - Zig
Vanilla TS Web App - Bun/Vite
Svelte Web App - Bun/Vite
React Web App - Bun/Vite`

type InitProject struct {
	Language string // Language name of the project
	BuildTool string // Build tool of the project
	InitTimes int // How many times the project has been init'd during the init process
	Flags []string
	AdditionalFlags []string
	Handler func()
	ManualInit bool
	ManualFiles []string
	ManualFolders []string
}

type ProjectRegistry struct {
	Projects []InitProject
}


func PrintInitProjectList() {
	fmt.Println(initProjectList)
}
