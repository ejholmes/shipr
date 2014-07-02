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
	ID             int
	RepoID         int    `db:"repo_id"`
	RawGuid        int    `db:"guid"`
	RawSha         string `db:"sha"`
	Force          bool
	RawDescription string `db:"description"`
	RawEnvironment string `db:"environment"`
	ExitStatus     *int   `db:"exit_status"`

	// Memoized repo instance
	repo *Repo `db:"-"`
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

// CreateFromDescriber takes a Jobable and inserts a new Job.
func (r *JobRepository) CreateFromDescriber(d Describer) (*Job, error) {
	repo, err := Repos.FindOrCreateByName(d.RepoName())
	if err != nil {
		return nil, err
	}

	job := &Job{
		repo:           repo,
		RepoID:         repo.ID,
		RawGuid:        d.Guid(),
		RawSha:         d.Sha(),
		RawEnvironment: d.Environment(),
		RawDescription: d.Description(),
	}

	err = r.Insert(job)
	if err != nil {
		return nil, err
	}

	return job, nil
}

// First finds the first job.
func (r *JobRepository) First() (*Job, error) {
	var job Job

	err := r.dbmap.SelectOne(&job, `SELECT * FROM jobs LIMIT 1`)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

// Insert inserts the Job into the database.
func (r *JobRepository) Insert(job *Job) error {
	return r.dbmap.Insert(job)
}

// Total returns the total number of Jobs.
func (r *JobRepository) Total() (int64, error) {
	count, err := r.dbmap.SelectInt(`SELECT count(*) from jobs`)
	if err != nil {
		return 0, err
	}
	return count, nil
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

// LogLines returns a LogLineRepository scoped to this job.
func (j *Job) LogLines() *LogLineRepository {
	return &LogLineRepository{j, dbmap}
}

// AddLine adds a line of log output to this job.
func (j *Job) AddLine(output string, timestamp time.Time) error {
	_, err := j.LogLines().Add(output, timestamp)
	return err
}

// Returns if the job is done or not.
func (j *Job) IsDone() bool {
	if j.Status() != StatusPending {
		return true
	} else {
		return false
	}
}

// Returns the repo for this job.
func (j *Job) Repo() (*Repo, error) {
	if j.repo == nil {
		repo, err := Repos.Find(j.RepoID)
		if err != nil {
			return nil, err
		}
		j.repo = repo
	}
	return j.repo, nil
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

// Methods to implement the Describer interface.

func (j *Job) RepoName() string {
	repo, err := j.Repo()
	if err != nil {
		return ""
	}
	return repo.Name
}

func (j *Job) Description() string { return j.RawDescription }
func (j *Job) Environment() string { return j.RawEnvironment }
func (j *Job) Guid() int           { return j.RawGuid }
func (j *Job) Sha() string         { return j.RawSha }
func (j *Job) Ref() string         { return "" }

// Run (Deploy) the job.
func (j *Job) Run() error {
	return deployer.Deploy(j)
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

	_, err := r.dbmap.Select(&lines, `SELECT * FROM log_lines WHERE job_id = $1 ORDER BY timestamp`, r.job.ID)
	if err != nil {
		return nil, err
	}

	return lines, nil
}
