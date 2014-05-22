package main

type GitHubRepository struct {
	ID   int
	Name string
}

type GitHubDeployment struct {
	ID          int
	Sha         string
	Ref         string
	Name        string
	Environment string
	Payload     map[string]interface{}
	Description string
	Repository  GitHubRepository
}

type GitHubDeploymentStatus struct {
	ID         int
	State      string
	Deployment GitHubDeployment
}
