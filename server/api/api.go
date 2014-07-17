package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
)

// Error represents an http error.
type Error struct {
	error
	Status int
}

// ErrorResponse is the format we respond with when there's an error.
type ErrorResponse struct {
	Error *Error `json:"error"`
}

var (
	// ErrNotFound is an error that represents a 404 error.
	ErrNotFound = &Error{error: errors.New("Not Found"), Status: 404}
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
	a.Handle("GET", "/jobs", JobsList)
	a.Handle("GET", "/jobs/{id}", JobsInfo)

	return a
}

// Handle takes a path and a Handler func to handle requests to path.
func (a *API) Handle(method, path string, hd Handler) {
	h := &handler{a.Shipr, hd, method}
	a.router.Handle(path, h).Methods(method)
}

// ServeHTTP implements the http.Handler interface.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

// Handler is a function signature that can handle a request and return a status code,
// and a response object.
type Handler func(*shipr.Shipr, *Response, *Request)

// Request wraps http.Request.
type Request struct {
	vars map[string]string
}

// Response is an object for building a response.
type Response struct {
	resource interface{}
	status   int
}

// Status sets the status code.
func (w *Response) Status(code int) {
	w.status = code
}

// Present presents the interface in JSON format.
func (w *Response) Present(resource interface{}) {
	w.resource = resource
}

// Error takes a string error message and presents it.
func (w *Response) Error(err *Error) {
	res := &ErrorResponse{Error: err}
	w.Status(err.Status)
	w.Present(res)
}

// NotFound returns a standard 404 Not Found response.
func (w *Response) NotFound() {
	w.Error(ErrNotFound)
}

// Var returns a single URL param.
func (r *Request) Var(v string) string {
	return r.vars[v]
}

// handler wraps a Handler to return a proper JSON response.
type handler struct {
	*shipr.Shipr
	handle Handler
	method string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := &Response{}
	req := &Request{vars: mux.Vars(r)}
	h.handle(h.Shipr, res, req)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.status)
	json.NewEncoder(w).Encode(res.resource)
}
