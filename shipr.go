package shipr

import "github.com/remind101/shipr/github"

type Options struct {
	Env         string
	DBDir       string
	GitHubToken string
	HerokuToken string
}

type Shipr struct {
	// The environment (e.g. production, staging, etc..)
	Env string

	// Access to services.
	DB DB
	*Datastore

	// The Provider we'll use to deploy jobs.
	Provider

	// Clients.
	GitHub github.Client
}

// Returns a new Shipr context.
func New(options *Options) (*Shipr, error) {
	db, err := NewDB(options.DBDir, options.Env)
	if err != nil {
		return nil, err
	}

	// Setup a client for talking to GitHub.
	g := github.New(options.GitHubToken)

	return &Shipr{
		Env:       options.Env,
		DB:        db,
		Datastore: NewDatastore(db),
		GitHub:    g,
	}, nil
}

// Deploy is the primary interface into deploying things. It takes an object that conforms
// to the Description interface, creates a Job then runs it.
func (c *Shipr) Deploy(d Description) error {
	j, err := c.Jobs.CreateFromDescription(d)
	if err != nil {
		return err
	}
	return c.Provider.Deploy(&deployment{c.Datastore, j})
}

func (c *Shipr) Close() error {
	return c.DB.Close()
}
