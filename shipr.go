package shipr

import "github.com/remind101/shipr/github"

// Options is a struct for passing options to New.
type Options struct {
	Env         string
	DBDir       string
	GitHubToken string
	HerokuToken string
}

// Shipr context.
type Shipr struct {
	// The environment (e.g. production, staging, etc..)
	Env string

	// Access to services.
	*Datastore

	// The Provider we'll use to deploy jobs.
	Provider

	// Clients.
	GitHub github.Client
}

// New returns a new Shipr instance.
func New(options *Options) (*Shipr, error) {
	db, err := NewDB(options.DBDir, options.Env)
	if err != nil {
		return nil, err
	}

	// Setup a client for talking to GitHub.
	g := github.New(options.GitHubToken)

	return &Shipr{
		Env:       options.Env,
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

// Close closes the database connection.
func (c *Shipr) Close() error {
	return c.DB.Close()
}
