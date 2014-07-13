package shipr

import (
	"github.com/remind101/shipr/github"
	"github.com/remind101/shipr/heroku"
)

type Options struct {
	Env         string
	DBDir       string
	GitHubToken string
	HerokuToken string
}

type Shipr struct {
	// The environment (e.g. production, staging, etc..)
	Env string

	// The DB connection.
	*DB

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

	// Setup a client for talking to GitHub and Heroku.
	g := github.New(options.GitHubToken)
	h := heroku.New(options.HerokuToken)

	// Setup the Heroku deployer.
	provider := newHerokuProvider(g, h)

	return &Shipr{
		Env:      options.Env,
		DB:       db,
		Provider: provider,
		GitHub:   g,
	}, nil
}

// Deploy is the primary interface into deploying things. It takes an object that conforms
// to the Description interface, creates a Job then runs it.
func (c *Shipr) Deploy(d Description) error {
	j, err := c.Jobs.CreateFromDescription(d)
	if err != nil {
		return err
	}
	return c.Provider.Deploy(&deployment{c.DB, j})
}

func (c *Shipr) Close() error {
	return c.DB.Close()
}
