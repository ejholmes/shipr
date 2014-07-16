package github

import (
	"encoding/json"
	"net/http"

	"github.com/ejholmes/go-github/github"
	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
)

// EventHeader is the name of the header that determines what type of event this is.
const EventHeader = "X-GitHub-Event"

// Handler demuxes incoming webhooks from GitHub and handles them.
type Handler struct {
	http.Handler
}

// New returns a new Handler.
func New(c *shipr.Shipr) http.Handler {
	m := mux.NewRouter()
	h := &Handler{m}

	var handlers = map[string]http.Handler{
		"deployment":        &DeploymentHandler{c},
		"deployment_status": &DeploymentStatusHandler{c},
	}

	for event, handler := range handlers {
		m.Methods("POST").Headers(EventHeader, event).Handler(handler)
	}

	return h
}

// DeploymentHandler handles "deployment" events.
type DeploymentHandler struct {
	*shipr.Shipr
}

func (h *DeploymentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var d github.Deployment
	decodeRequest(r, &d)

	h.Deploy(&description{&d})
}

// DeploymentStatusHandler handles "deployment_status" events.
type DeploymentStatusHandler struct {
	*shipr.Shipr
}

func (h *DeploymentStatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var d github.DeploymentStatus
	decodeRequest(r, &d)

	h.Notify(newNotification(&d))
}

func decodeRequest(r *http.Request, v interface{}) {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		panic(err)
	}
}
