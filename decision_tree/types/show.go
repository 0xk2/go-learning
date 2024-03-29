package types

import "gobyexample/decision_tree/votemachine"

type HistoryData struct {
	Name        string                      `json:"name"`
	Voted       interface{}                 `json:"voted"`
	MachineType votemachine.VoteMachineType `json:"machine_type"`
	// add originalOptions
}

type ShowResponse struct {
	MissionId          string                      `json:"mission_id"`
	MissionName        string                      `json:"mission_name"`
	MissionDescription string                      `json:"mission_description"`
	Name               string                      `json:"current"`
	CurrentVoteResult  interface{}                 `json:"current_vote_result"`
	Choices            interface{}                 `json:"choices"`
	VoteMachineType    votemachine.VoteMachineType `json:"vote_machine_type"`
	AllHistoryData     []HistoryData               `json:"all_history_data"`
}
