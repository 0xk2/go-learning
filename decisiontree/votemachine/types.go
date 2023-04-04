package votemachine

type VoteMachineType string

const (
	SingleChoice   VoteMachineType = "SingleChoiceRaceToMax"
	MultipleChoice VoteMachineType = "MultipleChoiceRaceToMax"
)

type RequiredVoteData struct {
	StartAfter int `json:"start_after"` // timestamp; start after Node start
	EndBefore  int `json:"end_before"`  // length in second; must > 1 min in production
}

type MultipleChoiceData struct {
	Options []string `json:"options"`
	Max     int      `json:"max"`
}
type SingleChoiceData struct {
	Options []string `json:"options"`
}
