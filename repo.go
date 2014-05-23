package main

// A Repo represents a GitHub repo.
type Repo struct {
	ID   int
	Name string
}

// CreateRepo creates a new Repo by name.
func CreateRepo(name string, hook bool) (*Repo, error) {
	repo := &Repo{Name: name}

	err := dbmap.Insert(repo)
	if err != nil {
		return nil, err
	}

	if hook {
		repo.InstallGitHubHook()
	}

	return repo, nil
}

// FindRepo trys to find a repo by name. If the repo is not found,
// returns nil.
func FindRepo(name string) (*Repo, error) {
	var r Repo

	err := dbmap.SelectOne(&r, `SELECT * FROM repos WHERE name = $1 LIMIT 1`, name)
	if err != nil {
		return nil, err
	}

	if r.ID == 0 {
		return nil, nil
	}

	return &r, err
}

// InstallGitHubHook sets the GitHub deployment and deployment_status
// webhook so that we can process these events.
func (r *Repo) InstallGitHubHook() error {
	return nil
}
