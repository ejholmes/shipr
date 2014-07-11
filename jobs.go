package shipr

import (
	"fmt"
	"time"
)

type JobStatus int

const (
	StatusPending JobStatus = iota
	StatusFailed
	StatusSucceeded
)

// JobsService manages the `jobs` table.
type JobsService struct {
	*DB
}

func (s *JobsService) CreateFromDeployment(d Deployment) (*Job, error) {
	repo, err := s.Repos.FindOrCreateByName(string(d.RepoName()))
	if err != nil {
		return nil, err
	}

	job := &Job{
		Repo:        repo,
		RepoID:      repo.ID,
		Guid:        d.Guid(),
		Sha:         d.Sha(),
		Environment: d.Environment(),
		Description: d.Description(),
	}

	return job, s.Insert(job)
}

type Job struct {
	ID          int
	RepoID      int `db:"repo_id"`
	Guid        int
	Sha         string
	Ref         string
	Environment string
	Description string
	Force       bool
	ExitStatus  *int `db:"exit_status"`

	Repo *Repo `db:"-"`
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

// Returns if the job is done or not.
func (j *Job) IsDone() bool {
	if j.Status() != StatusPending {
		return true
	} else {
		return false
	}
}

// DeploymentJob wraps Job to implement the Deployment interface.
type DeploymentJob struct {
	*Job
}

func (j *DeploymentJob) Guid() int           { return j.Job.Guid }
func (j *DeploymentJob) RepoName() RepoName  { return j.Repo.RepoName() }
func (j *DeploymentJob) Sha() string         { return j.Job.Sha }
func (j *DeploymentJob) Ref() string         { return j.Job.Ref }
func (j *DeploymentJob) Environment() string { return j.Job.Environment }
func (j *DeploymentJob) Description() string { return j.Job.Description }

func (j *DeploymentJob) AddLine(output string, timestamp time.Time) error {
	fmt.Println(output)
	return nil
}

func (j *DeploymentJob) SetExitCode(code int) error {
	fmt.Println(code)
	return nil
}
