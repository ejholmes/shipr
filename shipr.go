package shipr

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/ejholmes/go-github/github"
)

type Shipr struct {
	// The environment (e.g. production, staging, etc..)
	Env string

	// The DB connection.
	*DB

	// The deployer we'll use to deploy jobs.
	Deployer

	// Clients.
	GitHub *github.Client
}

// Returns a new Shipr context.
func New(path, env, githubToken string) (*Shipr, error) {
	db, err := NewDB(path, env)
	if err != nil {
		return nil, err
	}

	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: githubToken},
	}

	// Setup a client for talking to GitHub.
	gh := github.NewClient(t.Client())

	return &Shipr{
		Env:      env,
		DB:       db,
		Deployer: &HerokuDeployer{gh},
		GitHub:   gh,
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
