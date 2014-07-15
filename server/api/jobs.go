package api

import (
	"strconv"

	"github.com/remind101/shipr"
)

// Job Resource
type Job struct {
	ID          int    `json:"id"`
	Guid        int    `json:"guid"`
	Sha         string `json:"sha"`
	Ref         string `json:"ref"`
	Environment string `json:"environment"`
	Force       bool   `json:"force"`
	ExitStatus  *int   `json:"exit_status"`
}

func NewJob(j *shipr.Job) *Job {
	return &Job{
		ID:          j.ID,
		Guid:        j.Guid,
		Sha:         j.Sha,
		Ref:         j.Ref,
		Environment: j.Environment,
		ExitStatus:  j.ExitStatus,
	}
}

func JobsList(c *shipr.Shipr, w ResponseWriter, r *Request) {
}

func JobsInfo(c *shipr.Shipr, w ResponseWriter, r *Request) {
	id, err := strconv.Atoi(r.Var("id"))
	if err != nil {
		panic(err)
	}
	job, err := c.Jobs.Find(id)
	if err != nil {
		panic(err)
	}

	if job == nil {
		w.NotFound()
		return
	}

	w.Status(200)
	w.Present(NewJob(job))
}
