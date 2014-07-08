package shipr

import "time"

// Job is an implementation of the Deployment/Deployable interface that is backed
// by postgres.
type Job struct {
	ID             int    `db:"id"`
	RawGuid        int    `db:"guid"`
	RawSha         string `db:"sha"`
	RawRef         string `db:"ref"`
	RawEnvironment string `db:"environment"`
	RawDescription string `db:"description"`
}

func (j *Job) Guid() int           { return j.RawGuid }
func (j *Job) RepoName() RepoName  { return RepoName("") }
func (j *Job) Sha() string         { return j.RawSha }
func (j *Job) Ref() string         { return j.RawRef }
func (j *Job) Environment() string { return j.RawEnvironment }
func (j *Job) Description() string { return j.RawDescription }

func (j *Job) AddLine(output string, timestamp time.Time) error {
	return nil
}
