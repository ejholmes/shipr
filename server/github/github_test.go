package github

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/remind101/shipr"
)

// fakeShipr is a mock implementation of the Shipr interface.
type fakeShipr struct{}

func (sh *fakeShipr) Deploy(d shipr.Description) error {
	return nil
}

func (sh *fakeShipr) Notify(d shipr.Notification) error {
	return nil
}

func Test_GitHub(t *testing.T) {
	type expectation struct {
		status int
	}

	tests := []struct {
		payload   string
		event     string
		signature string
		expected  expectation
	}{
		{
			payload:   "deployment.json",
			event:     "deployment",
			signature: "sha1=invalid",
			expected: expectation{
				status: 403,
			},
		},
		{
			payload:   "deployment.json",
			event:     "deployment",
			signature: "sha1=017ae161492904b9b41244330be19c610428bbab",
			expected: expectation{
				status: 200,
			},
		},
		{
			payload:   "deployment_status.success.json",
			event:     "deployment_status",
			signature: "sha1=invalid",
			expected: expectation{
				status: 403,
			},
		},
		{
			payload:   "deployment_status.success.json",
			event:     "deployment_status",
			signature: "sha1=a57b521fd15ab4542f729b816f27226277d446aa",
			expected: expectation{
				status: 200,
			},
		},
	}

	for i, test := range tests {
		raw, err := ioutil.ReadFile("../../fixtures/github/" + test.payload)
		if err != nil {
			t.Error(err)
		}

		h := New(&fakeShipr{}, "1234")
		resp := httptest.NewRecorder()

		req, _ := http.NewRequest("POST", "/", bytes.NewReader(raw))
		req.Header.Set("X-GitHub-Event", test.event)
		req.Header.Set("X-Hub-Signature", test.signature)

		h.ServeHTTP(resp, req)

		if resp.Code != test.expected.status {
			t.Errorf("Status %v: Want %v; Got %v", i, test.expected.status, resp.Code)
		}
	}
}
