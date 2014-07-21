package github

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ejholmes/go-github/github"
	"github.com/remind101/shipr"
)

func fixture(t *testing.T, path string) []byte {
	raw, err := ioutil.ReadFile("../../fixtures/github/" + path)
	if err != nil {
		t.Error(err)
		return nil
	}
	return raw
}

// fakeShipr is a mock implementation of the Shipr interface.
type fakeShipr struct{}

func (sh *fakeShipr) Deploy(d shipr.Description) error {
	return nil
}

func (sh *fakeShipr) Notify(d shipr.Notification) error {
	return nil
}

func Test_GitHub(t *testing.T) {
	tests := []struct {
		payload   string
		event     string
		signature string

		status int
	}{
		{
			payload:   "deployment.json",
			event:     "deployment",
			signature: "sha1=invalid",
			status:    403,
		},
		{
			payload:   "deployment.json",
			event:     "deployment",
			signature: "sha1=017ae161492904b9b41244330be19c610428bbab",
			status:    200,
		},
		{
			payload:   "deployment_status.success.json",
			event:     "deployment_status",
			signature: "sha1=invalid",
			status:    403,
		},
		{
			payload:   "deployment_status.success.json",
			event:     "deployment_status",
			signature: "sha1=e8a9f11c1cac3bd321be0a4524c67a27b908a880",
			status:    200,
		},
	}

	for i, test := range tests {
		h := New(&fakeShipr{}, "1234")
		resp := httptest.NewRecorder()

		req, _ := http.NewRequest("POST", "/", bytes.NewReader(fixture(t, test.payload)))
		req.Header.Set("X-GitHub-Event", test.event)
		req.Header.Set("X-Hub-Signature", test.signature)

		h.ServeHTTP(resp, req)

		if resp.Code != test.status {
			t.Errorf("Status %v: Want %v; Got %v", i, test.status, resp.Code)
		}
	}
}

func Test_deployment(t *testing.T) {
	tests := []struct {
		fixture string

		guid        int
		repoName    string
		sha         string
		ref         string
		environment string
		description string
	}{
		{
			fixture:     "deployment.json",
			guid:        29104,
			repoName:    "ejholmes/acme-inc",
			sha:         "208db1b5d2bb89b0ff0b79cb7f702e21a750f3fc",
			environment: "production",
			description: "",
		},
	}

	for i, test := range tests {
		var gd github.Deployment
		json.Unmarshal(fixture(t, test.fixture), &gd)

		d := newDeployment(&gd)

		if d.Guid() != test.guid {
			t.Errorf("Guid %v: Want %v; Got %v", i, test.guid, d.Guid())
		}

		if d.RepoName().String() != test.repoName {
			t.Errorf("RepoName %v: Want %v; Got %v", i, test.repoName, d.RepoName())
		}

		if d.Sha() != test.sha {
			t.Errorf("Sha %v: Want %v; Got %v", i, test.sha, d.Sha())
		}

		if d.Environment() != test.environment {
			t.Errorf("Environment %v: Want %v; Got %v", i, test.environment, d.Environment())
		}

		if d.Description() != test.description {
			t.Errorf("Description %v: Want %v; Got %v", i, test.description, d.Description())
		}
	}
}

func Test_deploymentStatus(t *testing.T) {
	tests := []struct {
		fixture string

		guid        int
		repoName    string
		sha         string
		ref         string
		environment string
		description string
		url         string
		user        string
	}{
		{
			fixture:     "deployment_status.success.json",
			guid:        29104,
			repoName:    "ejholmes/acme-inc",
			sha:         "208db1b5d2bb89b0ff0b79cb7f702e21a750f3fc",
			environment: "production",
			description: "",
			url:         "http://shipr.test/1234",
			user:        "ejholmes",
		},
	}

	for i, test := range tests {
		var gd github.DeploymentStatus
		json.Unmarshal(fixture(t, test.fixture), &gd)

		d := newDeploymentStatus(&gd)

		if d.Guid() != test.guid {
			t.Errorf("Guid %v: Want %v; Got %v", i, test.guid, d.Guid())
		}

		if d.RepoName().String() != test.repoName {
			t.Errorf("RepoName %v: Want %v; Got %v", i, test.repoName, d.RepoName())
		}

		if d.Sha() != test.sha {
			t.Errorf("Sha %v: Want %v; Got %v", i, test.sha, d.Sha())
		}

		if d.Environment() != test.environment {
			t.Errorf("Environment %v: Want %v; Got %v", i, test.environment, d.Environment())
		}

		if d.Description() != test.description {
			t.Errorf("Description %v: Want %v; Got %v", i, test.description, d.Description())
		}

		if d.URL().String() != test.url {
			t.Errorf("URL %v: Want %v; Got %v", i, test.url, d.URL())
		}

		if d.User() != test.user {
			t.Errorf("User %v: Want %v; Got %v", i, test.user, d.User())
		}
	}
}
