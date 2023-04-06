package votemachine

type VoteMachineType string

type VoteMachineInterface interface {
	init(data interface{}, isOutput bool, noOfChildren int)
	Start(from VoteMachineType, seed interface{})
	IsValidChoice(who string, userSelectedOptions []interface{}) bool
	Vote(who string, userSelectedOptions []interface{})
	GetChoices() interface{}
	GetTallyResult() (machineType VoteMachineType, nextChildId int, votedResult interface{})
	GetCurrentVoteState() interface{}
}

const (
	VM_SingleChoiceRaceToMax   VoteMachineType = "SingleChoiceRaceToMax"
	VM_MultipleChoiceRaceToMax VoteMachineType = "MultipleChoiceRaceToMax"
)

type RequiredVoteData struct {
	StartAfter int `json:"start_after"` // timestamp; start after CheckPoint start
	EndBefore  int `json:"end_before"`  // length in second; must > 1 min in production
}

func BuildVoteMachine(machineType VoteMachineType, data interface{}, isOutput bool, noOfChildren int) VoteMachineInterface {
	switch machineType {
	case VM_SingleChoiceRaceToMax:
		tmp := &SingleChoiceRaceToMax{}
		tmp.init(data, isOutput, noOfChildren)
		return tmp
	case VM_MultipleChoiceRaceToMax:
		return &MultipleChoiceRaceToMax{}
	}
	return nil
}
