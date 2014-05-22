package main

// clean truncates all the tables in the database for pristine state
// between tests.
func clean() {
	dbmap.Exec(`truncate repos cascade`)
}
