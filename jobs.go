package shipr

// JobStatus represents the status of the job.
type JobStatus int

const (
	// StatusPending means the job is pending.
	StatusPending JobStatus = iota

	// StatusFailed means the job failed.
	StatusFailed

	// StatusSucceeded means the job succeeded.
	StatusSucceeded
)

// JobsService manages the `jobs` table.
type JobsService struct {
	*Datastore
}

// CreateFromDescription creates a new job from an object implementing the description
// interface.
func (s *JobsService) CreateFromDescription(d Description) (*Job, error) {
	repo, err := s.Repos.FindOrCreateByName(d.RepoName().String())
	if err != nil {
		return nil, err
	}

	job := &Job{
		Repo:        repo,
		RepoID:      repo.ID,
		Guid:        d.Guid(),
		Sha:         d.Sha(),
		Ref:         d.Ref(),
		Environment: d.Environment(),
		Description: d.Description(),
	}

	return job, s.Insert(job)
}

// All returns all jobs.
func (s *JobsService) All() ([]*Job, error) {
	var jobs []*Job
	return jobs, s.List("jobs", &jobs)
}

// Find finds a single job by id.
func (s *JobsService) Find(id int) (*Job, error) {
	return s.findBy("id", id)
}

// findBy finds a Job by a field.
func (s *JobsService) findBy(field string, v interface{}) (*Job, error) {
	var job Job

	err := s.Get("jobs", &job, field, v)
	if err != nil {
		return nil, err
	}

	if job.ID == 0 {
		return nil, nil
	}

	return &job, err
}

// Job maps the fields from the `jobs` table.
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

// Status returns the status for this job. Returns StatusPending if the exit code
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

// IsDone returns if the job is done or not.
func (j *Job) IsDone() bool {
	if j.Status() != StatusPending {
		return true
	}
	return false
}
