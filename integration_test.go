// +build integration

package main

import (
	"bytes"

	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GitHubDeploy(t *testing.T) {
	defer cleanup(t)

	server := NewServer()
	req, _ := http.NewRequest("POST", "/github", bytes.NewReader(fixture(t, "github/deployment.json")))
	req.Header.Set(GitHubEventHeader, "deployment")

	resp := httptest.NewRecorder()
	server.ServeHTTP(resp, req)

	total, _ := jobs.Total()
	if total != 1 {
		t.Error("Expected a job to be created")
	}

	job, err := jobs.First()
	if err != nil {
		t.Fatal(err)
	}

	if job.Sha != "13c6b1509c1c0f6a38cf9994cb510df5d39bb693" {
		t.Error("Expected the Sha to match the GitHub deployment event Sha.")
	}
}
