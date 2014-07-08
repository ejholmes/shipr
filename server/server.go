package main

import (
	"net/http"

	"github.com/remind101/shipr"
)

type Server struct {
	*shipr.Shipr
}

func NewServer(s *shipr.Shipr) *Server {
	return &Server{s}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
