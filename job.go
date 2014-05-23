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
	RepoID      int `db:"repo_id"`
	Guid        int
	Sha         string
	Force       bool
	Description string
	Environment string
	ExitStatus  *int `db:"exit_status"`
}

// LogLine represents a line of log output from the deploy job.
type LogLine struct {
	ID        int
	JobID     int `db:"job_id"`
	Output    string
	Timestamp time.Time
}

// CreateJob takes a Deployable and inserts a new Job.
func CreateJob(d Deployable) (*Job, error) {
	repo, err := FindOrCreateRepo(d.RepoName())
	if err != nil {
		return nil, err
	}

	job := &Job{
		RepoID:      repo.ID,
		Guid:        d.Guid(),
		Sha:         d.Sha(),
		Environment: d.Environment(),
		Description: d.Description(),
	}

	err = dbmap.Insert(job)
	if err != nil {
		return nil, err
	}

	return job, nil
}

// Output returns the log output for this job.
func (j *Job) Output() (string, error) {
	var lines []LogLine

	_, err := dbmap.Select(&lines, `SELECT * FROM log_lines WHERE job_id = $1 ORDER BY timestamp`, j.ID)
	if err != nil {
		return "", nil
	}

	output := ""
	for _, l := range lines {
		output += l.Output
	}

	return output, nil
}

// AddLine adds a line of log output to this job.
func (j *Job) AddLine(output string, timestamp time.Time) (*LogLine, error) {
	l := &LogLine{JobID: j.ID, Output: output, Timestamp: timestamp}
	err := dbmap.Insert(l)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// Returns if the job is done or not.
func (j *Job) IsDone() bool {
	if j.Status() != StatusPending {
		return true
	} else {
		return false
	}
}

// Returns the status for this job. Returns StatusPending if the exit code
// is nil.
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
