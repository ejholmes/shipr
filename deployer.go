package shipr

import (
	"fmt"
	"net/url"

	"github.com/ejholmes/go-github/github"
)

type Deployer interface {
	Deploy(Deployable) error
}

type HerokuDeployer struct {
	GitHub *github.Client
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
	url, err := d.SourceBlob()
	if err != nil {
		return err
	}
	fmt.Println(url)
	return nil
}

func (d *HerokuDeploy) SourceBlob() (*url.URL, error) {
	repoName := d.RepoName()

	url, _, err := d.GitHub.Repositories.GetArchiveLink(
		repoName.Owner(),
		repoName.Repo(),
		github.Tarball,
		&github.RepositoryContentGetOptions{Ref: d.Sha()},
	)
	if err != nil {
		return nil, err
	}
	return url, nil
}
