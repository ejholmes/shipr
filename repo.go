package main

// A Repo represents a GitHub repo.
type Repo struct {
	ID   int
	Name string
}

// InstallGitHubHook sets the GitHub deployment and deployment_status
// webhook so that we can process these events.
func (r *Repo) InstallGitHubHook() error {
	return nil
}

func (r *Repo) RepoName() RepoName {
	return RepoName(r.Name)
}
