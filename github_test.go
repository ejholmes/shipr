package main

import (
	"encoding/json"
	"errors"
	"testing"
)

func testGitHubDeployment(t *testing.T) *GitHubDeployment {
	var d GitHubDeployment
	err := json.Unmarshal(fixture(t, "deployment.json"), &d)
	if err != nil {
		t.Error(err)
	}
	return &d
}

func Test_GitHubDeployment_Deployable(t *testing.T) {
	d := testGitHubDeployment(t)
	notDeployable := errors.New("Expected GitHubDeployment to be Deployable but it's not.")

	if d.Guid() != 11939 {
		t.Error(notDeployable)
	}

	if d.RepoName() != "remind101/r101-api" {
		t.Error(notDeployable)
	}

	if d.Sha() != "13c6b1509c1c0f6a38cf9994cb510df5d39bb693" {
		t.Error(notDeployable)
	}

	if d.Ref() != "develop" {
		t.Error(notDeployable)
	}

	if d.Environment() != "production" {
		t.Error(notDeployable)
	}

	if d.Description() != "" {
		t.Error(notDeployable)
	}
}
