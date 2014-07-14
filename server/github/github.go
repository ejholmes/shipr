package github

import (
	"encoding/json"
	"net/http"

	"github.com/ejholmes/go-github/github"
	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
)

const EventHeader = "X-GitHub-Event"

type EventHandler http.Handler

// Handler demuxes incoming webhooks from GitHub and handles them.
type Handler struct {
	http.Handler
}

func New(c *shipr.Shipr) http.Handler {
	m := mux.NewRouter()
	h := &Handler{m}

	var handlers = map[string]EventHandler{
		"deployment":        &DeploymentHandler{c},
		"deployment_status": &DeploymentStatusHandler{c},
	}

	for event, handler := range handlers {
		m.Methods("POST").Headers(EventHeader, event).Handler(handler)
	}

	return h
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

func decodeRequest(r *http.Request, v interface{}) {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		panic(err)
	}
}
