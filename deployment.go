package shipr

import "time"

// deployment is an implementation of the Deployment interface backed by the
// jobs table.
type deployment struct {
	Datastore *Datastore
	*Job
}

func (d *deployment) Guid() int           { return d.Job.Guid }
func (d *deployment) RepoName() RepoName  { return d.Job.Repo.RepoName() }
func (d *deployment) Sha() string         { return d.Job.Sha }
func (d *deployment) Ref() string         { return d.Job.Ref }
func (d *deployment) Environment() string { return d.Job.Environment }
func (d *deployment) Description() string { return d.Job.Description }

func (d *deployment) AddLine(output string, timestamp time.Time) error {
	_, err := d.Datastore.LogLines.CreateLine(d.Job, output, timestamp)
	return err
}

func (d *deployment) SetExitCode(code int) error {
	d.Job.ExitStatus = &code
	d.Datastore.Jobs.Update(d.Job)
	return nil
}
