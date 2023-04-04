package network

import (
	. "gobyexample/decision_tree/dtree"
	"log"
	"net/http"
)

var Missions = make(map[string]*Tree)

func GoServeTree() {
	http.HandleFunc("/create", CreateHandler)
	http.HandleFunc("/vote", VoteHandler)
	http.HandleFunc("/show", ShowHandler)
	log.Println("Server start at 8080")

	http.ListenAndServe(":8080", nil)
}
