package shipr

type JobStatus int

const (
	StatusPending JobStatus = iota
	StatusFailed
	StatusSucceeded
)

// JobsService manages the `jobs` table.
type JobsService struct {
	*Datastore
}

func (s *JobsService) CreateFromDescription(d Description) (*Job, error) {
	repo, err := s.Repos.FindOrCreateByName(string(d.RepoName()))
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

func (s *JobsService) Find(id int) (*Job, error) {
	return s.findBy("id", id)
}

// findBy finds a Job by a field.
func (s *JobsService) findBy(field string, v interface{}) (*Job, error) {
	var job Job

	err := s.Get(&job, field, v)
	if err != nil {
		return nil, err
	}

	if job.ID == 0 {
		return nil, nil
	}

	return &job, err
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

func (j *Job) table() string {
	return "jobs"
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
