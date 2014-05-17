package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "dbname=shipr_dev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(db)
}
