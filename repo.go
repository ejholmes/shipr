package main

import "github.com/coopernurse/gorp"

// A Repository has methods for adding and removing repos.
type RepoRepository struct {
	dbmap *gorp.DbMap
}

// A Repo represents a GitHub repo.
type Repo struct {
	ID   int
	Name string
}

// FindOrCreate tries to find the repo by name or it creates it.
func (r *RepoRepository) FindOrCreate(name string) (*Repo, error) {
	repo, err := r.FindByName(name)
	if err != nil {
		return nil, err
	}

	if repo == nil {
		repo, err = r.Create(name)
		if err != nil {
			return nil, err
		}
	}

	return repo, nil
}

// Create creates a new Repo by name.
func (r *RepoRepository) Create(name string) (*Repo, error) {
	repo := &Repo{Name: name}

	err := r.dbmap.Insert(repo)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// FindByName trys to find a repo by name. If the repo is not found,
// returns nil.
func (r *RepoRepository) FindByName(name string) (*Repo, error) {
	var repo Repo

	err := r.dbmap.SelectOne(&repo, `SELECT * FROM repos WHERE name = $1 LIMIT 1`, name)
	if err != nil {
		return nil, err
	}

	if repo.ID == 0 {
		return nil, nil
	}

	return &repo, err
}

// InstallGitHubHook sets the GitHub deployment and deployment_status
// webhook so that we can process these events.
func (r *Repo) InstallGitHubHook() error {
	return nil
}
