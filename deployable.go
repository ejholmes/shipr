package shipr

// Deployable is an interface that's used for creating Job's. We implement this
// interface on the GitHubDeployment struct, so that we can deploy github's
// deployment events directly.
type Deployable interface {
	// Guid should return a unique identifier for this deployment.
	Guid() int

	// RepoName should return the string name of the repo to deploy.
	RepoName() string

	// Sha should return the git sha that we want to deploy.
	Sha() string

	// Ref should return the git ref that is being requested.
	Ref() string

	// Environment should return the name of the environment that the repo is being deploy to.
	Environment() string

	// Description should return a string description for the deploy.
	Description() string
}
