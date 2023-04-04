package votemachine

import (
	"fmt"
	"gobyexample/decision_tree/utils"
)

type SingleChoiceRaceToMaxData struct {
	Options []string `json:"options"`
	Max     int      `json:"max"`
}

func SingleChoiceRaceToMax_Parse(data interface{}) interface{} {
	tmp := data.(map[string]interface{})
	options := make([]string, 0)
	max := 0
	for key, value := range tmp {
		if key == "options" {
			for _, opt := range value.([]interface{}) {
				options = append(options, opt.(string))
			}
		}
		if key == "max" {
			max = utils.InterfaceToInt(value)
		}
	}
	return SingleChoiceRaceToMaxData{
		Options: options,
		Max:     max,
	}
}

func SingleChoiceRaceToMax_Vote(who string, userSelectedOptions []int, voted map[int]int, allData interface{}) (int, map[int]int) {
	result := -1
	votedResult := make(map[int]int)

	singleChoice := userSelectedOptions[0]
	voted[singleChoice] += 1
	optionName := allData.(SingleChoiceRaceToMaxData).Options[singleChoice]
	fmt.Printf("%s vote %d [%s]. There are %d person(s) choose %s\n", who, userSelectedOptions[0], optionName, voted[singleChoice], optionName)
	if voted[singleChoice] == 3 { // let's move on
		result = singleChoice
		votedResult = voted
	}
	return result, votedResult
}
