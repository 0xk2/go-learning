package types

type HistoryData struct {
	Name  string      `json:"name"`
	Voted map[int]int `json:"voted"`
	// add originalOptions
}

type ShowResponse struct {
	MissionId          string        `json:"mission_id"`
	MissionName        string        `json:"mission_name"`
	MissionDescription string        `json:"mission_description"`
	Name               string        `json:"current"`
	Vote               map[int]int   `json:"vote"`
	Choice             interface{}   `json:"choice"`
	NodeType           string        `json:"node_type"`
	AllHistoryData     []HistoryData `json:"all_history_data"`
}
