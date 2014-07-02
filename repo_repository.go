package main

import "github.com/coopernurse/gorp"

// A Repository has methods for adding and removing repos.
type RepoRepository struct {
	dbmap *gorp.DbMap
}

// FindOrCreateByName tries to find the repo by name or it creates it.
func (r *RepoRepository) FindOrCreateByName(name string) (*Repo, error) {
	repo, err := r.FindByName(name)
	if err != nil {
		return nil, err
	}

	if repo == nil {
		repo, err = r.CreateByName(name)
		if err != nil {
			return nil, err
		}
	}

	return repo, nil
}

// CreateByName creates a new Repo by name.
func (r *RepoRepository) CreateByName(name string) (*Repo, error) {
	repo := &Repo{Name: name}

	err := r.Insert(repo)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// Insert inserts the repo into postgres.
func (r *RepoRepository) Insert(repo *Repo) error {
	return r.dbmap.Insert(repo)
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

// Find finds a repo by id.
func (r *RepoRepository) Find(id int) (*Repo, error) {
	var repo Repo

	err := r.dbmap.SelectOne(&repo, `SELECT * FROM repos WHERE id = $1 LIMIT 1`, id)
	if err != nil {
		return nil, err
	}

	if repo.ID == 0 {
		return nil, nil
	}

	return &repo, err
}
