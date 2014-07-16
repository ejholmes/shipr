package github

import (
	"github.com/ejholmes/go-github/github"
	"github.com/remind101/shipr/util"
)

type notification struct {
	*github.DeploymentStatus
	*description
}

func newNotification(d *github.DeploymentStatus) *notification {
	return &notification{d, &description{d.Deployment}}
}

func (n *notification) State() string {
	return util.SafeString(n.DeploymentStatus.State)
}
