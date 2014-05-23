package main

import (
	"database/sql"
	"log"

	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

// Implementing this interface means that we're deployable.
type Deployable interface {
	// Guid should return a unique identifier for this deployment.
	Guid() int

	// RepoName should return the string name of the repo to deploy.
	RepoName() string

	// GitSha should return the git sha that we want to deploy.
	Sha() string

	// Ref should return the git ref that is being requested.
	Ref() string

	// Environment should return the name of the environment that the repo is being deploy to.
	Environment() string

	// Description should return a string description for the deploy.
	Description() string
}

var (
	// Database
	db    *sql.DB
	dbmap *gorp.DbMap
)

func init() {
	conn, err := sql.Open("postgres", "dbname=shipr_dev sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	dbmap = &gorp.DbMap{Db: conn, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Repo{}, "repos").SetKeys(true, "ID")
	dbmap.AddTableWithName(Job{}, "jobs").SetKeys(true, "ID")
	dbmap.AddTableWithName(LogLine{}, "log_lines").SetKeys(true, "ID")
}

func main() {
	defer db.Close()

	server := NewServer()
	server.Run()
}

func Deploy(d Deployable) error {
	_, err := CreateJob(d)
	if err != nil {
		return err
	}
	return nil
}
