package github

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"testing"
)

// fixture loads a JSON fixture from fixtures/.
func fixture(t *testing.T, path string) []byte {
	raw, err := ioutil.ReadFile("fixtures/" + path)
	if err != nil {
		t.Fatal(err)
	}
	return raw
}

func testDeployment(t *testing.T) *Deployment {
	var d Deployment
	err := json.Unmarshal(fixture(t, "deployment.json"), &d)
	if err != nil {
		t.Error(err)
	}
	return &d
}

func Test_Deployment_Deployable(t *testing.T) {
	d := testDeployment(t)
	notDeployable := errors.New("Expected Deployment to be Deployable but it's not.")

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
