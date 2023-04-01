package decisiontree

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateResponse struct {
	Id string `json:"id"`
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
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

	var missionData MissionData
	err := decoder.Decode(&missionData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	startId := missionData.Start
	var nodes = make(map[string]*Node)
	for _, nodeData := range missionData.Nodes {
		if nodeData.NodeType == "MultipleChoice" {
			max, options := ConverToMultipleChoiceData(nodeData.Data)
			nodeData.Data = MultipleChoiceData{
				Options: options,
				Max:     max,
			}
		} else if nodeData.NodeType == "SingleChoice" {
			options := ConvertToSingleChoiceData(nodeData.Data)
			nodeData.Data = SingleChoiceData{
				Options: options,
			}
		}
		nodes[nodeData.Id] = createEmptyNode(nodeData.Name, nodeData.IsOuput, nodeData.NodeType, nodeData.Data)
		nodes[nodeData.Id].id = nodeData.Id
	}
	for _, nodeData := range missionData.Nodes {
		if nodeData.Parent != "" {
			nodes[nodeData.Parent].attach(nodes[nodeData.Id])
		}
	}
	missionId := RandString(16)
	Missions[missionId] = createTree(nodes[startId], missionData.Name, missionData.Description)
	Missions[missionId].printFromCurrent()

	CreateResponse := CreateResponse{
		Id: missionId,
	}
	jsonData, err := json.Marshal(CreateResponse)
	if err != nil {
		http.Error(w, "Error crafting response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", jsonData)
}
