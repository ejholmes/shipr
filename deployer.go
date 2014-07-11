package shipr

import (
	"fmt"
	"net/url"

	"github.com/remind101/shipr/github"
	"github.com/remind101/shipr/heroku"
)

// Deployer is an interface that can be implemented for deploying a Deployable to
// some platform.
type Deployer interface {
	Deploy(Deployable) error
}

// herokuDeployer is an implementation of the Deployer interface for deploying
// to Heroku using the Platform API: https://devcenter.heroku.com/articles/platform-api-reference#build
type herokuDeployer struct {
	github github.Client
	heroku heroku.Client
}

// newHerokuDeployer builds a new herokuDeployer and returns it.
func newHerokuDeployer(g github.Client, h heroku.Client) Deployer {
	return &herokuDeployer{g, h}
}

// Deployer implements the Deployer interface. Builds a new herokuDeploy and runs it.
func (h *herokuDeployer) Deploy(d Deployable) error {
	return newHerokuDeploy(h, d).run()
}

// herokuDeploy wraps a deployable for managing the Heroku build process.
type herokuDeploy struct {
	*herokuDeployer
	Deployable
}

// newHerokuDeploy builds a new herokuDeploy and returns it.
func newHerokuDeploy(h *herokuDeployer, d Deployable) *herokuDeploy {
	return &herokuDeploy{h, d}
}

// run runs the build process.
func (d *herokuDeploy) run() error {
	b, err := d.createBuild()
	if err != nil {
		return err
	}

	d.poll(b)

	fmt.Println("Done")
	return nil
}

// polls polls the build output.
func (d *herokuDeploy) poll(b *heroku.Build) {
	lines := make(chan *heroku.BuildResultLine, 100)
	status := make(chan string)

	d.heroku.BuildOutputStream(d.app(), b.ID, lines, status)
	for {
		select {
		case l := <-lines:
			fmt.Println(l.Line)
		case s := <-status:
			fmt.Println(s)
			if s == "succeeded" {
				break
			}
		}
	}
}

// createBuild creates the Heroku build.
func (d *herokuDeploy) createBuild() (*heroku.Build, error) {
	source, err := d.sourceBlob()
	if err != nil {
		return nil, err
	}

	return d.heroku.BuildCreate(d.app(), source.String(), d.Sha())
}

// sourceBlob returns the archive link where the source can be downloaded by the heroku build.
func (d *herokuDeploy) sourceBlob() (*url.URL, error) {
	repoName := d.RepoName()
	return d.github.GetArchiveLink(repoName.Owner(), repoName.Repo(), d.Sha())
}

// app returns the name of the app.
func (d *herokuDeploy) app() string {
	return d.RepoName().Repo()
}
