package types

type MissionData struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Start       string     `json:"start"`
	Nodes       []NodeData `json:"nodes"`
}
type NodeData struct {
	Id       string      `json:"id"`
	Name     string      `json:"name"`
	Parent   string      `json:"parent"`
	IsOuput  bool        `json:"is_output"`
	NodeType string      `json:"node_type"`
	Data     interface{} `json:"data"`
}
