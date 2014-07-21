package server

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
	"github.com/remind101/shipr/server/api"
	"github.com/remind101/shipr/server/github"
)

const apiContentType = "application/vnd.shipr+json; version=1"

// Server is the http.Handler for serving the application.
type Server struct {
	*shipr.Shipr
	http.Handler
}

// NewServer returns a new Server.
func NewServer(c *shipr.Shipr) *Server {
	m := mux.NewRouter()

	// GitHub webhooks.
	m.Handle("/github", github.New(c, ""))

	// API
	m.Headers("Accept", apiContentType).Handler(api.New(c))

	// Middleware.
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewStatic(http.Dir("server/frontend")))
	n.UseHandler(m)

	return &Server{c, n}
}
