package github

import (
	"bytes"
	"fmt"
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
		payload, event string
		expected       expectation
	}{
		{
			payload: "deployment.json",
			event:   "deployment",
			expected: expectation{
				status: 200,
			},
		},
		{
			payload: "deployment_status.success.json",
			event:   "deployment_status",
			expected: expectation{
				status: 200,
			},
		},
	}

	for _, test := range tests {
		raw, err := ioutil.ReadFile("../../fixtures/github/" + test.payload)
		if err != nil {
			t.Error(err)
		}

		h := New(&fakeShipr{})
		resp := httptest.NewRecorder()

		req, _ := http.NewRequest("POST", "/", bytes.NewReader(raw))
		req.Header.Set("X-GitHub-Event", test.event)

		h.ServeHTTP(resp, req)

		if resp.Code != test.expected.status {
			fmt.Println(resp)
			t.Errorf("Status: Want %v; Got %v", test.expected.status, resp.Code)
		}
	}
}
