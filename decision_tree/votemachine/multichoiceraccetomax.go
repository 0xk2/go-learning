package votemachine

import (
	"fmt"
	"gobyexample/decision_tree/utils"
)

type MultipleChoiceRaceToMaxData struct {
	Options []string `json:"options"`
	Max     int      `json:"max"`
}

func MultipleChoiceRaceToMax_Parse(data interface{}) interface{} {
	tmp := data.(map[string]interface{})
	var max int
	options := make([]string, 0)
	for key, value := range tmp {
		if key == "options" {
			for _, opt := range value.([]interface{}) {
				options = append(options, opt.(string))
			}
		} else if key == "max" {
			max = utils.InterfaceToInt(value)
		}
	}
	return MultipleChoiceRaceToMaxData{
		Options: options,
		Max:     max,
	}
}

func MultipleChoiceRaceToMax_Vote(who string, userSelectedOptions []int, voted map[int]int, allData interface{}) (int, map[int]int) {
	choices := make([]int, 0)
	for _, opt := range userSelectedOptions {
		choices = append(choices, utils.InterfaceToInt(opt))
	}
	str := ""
	data := allData.(MultipleChoiceRaceToMaxData)
	max := data.Max
	options := data.Options
	for _, choice := range choices {
		str += options[choice] + ","
		voted[choice] += 1
	}
	choosen := make(map[int]int)

	fmt.Printf("%s vote [%s]; top %d choice win will\n", who, str, max)

	for opt, choice := range voted {
		if choice >= max {
			choosen[opt] = choice
		}
	}

	return 0, choosen
}
