package main

import "errors"

// A Repo represents a GitHub repo.
type Repo struct {
	ID   int
	Name string
}

// CreateRepo creates a new Repo by name.
func CreateRepo(name string) (*Repo, error) {
	repo := &Repo{Name: name}

	err := dbmap.Insert(repo)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// FindRepo trys to find a repo by name. If the repo is not found,
// returns nil.
func FindRepo(name string) (*Repo, error) {
	var r Repo

	err := dbmap.SelectOne(&r, `SELECT * FROM repos WHERE name = $1`, name)
	if err != nil {
		return nil, err
	}

	if r.ID == 0 {
		return nil, errors.New("Not found.")
	}

	return &r, err
}
