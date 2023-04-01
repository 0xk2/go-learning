package decisiontree

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Vote struct {
	Who       string `json:"who"`
	Options   []int  `json:"options"`
	MissionId string `json:"missionId"`
}

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	SetupCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var vote Vote
	err := decoder.Decode(&vote)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	who := vote.Who
	options := vote.Options
	mission := Missions[vote.MissionId]
	if mission == nil {
		log.Print("Invalid mission id")
		return
	}
	isValid := mission.isValidChoice(options)
	if !isValid {
		fmt.Printf("*%s* %s vote %v, this is an invalid vote\n", mission.current.nodeType, who, options)
		return
	}

	selectedOption, votedResult := mission.current.Vote(who, options)
	var originalOptions []string
	if mission.current.nodeType == "MultipleChoice" {
		_, originalOptions = ConverToMultipleChoiceData(mission.current.data)
	} else if mission.current.nodeType == "SingleChoice" {
		originalOptions = mission.current.data.(SingleChoiceData).Options
	}
	if selectedOption != -1 && votedResult != nil {
		mission.choose(selectedOption)
		mission.current.Start(votedResult, originalOptions)
		log.Print("current: " + mission.current.name)
		if mission.current.nodeType == "MultipleChoice" {
			max, options := ConverToMultipleChoiceData(mission.current.data)
			log.Printf("max: %d, options: %s", max, options)
		}
		if mission.current.isOuput == true {
			fmt.Printf("The mission is done, the result is %s\n; can tweet now", mission.current.name)
			log.Println(votedResult)
			if mission.current.nodeType == "MultipleChoice" {
				tweet("result of " + vote.MissionId + " is: " + fmt.Sprintf("%v", options))
			} else {
				tweet("result of " + vote.MissionId + " is: " + mission.current.name)
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", "{status: \"ok\"}")
}
