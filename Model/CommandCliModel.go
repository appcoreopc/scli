package Model

type CommandCliModel struct {
	Name        string
	Description string
	Version     int
	Tools       []ToolModel
}

type ToolModel struct {
	Name        string
	Description string
	Version     int
	Command     string
	Packageurl  string
}
