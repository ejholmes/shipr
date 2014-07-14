package github

import (
	"github.com/ejholmes/go-github/github"
	"github.com/remind101/shipr"
	"github.com/remind101/shipr/util"
)

// description wraps a github.Deployment to implement the shipr.Description interface.
type description struct {
	*github.Deployment
}

func (d *description) Guid() int {
	return *d.Deployment.ID
}

func (d *description) RepoName() shipr.RepoName {
	return shipr.RepoName(util.SafeString(d.Deployment.Repository.FullName))
}

func (d *description) Sha() string {
	return util.SafeString(d.Deployment.Sha)
}

func (d *description) Ref() string {
	return util.SafeString(d.Deployment.Ref)
}

func (d *description) Environment() string {
	return util.SafeString(d.Deployment.Environment)
}

func (d *description) Description() string {
	return util.SafeString(d.Deployment.Description)
}
