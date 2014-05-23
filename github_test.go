package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func fixture(t *testing.T, path string) []byte {
	raw, err := ioutil.ReadFile("fixtures/" + path)
	if err != nil {
		t.Fatal(err)
	}
	return raw
}

func testGitHubDeployment(t *testing.T) *GitHubDeployment {
	var d GitHubDeployment
	err := json.Unmarshal(fixture(t, "deployment.json"), &d)
	if err != nil {
		t.Error(err)
	}
	return &d
}

func Test_GitHubDeployment_RepoName(t *testing.T) {
	d := testGitHubDeployment(t)

	got := d.RepoName()
	want := "remind101/r101-api"
	if got != want {
		t.Error("RepoName() got %v; want %v", got, want)
	}
}
