package Model

type CommandCliModel struct {
	Name        string
	Description string
	Tools       []ToolModel
}

type ToolModel struct {
	Name        string
	Description string
	Version     string
	Command     string
}
