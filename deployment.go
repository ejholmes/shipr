package shipr

import (
	"fmt"
	"time"
)

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

type Deployment interface {
	Description

	AddLine(string, time.Time) error
	SetExitCode(int) error
}

// deployment is an implementation of the Deployment interface backed by the
// jobs table.
type deployment struct {
	*Job
}

func (j *deployment) Guid() int           { return j.Job.Guid }
func (j *deployment) RepoName() RepoName  { return j.Repo.RepoName() }
func (j *deployment) Sha() string         { return j.Job.Sha }
func (j *deployment) Ref() string         { return j.Job.Ref }
func (j *deployment) Environment() string { return j.Job.Environment }
func (j *deployment) Description() string { return j.Job.Description }

func (j *deployment) AddLine(output string, timestamp time.Time) error {
	fmt.Println(output)
	return nil
}

func (j *deployment) SetExitCode(code int) error {
	fmt.Println(code)
	return nil
}
