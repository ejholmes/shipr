package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
)

type Server struct {
	*shipr.Shipr
	handler http.Handler
}

func NewServer(c *shipr.Shipr) *Server {
	m := mux.NewRouter()

	// GitHub webhooks.
	m.Handle("/github", NewGitHubHandler(c))

	// Middleware.
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.UseHandler(m)

	return &Server{c, n}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}
