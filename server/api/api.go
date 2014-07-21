package api

import (
	"errors"
	"net/http"

	"github.com/ejholmes/buble"
	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
)

var (
	// ErrNotFound is an error that represents a 404 error.
	ErrNotFound = &Error{error: errors.New("Not Found"), Status: 404}

	// ErrBadRequest is an error that represents a 400 error.
	ErrBadRequest = &Error{error: errors.New("Bad Request"), Status: 400}
)

// API implements the http.ServeHTTP interface for serving up the api.
type API struct {
	shipr  *shipr.Shipr
	router *mux.Router
}

// New returns a new API.
func New(c *shipr.Shipr) *API {
	a := &API{shipr: c, router: mux.NewRouter()}

	// Routes
	a.Handle("GET", "/jobs", JobsList)
	a.Handle("GET", "/jobs/{id}", JobsInfo)

	return a
}

// Handle takes a method, path and a HandlerFunc and adds a route to handle
// requests.
func (a *API) Handle(method, path string, fn HandlerFunc) {
	h := &buble.Handler{
		HandlerFunc: func(w buble.ResponseWriter, r *buble.Request) {
			resp := &Response{ResponseWriter: w}
			req := &Request{Request: r, vars: mux.Vars(r.Request)}
			fn(a.shipr, resp, req)
		},
	}
	a.router.Handle(path, h).Methods(method)
}

// ServeHTTP implements the http.ServeHTTP interface.
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.router.ServeHTTP(w, r)
}

// Error represents our error interface.
type Error struct {
	error
	Status int
}

// ErrorResponse represents our api error resource.
type ErrorResponse struct {
	Error *Error `json:"error"`
}

// ResponseWriter wraps an http.ResponseWriter with convenience methods.
type ResponseWriter interface {
	buble.ResponseWriter

	Error(*Error)
	NotFound()
	BadRequest()
}

// Response wraps a buble.Response.
type Response struct {
	buble.ResponseWriter
}

// Error sets the status code associated with this error then encodes the error resource
// into the response.
func (r *Response) Error(err *Error) {
	r.WriteHeader(err.Status)
	r.Encode(&ErrorResponse{Error: err})
}

// NotFound responds with an ErrNotFound response.
func (r *Response) NotFound() {
	r.Error(ErrNotFound)
}

// BadRequest resposne with an ErrBadRequest response.
func (r *Response) BadRequest() {
	r.Error(ErrBadRequest)
}

// Request wraps a buble.Request.
type Request struct {
	*buble.Request
	vars map[string]string
}

// Var returns a mux variable from the url.
func (r *Request) Var(key string) string {
	return r.vars[key]
}

// HandlerFunc is the method signature for api handlers.
type HandlerFunc func(*shipr.Shipr, ResponseWriter, *Request)
