package main

import "github.com/coopernurse/gorp"

// JobRepository has methods for adding and removing jobs.
type JobRepository struct {
	dbmap *gorp.DbMap
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
