package shipr

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

	// The deployer we'll use to deploy jobs.
	Deployer

	// Clients.
	GitHub *GitHubClient
}

// Returns a new Shipr context.
func New(options *Options) (*Shipr, error) {
	db, err := NewDB(options.DBDir, options.Env)
	if err != nil {
		return nil, err
	}

	// Setup a client for talking to GitHub.
	gh := NewGitHubClient(options.GitHubToken)

	deployer := NewHerokuDeployer(gh, options.HerokuToken)

	return &Shipr{
		Env:      options.Env,
		DB:       db,
		Deployer: deployer,
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
