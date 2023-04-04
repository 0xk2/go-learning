package votemachine

type VoteMachineType string

const (
	SingleChoiceRaceToMax   VoteMachineType = "SingleChoiceRaceToMax"
	MultipleChoiceRaceToMax VoteMachineType = "MultipleChoiceRaceToMax"
)

type RequiredVoteData struct {
	StartAfter int `json:"start_after"` // timestamp; start after Node start
	EndBefore  int `json:"end_before"`  // length in second; must > 1 min in production
}

func Vote(nodeType VoteMachineType, who string, userSelectedOptions []int, voted map[int]int, allData interface{}) (int, map[int]int) {
	switch nodeType {
	case SingleChoiceRaceToMax:
		return SingleChoiceRaceToMax_Vote(who, userSelectedOptions, voted, allData)
	case MultipleChoiceRaceToMax:
		return MultipleChoiceRaceToMax_Vote(who, userSelectedOptions, voted, allData)
	}
	return -1, nil
}

func Parse(nodeType VoteMachineType, allData interface{}) interface{} {
	switch nodeType {
	case SingleChoiceRaceToMax:
		return SingleChoiceRaceToMax_Parse(allData)
	case MultipleChoiceRaceToMax:
		return MultipleChoiceRaceToMax_Parse(allData)
	}
	return nil
}

func GetOptions(nodeType VoteMachineType, allData interface{}) []string {
	switch nodeType {
	case SingleChoiceRaceToMax:
		return allData.(SingleChoiceRaceToMaxData).Options
	case MultipleChoiceRaceToMax:
		return allData.(MultipleChoiceRaceToMaxData).Options
	}
	return nil
}
