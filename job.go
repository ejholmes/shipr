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

// Methods to implement the Deployment interface.

func (j *Job) RepoName() RepoName {
	repo, err := j.Repo()
	if err != nil {
		return ""
	}
	return repo.RepoName()
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
