package shipr

import (
	"net/url"
	"time"

	"github.com/remind101/shipr/github"
)

// Description is an interface that's describes information about a deployment.
// This can be used to create jobs, and is also implemented by Job.
type Description interface {
	// Guid should return a unique identifier for this deployment.
	Guid() int

	// RepoName should return the string name of the repo to deploy.
	RepoName() RepoName

	// Sha should return the git sha that we want to deploy.
	Sha() string

	// Ref should return the git ref that is being requested.
	Ref() string

	// Environment should return the name of the environment that the repo is being deploy to.
	Environment() string

	// Description should return a string description for the deploy.
	Description() string
}

// Deployment is an interface that is sent to Providers when deploying. Deployments implement the
// Description interface and also methods for updating information about the deployment.
type Deployment interface {
	Description

	AddLine(string, time.Time) error
	SetExitCode(int) error
}

// Notification is an interface that is provided to Notifiers when there is a status
// update on a deploy.
type Notification interface {
	Description

	URL() *url.URL
	User() string
	State() string
}

// Provider is an interface that can be implemented for deploying a Deployment to
// some platform.
type Provider interface {
	Deploy(Deployment) error
}

// Notifier is an interface that notifiers can implement to forward notifications
// to an external system, like HipChat or Slack.
type Notifier interface {
	Notify(Notification) error
}

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

	// The Notifier to use to send notifications.
	Notifier

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
