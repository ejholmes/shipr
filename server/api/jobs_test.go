package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/remind101/shipr"
)

type jobsInfoTest struct {
	id       string
	expected *Response
}

func Test_JobsInfo(t *testing.T) {
	c := testShipr(t)
	c.Reset()

	// Insert some test data.
	repo := &shipr.Repo{}
	err := c.Repos.Insert(repo)
	if err != nil {
		t.Fatal(err)
	}

	job := &shipr.Job{Repo: repo, RepoID: repo.ID}
	err = c.Jobs.Insert(job)
	if err != nil {
		t.Fatal(err)
	}

	tests := []jobsInfoTest{
		{
			id: "1024",
			expected: &Response{
				status: 404,
				resource: &ErrorResponse{
					Error: "Not Found",
				},
			},
		},
		{
			id: fmt.Sprintf("%v", job.ID),
			expected: &Response{
				status: 200,
				resource: &Job{
					ID: 1,
				},
			},
		},
	}

	for i, test := range tests {
		res := &Response{}
		req := &Request{vars: map[string]string{"id": test.id}}

		JobsInfo(c, res, req)

		if res.status != test.expected.status {
			t.Errorf("%v: Response.status - Want %v, Got %v", i, test.expected.status, res.status)
		}

		if !reflect.DeepEqual(res.resource, test.expected.resource) {
			t.Errorf("%v: Response.resource - Want %v, Got %v", i, test.expected.resource, res.resource)
		}
	}
}
