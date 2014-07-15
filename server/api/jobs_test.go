package api

import (
	"reflect"
	"testing"

	"github.com/remind101/shipr"
)

type jobsInfoTest struct {
	id       func(*shipr.Shipr) string
	expected *Response
}

func Test_JobsInfo(t *testing.T) {
	tests := []jobsInfoTest{
		{
			id: func(c *shipr.Shipr) string {
				return "1"
			},
			expected: &Response{
				status: 404,
				resource: &ErrorResponse{
					Error: "Not Found",
				},
			},
		},
	}

	for i, test := range tests {
		c := testShipr(t)
		res := &Response{}
		req := &Request{vars: map[string]string{"id": test.id(c)}}

		JobsInfo(c, res, req)

		if res.status != test.expected.status {
			t.Errorf("%v: Response.status - Want %v, Got %v", i, test.expected.status, res.status)
		}

		if !reflect.DeepEqual(res.resource, test.expected.resource) {
			t.Errorf("%v: Response.resource - Want %v, Got %v", i, test.expected.resource, res.resource)
		}
	}
}
