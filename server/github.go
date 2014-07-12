package main

import (
	"net/http"

	"github.com/ejholmes/go-github/github"
	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
	"github.com/remind101/shipr/util"
)

const GitHubEventHeader = "X-GitHub-Event"

type GitHubEventHandler http.Handler

// GitHubHandler demuxes incoming webhooks from GitHub and handles them.
type GitHubHandler struct {
	http.Handler
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

// description wraps a github.Deployment to implement the shipr.Description interface.
type description struct {
	*github.Deployment
}

func (d *description) Guid() int {
	return *d.Deployment.ID
}

func (d *description) RepoName() shipr.RepoName {
	return shipr.RepoName(util.SafeString(d.Deployment.Repository.FullName))
}

func (d *description) Sha() string {
	return util.SafeString(d.Deployment.Sha)
}

func (d *description) Ref() string {
	return util.SafeString(d.Deployment.Ref)
}

func (d *description) Environment() string {
	return util.SafeString(d.Deployment.Environment)
}

func (d *description) Description() string {
	return util.SafeString(d.Deployment.Description)
}

type DeploymentHandler struct {
	*shipr.Shipr
}

func (h *DeploymentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var d github.Deployment
	decodeRequest(r, &d)

	h.Deploy(&description{&d})
}

type DeploymentStatusHandler struct {
	*shipr.Shipr
}

func (h *DeploymentStatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
