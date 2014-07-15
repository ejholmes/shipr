package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
)

// API serves http requests for the API.
type API struct {
	*shipr.Shipr
	router *mux.Router
}

// New returns a new instance of API, with all of the routes added.
func New(c *shipr.Shipr) http.Handler {
	a := &API{c, mux.NewRouter()}

	// Routes
	a.Handle("/jobs", JobsList)
	a.Handle("/jobs/{id}", JobsInfo)

	return a
}

// Handle takes a path and a Handler func to handle requests to path.
func (a *API) Handle(path string, hd Handler) {
	h := &handler{a.Shipr, hd}
	a.router.Handle(path, h)
}

// ServeHTTP implements the http.Handler interface.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

// Handler is a function signature that can handle a request and return a status code,
// and a response object.
type Handler func(*shipr.Shipr, ResponseWriter, *Request)

// Request wraps http.Request.
type Request struct {
	*http.Request
	Vars map[string]string
}

// ResponseWriter wraps an http.ResponseWriter.
type ResponseWriter interface {
	Status(int)
	Present(int, interface{})
	Error(int, string)
}

type responseWriter struct {
	http.ResponseWriter
}

// Status sets the status code.
func (w *responseWriter) Status(code int) {
	w.WriteHeader(code)
}

// Present presents the interface in JSON format.
func (w *responseWriter) Present(code int, v interface{}) {
	w.Status(code)
	json.NewEncoder(w).Encode(v)
}

// Error takes a string error message and presents it.
func (w *responseWriter) Error(code int, msg string) {
	err := &ErrorResponse{Message: msg}
	w.Present(code, err)
}

// ErrorResponse is the format we respond with when there's an error.
type ErrorResponse struct {
	Message string
}

// Var returns a single URL param.
func (r *Request) Var(v string) string {
	return r.Vars[v]
}

// handler wraps a Handler to return a proper JSON response.
type handler struct {
	*shipr.Shipr
	Handle Handler
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := &responseWriter{w}
	h.Handle(h.Shipr, resp, &Request{Request: r, Vars: mux.Vars(r)})
}
