package main

import (
	"database/sql"
	"log"

	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
)

var (
	// Sql connection
	db *sql.DB

	// db holds a global connection to the database.
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
