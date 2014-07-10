package shipr

import (
	"fmt"
	"net/url"

	"github.com/remind101/shipr/github"
	"github.com/remind101/shipr/heroku"
)

type Deployer interface {
	Deploy(Deployable) error
}

type HerokuDeployer struct {
	github github.Client
	heroku heroku.Client
}

func NewHerokuDeployer(g github.Client, h heroku.Client) *HerokuDeployer {
	return &HerokuDeployer{g, h}
}

func (h *HerokuDeployer) Deploy(d Deployable) error {
	hd := &HerokuDeploy{h, d}
	return hd.Run()
}

type HerokuDeploy struct {
	*HerokuDeployer
	Deployable
}

func (d *HerokuDeploy) Run() error {
	build, err := d.CreateBuild()
	if err != nil {
		return err
	}
	fmt.Println(build)
	return nil
}

func (d *HerokuDeploy) CreateBuild() (*heroku.Build, error) {
	source, err := d.SourceBlob()
	if err != nil {
		return nil, err
	}

	return d.heroku.BuildCreate(d.App(), source.String(), d.Sha())
}

func (d *HerokuDeploy) SourceBlob() (*url.URL, error) {
	repoName := d.RepoName()

	url, err := d.github.GetArchiveLink(repoName.Owner(), repoName.Repo(), d.Sha())
	return url, err
}

func (d *HerokuDeploy) App() string {
	return d.RepoName().Repo()
}
