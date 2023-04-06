package types

import "gobyexample/decision_tree/votemachine"

type Vote struct {
	Who       string        `json:"who"`
	Options   []interface{} `json:"options"`
	MissionId string        `json:"missionId"`
}

type VoteResponse struct {
	Status      bool                        `json:"status"`
	Option      []interface{}               `json:"option"`
	Message     string                      `json:"message"`
	MachineType votemachine.VoteMachineType `json:"machine_type"`
}
