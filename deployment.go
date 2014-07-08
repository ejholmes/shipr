package shipr

import "time"

// Deployment is an interface that's describes information about a deployment.
// This can be used to create jobs, and is also implemented by Job.
type Deployment interface {
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

type Deployable interface {
	Deployment

	AddLine(string, time.Time) error
}
