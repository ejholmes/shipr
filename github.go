package main

// GitHubRepository is embedded inside GitHubDeployment and GitHubDeploymentStatus.
type GitHubRepository struct {
	ID       int
	Name     string
	FullName string `json:"full_name"`
}

// GitHubDeployment represents the webhook payload that GitHub sends us for the
// `deployment` event.
type GitHubDeployment struct {
	ID             int
	RawSha         string `json:"sha"`
	RawRef         string `json:"ref"`
	RawEnvironment string `json:"environment"`
	RawDescription string `json:"description"`
	Name           string
	Payload        map[string]interface{}
	Repository     GitHubRepository
}

// GitHubDeploymentStatus represents the webhook payload that GitHub sends us for the
// `deployment_status` event.
type GitHubDeploymentStatus struct {
	ID         int
	State      string
	Deployment GitHubDeployment
}

// Methods to implement the Deployable interface.

// Guid returns the unique identifier for this github deployment.
func (d *GitHubDeployment) Guid() int {
	return d.ID
}

// RepoName returns the GitHub repository name related to this deployment.
func (d *GitHubDeployment) RepoName() string {
	return d.Repository.FullName
}

// Sha returns the git sha to deploy.
func (d *GitHubDeployment) Sha() string {
	return d.RawSha
}

// Ref returns the friendly ref name for the git sha.
func (d *GitHubDeployment) Ref() string {
	return d.RawRef
}

// Environment returns the deploy environment that was requested.
func (d *GitHubDeployment) Environment() string {
	return d.RawEnvironment
}

// Description returns the description for the deploy.
func (d *GitHubDeployment) Description() string {
	return d.RawDescription
}
