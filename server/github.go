package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
)

const GitHubEventHeader = "X-GitHub-Event"

type GitHubEventHandler func(*shipr.Shipr, http.ResponseWriter, *http.Request)

var GitHubEventHandlers = map[string]GitHubEventHandler{
	"deployment":        HandleDeployment,
	"deployment_status": HandleDeploymentStatus,
}

// GitHubHandler demuxes incoming webhooks from GitHub and handles them.
type GitHubHandler struct {
	handler http.Handler
}

func NewGitHubHandler() *GitHubHandler {
	m := mux.NewRouter()
	h := &GitHubHandler{m}

	for event, f := range GitHubEventHandlers {
		handler := func(w http.ResponseWriter, r *http.Request) { f(nil, w, r) }
		m.HandleFunc("/github", handler).Methods("POST").Headers(GitHubEventHeader, event)
	}

	return h
}

func (h *GitHubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}

func HandleDeployment(c *shipr.Shipr, w http.ResponseWriter, r *http.Request) {
}

func HandleDeploymentStatus(c *shipr.Shipr, w http.ResponseWriter, r *http.Request) {
}
