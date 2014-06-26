package main

import (
	"time"

	"github.com/coopernurse/gorp"
)

type JobStatus int

const (
	StatusPending JobStatus = iota
	StatusFailed
	StatusSucceeded
)

// JobRepository has methods for adding and removing jobs.
type JobRepository struct {
	dbmap *gorp.DbMap
}

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

// LogLineRepository has methods for adding and removing log lines.
type LogLineRepository struct {
	job   *Job
	dbmap *gorp.DbMap
}

// LogLine represents a line of log output from the deploy job.
type LogLine struct {
	ID        int
	JobID     int `db:"job_id"`
	Output    string
	Timestamp time.Time
}

// CreateByDeployable takes a Deployable and inserts a new Job.
func (r *JobRepository) CreateByDeployable(d Deployable) (*Job, error) {
	repo, err := repos.FindOrCreateByName(d.RepoName())
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

	err = r.Insert(job)
	if err != nil {
		return nil, err
	}

	return job, nil
}

// Insert inserts the job into the database.
func (r *JobRepository) Insert(job *Job) error {
	return r.dbmap.Insert(job)
}

// Output returns the log output for this job.
func (j *Job) Output() (string, error) {
	lines, err := j.LogLines().All()
	if err != nil {
		return "", err
	}

	output := ""
	for _, l := range lines {
		output += l.Output
	}

	return output, nil
}

func (j *Job) LogLines() *LogLineRepository {
	return &LogLineRepository{j, dbmap}
}

// AddLine adds a line of log output to this job.
func (j *Job) AddLine(output string, timestamp time.Time) (*LogLine, error) {
	return j.LogLines().Add(output, timestamp)
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

// Run (Deploy) the job.
func (j *Job) Run() error {
	return herokuDeployer.Deploy(j)
}

// Add adds a LogLine.
func (r *LogLineRepository) Add(output string, timestamp time.Time) (*LogLine, error) {
	l := &LogLine{JobID: r.job.ID, Output: output, Timestamp: timestamp}
	err := r.Insert(l)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// Insert inserts a LogLine.
func (r *LogLineRepository) Insert(logLine *LogLine) error {
	return r.dbmap.Insert(logLine)
}

// AllForJob returns a slice of LogLine for the Job.
func (r *LogLineRepository) All() ([]LogLine, error) {
	var lines []LogLine

	_, err := dbmap.Select(&lines, `SELECT * FROM log_lines WHERE job_id = $1 ORDER BY timestamp`, r.job.ID)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
