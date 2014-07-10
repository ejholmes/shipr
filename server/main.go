package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/remind101/shipr"
)

func main() {
	var (
		dbdir       = flag.String("dbdir", "db", "The db dir containing a dbconf.yml file.")
		env         = flag.String("env", Env("GOENV", "development"), "The environment to run in.")
		port        = flag.String("port", Env("PORT", "3001"), "Port to run the server on.")
		githubToken = flag.String("github-token", Env("SHIPR_GITHUB_TOKEN", ""), "GitHub api token.")
		herokuToken = flag.String("heroku-token", Env("SHIPR_HEROKU_TOKEN", ""), "Heroku api token.")
	)
	flag.Parse()

	c, err := shipr.New(&shipr.Options{
		DBDir:       *dbdir,
		Env:         *env,
		GitHubToken: *githubToken,
		HerokuToken: *herokuToken,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	log.Fatal(http.ListenAndServe(":"+*port, NewServer(c)))
}

func Env(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		v = fallback
	}
	return v
}
