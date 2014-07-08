package shipr

// RepoRepository manages the `repos` table.
type RepoRepository struct {
	*DB
}

// Insert inserts the Repo into the database.
func (r *RepoRepository) Insert(repo *Repo) error {
	return r.Insert(repo)
}

// CreateByName creates a new Repo by name.
func (r *RepoRepository) CreateByName(name string) (*Repo, error) {
	repo := &Repo{Name: name}
	return repo, r.Insert(repo)
}

// FindByName trys to find a repo by name. If the repo is not found,
// returns nil.
func (r *RepoRepository) FindByName(name string) (*Repo, error) {
	return r.findBy("name", name)
}

// Find finds a repo by id.
func (r *RepoRepository) Find(id int) (*Repo, error) {
	return r.findBy("id", id)
}

// findBy finds a Repo by a field.
func (r *RepoRepository) findBy(field string, v interface{}) (*Repo, error) {
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
