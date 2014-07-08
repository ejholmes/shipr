package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/remind101/shipr"
)

func main() {
	var (
		dbdir = flag.String("dbdir", "db", "The db dir containing a dbconf.yml file.")
		env   = flag.String("env", "development", "The environment to run in.")
	)
	c, err := shipr.New(*dbdir, *env)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	log.Fatal(http.ListenAndServe(":3001", NewServer(c)))
}
