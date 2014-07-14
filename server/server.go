package server

import (
	"encoding/json"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
)

type Server struct {
	*shipr.Shipr
	http.Handler
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

func decodeRequest(r *http.Request, v interface{}) {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		panic(err)
	}
}
