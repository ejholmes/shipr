package shipr

// JobRepository manages the `jobs` table.
type JobRepository struct {
	*DB
}

func (r *JobRepository) CreateFromDeployment(d Deployment) (*Job, error) {
	return &Job{}, nil
}

// Insert inserts the Job into the database.
func (r *JobRepository) Insert(job *Job) error {
	return r.Insert(job)
}
