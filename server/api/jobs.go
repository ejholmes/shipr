package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

type JobsHandler struct {
	*shipr.Shipr
}

func (h *JobsHandler) Info(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	job, err := h.Jobs.Find(id)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(NewJob(job))
}
