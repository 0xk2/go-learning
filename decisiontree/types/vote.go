package types

type Vote struct {
	Who       string `json:"who"`
	Options   []int  `json:"options"`
	MissionId string `json:"missionId"`
}

type VoteResponse struct {
	Status  bool     `json:"status"`
	Option  []string `json:"option"`
	Message string   `json:"message"`
}
