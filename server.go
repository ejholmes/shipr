package main

import (
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
)

const GitHubEventHeader = "X-GitHub-Event"

var GitHubEventHandlers = map[string]http.HandlerFunc{
	"deployment":        HandleDeployment,
	"deployment_status": HandleDeploymentStatus,
}

func NewServer() http.Handler {
	m := mux.NewRouter()

	for event, handler := range GitHubEventHandlers {
		m.HandleFunc("/github", handler).Methods("POST").Headers(GitHubEventHeader, event)
	}

	return m
}

func HandleDeployment(w http.ResponseWriter, r *http.Request) {
	var d GitHubDeployment
	decodeRequest(r, &d)

	err := Deploy(&d)
	if err != nil {
		panic(err)
	}
}

func HandleDeploymentStatus(w http.ResponseWriter, r *http.Request) {
	var s GitHubDeploymentStatus
	decodeRequest(r, &s)
}

func decodeRequest(r *http.Request, v interface{}) {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		panic(err)
	}
}
