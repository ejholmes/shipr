package shipr

import (
	"io/ioutil"
	"testing"
)

// cleanup truncates all the tables in the database for pristine state
// between tests.
func cleanup(t *testing.T) {
	_, err := dbmap.Exec(`truncate repos cascade`)
	if err != nil {
		t.Fatal(err)
	}
}

// fixture loads a JSON fixture from fixtures/.
func fixture(t *testing.T, path string) []byte {
	raw, err := ioutil.ReadFile("fixtures/" + path)
	if err != nil {
		t.Fatal(err)
	}
	return raw
}
