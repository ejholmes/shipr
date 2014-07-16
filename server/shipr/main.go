package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/remind101/shipr"
	"github.com/remind101/shipr/providers/heroku"
	"github.com/remind101/shipr/server"
)

type config struct {
	dbdir, env, port string
	tokens           struct {
		github, heroku string
	}
}

func main() {
	config := parseFlags()

	c, err := shipr.New(&shipr.Options{
		DBDir:       config.dbdir,
		Env:         config.env,
		GitHubToken: config.tokens.github,
	})
	if err != nil {
		log.Fatal(err)
	}

	c.Provider = heroku.New(c.GitHub, config.tokens.heroku)

	defer c.Close()

	log.Fatal(http.ListenAndServe(":"+config.port, server.NewServer(c)))
}

func parseFlags() *config {
	var c config

	flag.StringVar(&c.dbdir, "dbdir", "db", "The db dir containing a dbconf.yml file.")
	flag.StringVar(&c.env, "env", Env("GOENV", "development"), "The environment to run in.")
	flag.StringVar(&c.port, "port", Env("PORT", "3001"), "Port to run the server on.")
	flag.StringVar(&c.tokens.github, "github-token", Env("SHIPR_GITHUB_TOKEN", ""), "GitHub api token.")
	flag.StringVar(&c.tokens.heroku, "heroku-token", Env("SHIPR_HEROKU_TOKEN", ""), "Heroku api token.")
	flag.Parse()

	return &c
}

func Env(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		v = fallback
	}
	return v
}
