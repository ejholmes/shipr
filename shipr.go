package shipr

import "github.com/remind101/shipr/clients/github"

type Shipr struct {
	// The environment (e.g. production, staging, etc..)
	Env string

	// The DB connection.
	*DB

	// Repositories.
	Repos *RepoRepository
	Jobs  *JobRepository

	// The deployer we'll use to deploy jobs.
	Deployer

	// Clients.
	Github *github.Client
}

// Returns a new Shipr context.
func New(path, env string) (*Shipr, error) {
	db, err := NewDB(path, env)
	if err != nil {
		return nil, err
	}

	// Setup a client for talking to GitHub.
	gh := github.NewClient(nil)

	return &Shipr{
		Env:    env,
		DB:     db,
		Repos:  &RepoRepository{db},
		Jobs:   &JobRepository{db},
		Github: gh,
	}, nil
}

// Deploy is the primary interface into deploying things. It takes an object that conforms
// to the Deployment interface, creates a Job then runs it.
func (c *Shipr) Deploy(d Deployment) error {
	j, err := c.Jobs.CreateFromDeployment(d)
	if err != nil {
		return err
	}
	return c.Deployer.Deploy(&DeploymentJob{j})
}

func (c *Shipr) Close() error {
	return c.DB.Close()
}
