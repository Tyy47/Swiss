package initialize


// Language is the key guardrail for the initialize map
type Language string


// project is the structure that holds all the needed information
// to initialize a project using swiss.
type project struct {

	// Build tool
	Tool       string   

	// Arguments needed to init project
	Arguments  []string 	

	// Additional folders needed for project
	Folders    []string 

	// Additional files needed for project
	Files      []string 

	// Toggle if a project needs a manual init like C.
	// C doesn't have a traditional init tool like typescript/bun.
	ManualInit bool     

	// NeedsProjectName stores the state if a project needs a name for init
	NeedsProjectName bool
}
