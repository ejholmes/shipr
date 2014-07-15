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
	*http.Request
	Vars map[string]string
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
func (w *Response) Error(code int, msg string) {
	res := &ErrorResponse{Error: msg}
	w.Status(code)
	w.Present(res)
}

// NotFound returns a standard 404 Not Found response.
func (w *Response) NotFound() {
	w.Error(404, "Not Found")
}

// ErrorResponse is the format we respond with when there's an error.
type ErrorResponse struct {
	Error string `json:"error"`
}

// Var returns a single URL param.
func (r *Request) Var(v string) string {
	return r.Vars[v]
}

// handler wraps a Handler to return a proper JSON response.
type handler struct {
	*shipr.Shipr
	Handle Handler
	Method string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := &Response{}
	req := &Request{Request: r, Vars: mux.Vars(r)}
	h.Handle(h.Shipr, res, req)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.status)
	json.NewEncoder(w).Encode(res.resource)
}
