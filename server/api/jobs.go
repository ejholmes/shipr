package api

import (
	"strconv"

	"github.com/remind101/shipr"
)

// job is the api representation of a shipr.Job.
type job struct {
	ID          int     `json:"id"`
	Guid        int     `json:"guid"`
	Sha         string  `json:"sha"`
	Ref         string  `json:"ref"`
	Environment string  `json:"environment"`
	Force       bool    `json:"force"`
	ExitStatus  *int    `json:"exit_status"`
	Status      string  `json:"status"`
	Output      *string `json:"output,omitempty"`
}

// Job decorates a shipr.Job as a Job resource.
func Job(j *shipr.Job) *job {
	return &job{
		ID:          j.ID,
		Guid:        j.Guid,
		Sha:         j.Sha,
		Ref:         j.Ref,
		Environment: j.Environment,
		ExitStatus:  j.ExitStatus,
		Status:      j.Status().String(),
	}
}

// Jobs decorates a slice of shipr.Job as a Job resource.
func Jobs(jobs []*shipr.Job) []*job {
	r := make([]*job, len(jobs))
	for i, j := range jobs {
		r[i] = Job(j)
	}
	return r
}

// JobsList responds with all jobs.
func JobsList(c *shipr.Shipr, res *Response, req *Request) {
	jobs, err := c.Jobs.All()
	if err != nil {
		panic(err)
	}
	res.Status(200)
	res.Present(Jobs(jobs))
}

// JobsInfo responds with a single job.
func JobsInfo(c *shipr.Shipr, res *Response, req *Request) {
	id, err := strconv.Atoi(req.Var("id"))
	if err != nil {
		panic(err)
	}
	j, err := c.Jobs.Find(id)
	if err != nil {
		panic(err)
	}

	if j == nil {
		res.NotFound()
		return
	}

	output, err := c.LogLines.Output(j)
	if err != nil {
		panic(err)
	}

	resource := Job(j)
	resource.Output = &output

	res.Status(200)
	res.Present(resource)
}
