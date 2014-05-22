package main

type GitHubRepository struct {
	ID   int `json:"id"`
	Name string
}

type GitHubDeployment struct {
	ID          int `json:"id"`
	Sha         string
	Ref         string
	Name        string
	Environment string
	Payload     map[string]interface{}
	Description string
	Repository  GitHubRepository
}

type GitHubDeploymentStatus struct {
	ID         int `json:"id"`
	State      string
	Deployment GitHubDeployment
}
