package shipr

import "strings"

// ReposService manages the `repos` table.
type ReposService struct {
	*DB
}

// FindOrCreateByName tries to find the repo by name or it creates it.
func (s *ReposService) FindOrCreateByName(name string) (*Repo, error) {
	repo, err := s.FindByName(name)
	if err != nil {
		return nil, err
	}

	if repo == nil {
		return s.CreateByName(name)
	}

	return repo, nil
}

// CreateByName creates a new Repo by name.
func (s *ReposService) CreateByName(name string) (*Repo, error) {
	repo := &Repo{Name: name}
	return repo, s.Insert(repo)
}

// FindByName trys to find a repo by name. If the repo is not found,
// returns nil.
func (s *ReposService) FindByName(name string) (*Repo, error) {
	return s.findBy("name", name)
}

// Find finds a repo by id.
func (s *ReposService) Find(id int) (*Repo, error) {
	return s.findBy("id", id)
}

// findBy finds a Repo by a field.
func (s *ReposService) findBy(field string, v interface{}) (*Repo, error) {
	var repo Repo

	err := s.SelectOne(&repo, `SELECT * FROM repos WHERE `+field+` = $1 LIMIT 1`, v)
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
