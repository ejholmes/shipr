package main

import (
	"io/ioutil"
	"testing"
	"time"
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

// A fake Deployment implementation.
type FakeDeployment struct {
	guid                                         int
	description, environment, ref, repoName, sha string
}

func (d *FakeDeployment) AddLine(output string, timestamp time.Time) error {
	return nil
}

func (d *FakeDeployment) Description() string { return d.description }
func (d *FakeDeployment) Environment() string { return d.environment }
func (d *FakeDeployment) Guid() int           { return d.guid }
func (d *FakeDeployment) Ref() string         { return d.ref }
func (d *FakeDeployment) RepoName() RepoName  { return RepoName(d.repoName) }
func (d *FakeDeployment) Sha() string         { return d.sha }
