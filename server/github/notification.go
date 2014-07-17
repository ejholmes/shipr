package github

import (
	"net/url"

	"github.com/ejholmes/go-github/github"
	"github.com/remind101/shipr"
	"github.com/remind101/shipr/util"
)

type notification struct {
	*github.DeploymentStatus
	*description
}

func newNotification(d *github.DeploymentStatus) *notification {
	return &notification{d, &description{d.Deployment}}
}

func (n *notification) RepoName() shipr.RepoName {
	return shipr.RepoName("TODO")
}

func (n *notification) URL() *url.URL {
	u, err := url.Parse("http://www.google.com")
	if err != nil {
		panic(err)
	}
	return u
}

func (n *notification) User() string {
	return "TODO"
}

func (n *notification) State() string {
	return util.SafeString(n.DeploymentStatus.State)
}
