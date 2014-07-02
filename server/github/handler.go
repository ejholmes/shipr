package github

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
)

const EventHeader = "X-GitHub-Event"

var EventHandlers = map[string]http.HandlerFunc{
	"deployment":        HandleDeployment,
	"deployment_status": HandleDeploymentStatus,
}

// Handler implements the http.Handler interface and is capable of handling
// github webhooks.
type Handler struct {
	handler http.Handler
}

// NewHandler returns a new Handler instance.
func NewHandler() *Handler {
	m := mux.NewRouter()

	for event, handler := range EventHandlers {
		m.HandleFunc("/", handler).Methods("POST").Headers(EventHeader, event)
	}

	return &Handler{m}
}

// ServeHTTP implements the http.Handler interface.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}

func HandleDeployment(w http.ResponseWriter, r *http.Request) {
	var d Deployment
	decodeRequest(r, &d)

	err := shipr.Deploy(&d)
	if err != nil {
		panic(err)
	}
}

func HandleDeploymentStatus(w http.ResponseWriter, r *http.Request) {
	var s DeploymentStatus
	decodeRequest(r, &s)
}

func decodeRequest(r *http.Request, v interface{}) {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		panic(err)
	}
}
