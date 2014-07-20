package api

import (
	"strconv"

	"github.com/remind101/shipr"
)

type JobResource struct {
	ID          int    `json:"id"`
	Guid        int    `json:"guid"`
	Sha         string `json:"sha"`
	Ref         string `json:"ref"`
	Environment string `json:"environment"`
	Force       bool   `json:"force"`
	ExitStatus  *int   `json:"exit_status"`
}

func NewJobResource(j *shipr.Job) *JobResource {
	return &JobResource{
		ID:          j.ID,
		Guid:        j.Guid,
		Sha:         j.Sha,
		Ref:         j.Ref,
		Environment: j.Environment,
		ExitStatus:  j.ExitStatus,
	}
}

func JobsList(c *shipr.Shipr, res *Response, req *Request) {
	res.Status(200)
	res.Present([]string{})
}

func JobsInfo(c *shipr.Shipr, res *Response, req *Request) {
	id, err := strconv.Atoi(req.Var("id"))
	if err != nil {
		panic(err)
	}
	job, err := c.Jobs.Find(id)
	if err != nil {
		panic(err)
	}

	if job == nil {
		res.NotFound()
		return
	}

	res.Status(200)
	res.Present(NewJobResource(job))
}
