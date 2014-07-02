package github

// Repository is embedded inside GitHubDeployment and GitHubDeploymentStatus.
type Repository struct {
	ID       int
	Name     string
	FullName string `json:"full_name"`
}

// Deployment represents the webhook payload that GitHub sends us for the
// `deployment` event.
type Deployment struct {
	ID             int
	RawSha         string `json:"sha"`
	RawRef         string `json:"ref"`
	RawEnvironment string `json:"environment"`
	RawDescription string `json:"description"`
	Name           string
	Payload        map[string]interface{}
	Repository     Repository
}

// DeploymentStatus represents the webhook payload that GitHub sends us for the
// `deployment_status` event.
type DeploymentStatus struct {
	ID         int
	State      string
	Deployment Deployment
}

// Methods to implement the Deployable interface.

// Guid returns the unique identifier for this github deployment.
func (d *Deployment) Guid() int {
	return d.ID
}

// RepoName returns the GitHub repository name related to this deployment.
func (d *Deployment) RepoName() string {
	return d.Repository.FullName
}

// Sha returns the git sha to deploy.
func (d *Deployment) Sha() string {
	return d.RawSha
}

// Ref returns the friendly ref name for the git sha.
func (d *Deployment) Ref() string {
	return d.RawRef
}

// Environment returns the deploy environment that was requested.
func (d *Deployment) Environment() string {
	return d.RawEnvironment
}

// Description returns the description for the deploy.
func (d *Deployment) Description() string {
	return d.RawDescription
}
