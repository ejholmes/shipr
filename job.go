package main

import "time"

type JobStatus int

const (
	StatusPending JobStatus = iota
	StatusFailed
	StatusSucceeded
)

// Job is our reference to a deployment.
type Job struct {
	ID          int
	Sha         string
	Environment string
	ExitStatus  *int
}

func (j *Job) AppendOutput(output string, timestamp time.Time) error {
	// Create a log_lines row
	return nil
}

// Returns if the job is done or not.
func (j *Job) IsDone() bool {
	if j.Status() != StatusPending {
		return true
	} else {
		return false
	}
}

// Returns the status for this job.
func (j *Job) Status() (status JobStatus) {
	if j.ExitStatus != nil {
		if *j.ExitStatus == 0 {
			status = StatusSucceeded
		} else {
			status = StatusFailed
		}
	}
	return
}
