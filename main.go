package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// db holds a global connection to the database.
var db *sqlx.DB

func init() {
	db = sqlx.MustConnect("postgres", "dbname=shipr_dev sslmode=disable")
}

func main() {
	defer db.Close()
}
