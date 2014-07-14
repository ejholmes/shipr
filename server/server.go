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

type Server struct {
	*shipr.Shipr
	http.Handler
}

func NewServer(c *shipr.Shipr) *Server {
	m := mux.NewRouter()

	// Health checking.
	m.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok\n"))
	}))

	// GitHub webhooks.
	m.Handle("/github", github.New(c))

	// API
	m.Headers("Accept", apiContentType).Handler(api.New(c))

	// Middleware.
	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.UseHandler(m)

	return &Server{c, n}
}
