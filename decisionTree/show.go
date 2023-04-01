package decisiontree

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

func ShowHandler(w http.ResponseWriter, r *http.Request) {
	SetupCORS(&w, r)
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
	parent := mission.current.parent
	for parent != nil {
		allHistoryData = append(allHistoryData, HistoryData{
			Name:  parent.name,
			Voted: parent.voted,
		})
		parent = parent.parent
	}
	// reverse allHistoryData
	for i, j := 0, len(allHistoryData)-1; i < j; i, j = i+1, j-1 {
		allHistoryData[i], allHistoryData[j] = allHistoryData[j], allHistoryData[i]
	}

	resp := ShowResponse{
		MissionId:          missionId,
		MissionName:        mission.name,
		MissionDescription: mission.description,
		Name:               mission.current.name,
		Vote:               mission.current.voted,
		NodeType:           mission.current.nodeType,
		AllHistoryData:     allHistoryData,
	}
	if mission.current.nodeType == "MultipleChoiceData" {
		resp.Choice = mission.current.data.(MultipleChoiceData).Options
	} else if mission.current.nodeType == "SingleChoice" {
		resp.Choice = mission.current.data.(SingleChoiceData).Options
	}
	mission.current.print()
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
