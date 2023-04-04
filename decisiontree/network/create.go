package network

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gobyexample/decisiontree/dtree"
	. "gobyexample/decisiontree/types"
	"gobyexample/decisiontree/utils"
	"gobyexample/decisiontree/votemachine"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
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

	var missionData MissionData
	err := decoder.Decode(&missionData)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	startId := missionData.Start
	var nodes = make(map[string]*dtree.Node)
	for _, nodeData := range missionData.Nodes {
		if nodeData.NodeType == "MultipleChoice" {
			max, options := utils.ConverToMultipleChoiceData(nodeData.Data)
			nodeData.Data = votemachine.MultipleChoiceData{
				Options: options,
				Max:     max,
			}
		} else if nodeData.NodeType == "SingleChoice" {
			options := utils.ConvertToSingleChoiceData(nodeData.Data)
			nodeData.Data = votemachine.SingleChoiceData{
				Options: options,
			}
		}
		nodes[nodeData.Id] = dtree.CreateEmptyNode(nodeData.Name, nodeData.IsOuput, nodeData.NodeType, nodeData.Data)
		nodes[nodeData.Id].Id = nodeData.Id
	}
	for _, nodeData := range missionData.Nodes {
		if nodeData.Parent != "" {
			nodes[nodeData.Parent].Attach(nodes[nodeData.Id])
		}
	}
	missionId := utils.RandString(16)
	Missions[missionId] = dtree.CreateTree(nodes[startId], missionData.Name, missionData.Description)
	Missions[missionId].PrintFromCurrent()

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
