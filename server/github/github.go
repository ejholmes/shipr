package github

import (
	"net/http"

	"github.com/ejholmes/buble"
	"github.com/ejholmes/go-github/github"
	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
)

// ResponseWriter wraps a buble.ResponseWriter.
type ResponseWriter interface {
	buble.ResponseWriter
}

// Response wraps a buble.Response.
type Response struct {
	buble.ResponseWriter
}

// Request wraps a buble.Request.
type Request struct {
	*buble.Request
}

// HandlerFunc defines the method signature for event handlers.
type HandlerFunc func(Shipr, ResponseWriter, *Request)

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
	shipr  Shipr
	router *mux.Router
}

// New returns a new Handler.
func New(sh Shipr) *GitHub {
	h := &GitHub{shipr: sh, router: mux.NewRouter()}

	// Handlers
	h.Handle("deployment", Deployment)
	h.Handle("deployment_status", DeploymentStatus)

	return h
}

// Handle maps an event to a HandlerFunc.
func (g *GitHub) Handle(event string, fn HandlerFunc) {
	h := &buble.Handler{
		HandlerFunc: func(w buble.ResponseWriter, r *buble.Request) {
			resp := &Response{ResponseWriter: w}
			req := &Request{Request: r}
			fn(g.shipr, resp, req)
		},
	}
	g.router.Methods("POST").Headers(EventHeader, event).Handler(h)
}

func (g *GitHub) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.router.ServeHTTP(w, r)
}

// Deployment handles "deployment" events.
func Deployment(sh Shipr, w ResponseWriter, r *Request) {
	var d github.Deployment
	r.Decode(&d)

	err := sh.Deploy(&description{&d})
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
}

// DeploymentStatus handles "deployment_status" events.
func DeploymentStatus(sh Shipr, w ResponseWriter, r *Request) {
	var d github.DeploymentStatus
	r.Decode(&d)

	err := sh.Notify(newNotification(&d))
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
}
