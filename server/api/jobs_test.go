package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/remind101/shipr"
)

type jobsInfoTest struct {
	id       func(*testing.T, *shipr.Shipr) string
	expected *Response
}

func Test_JobsInfo(t *testing.T) {
	tests := []jobsInfoTest{
		{
			id: func(t *testing.T, c *shipr.Shipr) string {
				return "1"
			},
			expected: &Response{
				status: 404,
				resource: &ErrorResponse{
					Error: "Not Found",
				},
			},
		},
		{
			id: func(t *testing.T, c *shipr.Shipr) string {
				r := &shipr.Repo{}
				err := c.Repos.Insert(r)
				if err != nil {
					t.Fatal(err)
				}

				j := &shipr.Job{Repo: r, RepoID: r.ID}
				err = c.Jobs.Insert(j)
				if err != nil {
					t.Fatal(err)
				}
				return fmt.Sprintf("%v", j.ID)
			},
			expected: &Response{
				status: 200,
				resource: &Job{
					ID: 1,
				},
			},
		},
	}

	for i, test := range tests {
		c := testShipr(t)
		c.Reset()

		res := &Response{}
		req := &Request{vars: map[string]string{"id": test.id(t, c)}}

		JobsInfo(c, res, req)

		if res.status != test.expected.status {
			t.Errorf("%v: Response.status - Want %v, Got %v", i, test.expected.status, res.status)
		}

		if !reflect.DeepEqual(res.resource, test.expected.resource) {
			t.Errorf("%v: Response.resource - Want %v, Got %v", i, test.expected.resource, res.resource)
		}
	}
}
