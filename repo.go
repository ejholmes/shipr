package shipr

// A Repo represents a GitHub repo.
type Repo struct {
	ID   int
	Name string
}

// RepoName returns a RepoName for this Repo.
func (r *Repo) RepoName() RepoName {
	return RepoName(r.Name)
}
