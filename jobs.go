package shipr

import "time"

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
		repo:        repo,
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
	RepoID      int
	Guid        int
	Sha         string
	Ref         string
	Environment string
	Description string
	Force       bool
	ExitStatus  int

	repo *Repo `db:"-"`
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
