package network

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gobyexample/decision_tree/enforcer"
	. "gobyexample/decision_tree/types"
	"gobyexample/decision_tree/utils"
	"gobyexample/decision_tree/votemachine"
)

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	utils.SetupCORS(&w, r)
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := VoteResponse{
			Status:  false,
			Message: "Invalid mission id",
		}
		// golang json encode
		jsonData, _ := json.Marshal(resp)
		fmt.Fprintf(w, "%s", jsonData)
		log.Print("Invalid mission id")
		return
	}
	resp := VoteResponse{
		Option: mission.Current.ConvertOptionIdxToString(options),
	}

	isValid := mission.IsValidChoice(options)
	if !isValid {
		resp.Status = false
		jsonData, _ := json.Marshal(resp)
		fmt.Printf("*%s* %s vote %v, this is an invalid vote\n", mission.Current.NodeType, who, options)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", jsonData)
		return
	}
	resp.Status = true
	current := mission.Current
	selectedOption, votedResult := current.Vote(who, options)
	originalOptions := votemachine.GetOptions(current.NodeType, current.AllData)
	if selectedOption != -1 && votedResult != nil {
		mission.Choose(selectedOption)
		current = mission.Current
		current.Start(votedResult, originalOptions)
		log.Print("current: " + current.Name)
		if mission.Current.IsOuput == true {
			fmt.Printf("The mission is done, the result is %s\n; can tweet now", mission.Current.Name)
			log.Println(votedResult)
			if mission.Current.NodeType == "MultipleChoice" {
				enforcer.Tweet("result of " + vote.MissionId + " is: " + fmt.Sprintf("%v", options))
			} else {
				enforcer.Tweet("result of " + vote.MissionId + " is: " + mission.Current.Name)
			}
		}
	}
	jsonData, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", jsonData)
}
