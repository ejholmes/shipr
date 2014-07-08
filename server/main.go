package main

import (
	"log"
	"net/http"

	"github.com/remind101/shipr"
)

func main() {
	s := shipr.New()
	log.Fatal(http.ListenAndServe(":3001", NewServer(s)))
}
