package shipr

import "time"

type Job struct {
	ID          int
	Guid        int
	Sha         string
	Ref         string
	Environment string
	Description string
}

// DeploymentJob wraps Job to implement the Deployment interface.
type DeploymentJob struct {
	*Job
}

func (j *DeploymentJob) Guid() int           { return j.Job.Guid }
func (j *DeploymentJob) RepoName() RepoName  { return RepoName("") }
func (j *DeploymentJob) Sha() string         { return j.Job.Sha }
func (j *DeploymentJob) Ref() string         { return j.Job.Ref }
func (j *DeploymentJob) Environment() string { return j.Job.Environment }
func (j *DeploymentJob) Description() string { return j.Job.Description }

func (j *DeploymentJob) AddLine(output string, timestamp time.Time) error {
	return nil
}
