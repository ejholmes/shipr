package github

import (
	"encoding/json"
	"net/http"

	"github.com/ejholmes/go-github/github"
	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
)

const (
	// EventHeader is the name of the header that determines what type of event this is.
	EventHeader = "X-GitHub-Event"

	// SigHeader is the name of the header that contains the sha1 of the payload.
	SigHeader = "X-Hub-Signature"
)

// Shipr is an interface that declares the methods we use from shipr.
type Shipr interface {
	Deploy(shipr.Description) error
	Notify(shipr.Notification) error
}

// GitHub demuxes incoming webhooks from GitHub and handles them.
type GitHub struct {
	router *mux.Router
}

// New returns a new Handler.
func New(sh Shipr) *GitHub {
	r := mux.NewRouter()
	h := &GitHub{router: r}

	var handlers = map[string]http.Handler{
		"deployment":        &DeploymentHandler{shipr: sh},
		"deployment_status": &DeploymentStatusHandler{shipr: sh},
	}

	for event, handler := range handlers {
		r.Methods("POST").Headers(EventHeader, event).Handler(handler)
	}

	return h
}

func (h *GitHub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

// DeploymentHandler handles "deployment" events.
type DeploymentHandler struct {
	shipr Shipr
}

func (h *DeploymentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var d github.Deployment
	decodeRequest(r, &d)

	err := h.shipr.Deploy(&description{&d})
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
}

// DeploymentStatusHandler handles "deployment_status" events.
type DeploymentStatusHandler struct {
	shipr Shipr
}

func (h *DeploymentStatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var d github.DeploymentStatus
	decodeRequest(r, &d)

	err := h.shipr.Notify(newNotification(&d))
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
}

func decodeRequest(r *http.Request, v interface{}) {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		panic(err)
	}
}
