package main

import "testing"

// clean truncates all the tables in the database for pristine state
// between tests.
func cleanup(t *testing.T) {
	_, err := dbmap.Exec(`truncate repos cascade`)
	if err != nil {
		t.Fatal(err)
	}
}
