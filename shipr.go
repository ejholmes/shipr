package shipr

type Shipr struct {
	// The environment (e.g. production, staging, etc..)
	Env string

	// The DB connection.
	*DB

	// Repositories
	Repos *RepoRepository
}

// Returns a new Shipr context.
func New(path, env string) (*Shipr, error) {
	db, err := NewDB(path, env)
	if err != nil {
		return nil, err
	}
	return &Shipr{
		Env:   env,
		DB:    db,
		Repos: &RepoRepository{},
	}, nil
}

// Deploy is the primary interface into deploy things. It takes an object that conforms
// to the Deployment interface, creates a Job then runs it.
func (c *Shipr) Deploy(d Deployment) error {
	return nil
}

func (c *Shipr) Close() error {
	return c.DB.Close()
}
