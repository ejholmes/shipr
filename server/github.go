package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
)

const GitHubEventHeader = "X-GitHub-Event"

type GitHubEventHandler http.Handler

// GitHubHandler demuxes incoming webhooks from GitHub and handles them.
type GitHubHandler struct {
	handler http.Handler
}

func NewGitHubHandler(c *shipr.Shipr) *GitHubHandler {
	m := mux.NewRouter()
	h := &GitHubHandler{m}

	var handlers = map[string]GitHubEventHandler{
		"deployment":        &DeploymentHandler{c},
		"deployment_status": &DeploymentStatusHandler{c},
	}

	for event, handler := range handlers {
		m.Methods("POST").Headers(GitHubEventHeader, event).Handler(handler)
	}

	return h
}

func (h *GitHubHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	h.handler.ServeHTTP(w, r)
}

type DeploymentHandler struct {
	*shipr.Shipr
}

func (h *DeploymentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deployment")
}

type DeploymentStatusHandler struct {
	*shipr.Shipr
}

func (h *DeploymentStatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeploymentStatus")
}
