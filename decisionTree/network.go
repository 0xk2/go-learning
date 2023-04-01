package decisiontree

import (
	"log"
	"net/http"
)

type MultipleChoiceData struct {
	Options []string `json:"options"`
	Max     int      `json:"max"`
}
type SingleChoiceData struct {
	Options []string `json:"options"`
}
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

var Missions = make(map[string]*Tree)

func GoServeTree() {
	http.HandleFunc("/create", CreateHandler)
	http.HandleFunc("/vote", VoteHandler)
	http.HandleFunc("/show", ShowHandler)
	log.Println("Server start at 8080")

	http.ListenAndServe(":8080", nil)
}
