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

// LogLine represents a line of log output from the deploy job.
type LogLine struct {
	JobID     int `db:"job_id"`
	Output    string
	Timestamp time.Time
}

// Output returns the log output for this job.
func (j *Job) Output() string {
	return ""
}

// AddLine adds a line of log output to this job.
func (j *Job) AddLine(output string, timestamp time.Time) error {
	query := `INSERT INTO log_lines (job_id, output, timestamp) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, j.ID, output, timestamp)
	if err != nil {
		return err
	}
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
