package shipr

import (
	"fmt"
	"net/url"

	"code.google.com/p/goauth2/oauth"

	"github.com/ejholmes/go-github/github"
	"github.com/ejholmes/heroku-go/v3"
)

type Deployer interface {
	Deploy(Deployable) error
}

type HerokuDeployer struct {
	github *github.Client
	heroku *heroku.Service
}

func NewHerokuDeployer(gh *github.Client, herokuToken string) *HerokuDeployer {
	t := &oauth.Transport{
		Token:     &oauth.Token{AccessToken: herokuToken},
		Transport: heroku.DefaultTransport,
	}

	h := heroku.NewService(t.Client())
	return &HerokuDeployer{gh, h}
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

	url, version := source.String(), d.Sha()

	return d.heroku.BuildCreate(d.App(), heroku.BuildCreateOpts{
		SourceBlob: struct {
			URL     *string `json:"url,omitempty"`
			Version *string `json:"version,omitempty"`
		}{&url, &version},
	})
}

func (d *HerokuDeploy) SourceBlob() (*url.URL, error) {
	repoName := d.RepoName()

	url, _, err := d.github.Repositories.GetArchiveLink(
		repoName.Owner(),
		repoName.Repo(),
		github.Tarball,
		&github.RepositoryContentGetOptions{Ref: d.Sha()},
	)
	return url, err
}

func (d *HerokuDeploy) App() string {
	return d.RepoName().Repo()
}
