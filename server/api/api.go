package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/remind101/shipr"
)

type API struct {
	*shipr.Shipr
	http.Handler
}

func New(c *shipr.Shipr) http.Handler {
	m := mux.NewRouter()
	api := &API{c, m}

	// Routes.
	//m.HandleFunc("/jobs/{id}", api.JobInfo)

	return api
}
