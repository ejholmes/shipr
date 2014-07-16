package shipr

import "time"

// Description is an interface that's describes information about a deployment.
// This can be used to create jobs, and is also implemented by Job.
type Description interface {
	// Guid should return a unique identifier for this deployment.
	Guid() int

	// RepoName should return the string name of the repo to deploy.
	RepoName() RepoName

	// Sha should return the git sha that we want to deploy.
	Sha() string

	// Ref should return the git ref that is being requested.
	Ref() string

	// Environment should return the name of the environment that the repo is being deploy to.
	Environment() string

	// Description should return a string description for the deploy.
	Description() string
}

// Deployment is an interface that is sent to Providers when deploying. Deployments implement the
// Description interface and also methods for updating information about the deployment.
type Deployment interface {
	Description

	AddLine(string, time.Time) error
	SetExitCode(int) error
}

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
