package shipr

import "strings"

// ReposService manages the `repos` table.
type ReposService struct {
	*DB
}

// CreateByName creates a new Repo by name.
func (r *ReposService) CreateByName(name string) (*Repo, error) {
	repo := &Repo{Name: name}
	return repo, r.Insert(repo)
}

// FindByName trys to find a repo by name. If the repo is not found,
// returns nil.
func (r *ReposService) FindByName(name string) (*Repo, error) {
	return r.findBy("name", name)
}

// Find finds a repo by id.
func (r *ReposService) Find(id int) (*Repo, error) {
	return r.findBy("id", id)
}

// findBy finds a Repo by a field.
func (r *ReposService) findBy(field string, v interface{}) (*Repo, error) {
	var repo Repo

	err := r.SelectOne(&repo, `SELECT * FROM repos WHERE `+field+` = $1 LIMIT 1`, v)
	if err != nil {
		return nil, err
	}

	if repo.ID == 0 {
		return nil, nil
	}

	return &repo, err
}

// A Repo represents a GitHub repo.
type Repo struct {
	ID   int
	Name string
}

// RepoName returns a RepoName for this Repo.
func (r *Repo) RepoName() RepoName {
	return RepoName(r.Name)
}

// RepoName is value object that represents the <owner/repo> format.
type RepoName string

// Owner returns the owner part of the repo name.
func (n RepoName) Owner() string {
	return n.parts()[0]
}

// Repo returns the repo part of the repo name.
func (n RepoName) Repo() string {
	return n.parts()[1]
}

// Parts returns the owner/repo parts.
func (n RepoName) parts() []string {
	return strings.SplitN(string(n), "/", 2)
}
