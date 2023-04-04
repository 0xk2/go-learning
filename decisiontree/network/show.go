package network

import (
	"encoding/json"
	"fmt"
	"gobyexample/decisiontree/utils"
	"gobyexample/decisiontree/votemachine"
	"log"
	"net/http"

	. "gobyexample/decisiontree/types"
)

func ShowHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	missionId := r.URL.Query().Get("id")

	mission := Missions[missionId]
	if mission == nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	var allHistoryData []HistoryData
	parent := mission.Current.Parent
	for parent != nil {
		allHistoryData = append(allHistoryData, HistoryData{
			Name:  parent.Name,
			Voted: parent.Voted,
		})
		parent = parent.Parent
	}
	// reverse allHistoryData
	for i, j := 0, len(allHistoryData)-1; i < j; i, j = i+1, j-1 {
		allHistoryData[i], allHistoryData[j] = allHistoryData[j], allHistoryData[i]
	}

	resp := ShowResponse{
		MissionId:          missionId,
		MissionName:        mission.Name,
		MissionDescription: mission.Description,
		Name:               mission.Current.Name,
		Vote:               mission.Current.Voted,
		NodeType:           mission.Current.NodeType,
		AllHistoryData:     allHistoryData,
	}
	if mission.Current.NodeType == "MultipleChoiceData" {
		resp.Choice = mission.Current.AllData.(votemachine.MultipleChoiceData).Options
	} else if mission.Current.NodeType == "SingleChoice" {
		resp.Choice = mission.Current.AllData.(votemachine.SingleChoiceData).Options
	}
	mission.Current.Print()
	log.Print(resp)
	jsonData, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Error crafting response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", jsonData)
}
