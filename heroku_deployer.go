package main

import "github.com/cyberdelia/heroku-go/v3"

type HerokuDeployer struct {
	Client *heroku.Service
}

func NewHerokuDeployer(username, password string) *HerokuDeployer {
	client := heroku.NewService(heroku.DefaultClient)
	return &HerokuDeployer{client}
}

// Methods to implement the Deployer interface.

// Deploy deploys a job to Heroku.
func (h *HerokuDeployer) Deploy(d Deployable) error {
	_ = &HerokuDeploy{d}

	// Create Build

	// Poll Build Result

	return nil
}

// HerokuDeploy manages the lifecycle of a deployment via heroku builds.
type HerokuDeploy struct {
	Deployable
}

// AppName returns the name of the app for this github repo/environment combo.
func (d *HerokuDeploy) AppName() string {
	repoName := d.RepoName()

	switch d.Environment() {
	case "production":
		return repoName.Repo()
	default:
		return repoName.Repo() + "-" + d.Environment()
	}
}

// ArchiveLink returns the URL to download the source for this Sha.
func (d *HerokuDeploy) ArchiveLink() string {
	return ""
}
