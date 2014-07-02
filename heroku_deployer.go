package main

import (
	"fmt"

	"github.com/cyberdelia/heroku-go/v3"
)

type HerokuDeployer struct {
	Client *heroku.Service
}

type HerokuDeploy struct {
	Deployable
}

func NewHerokuDeployer(username, password string) *HerokuDeployer {
	client := heroku.NewService(heroku.DefaultClient)
	return &HerokuDeployer{client}
}

// Methods to implement the Deployer interface.

// Deploy deploys a job to Heroku.
func (h *HerokuDeployer) Deploy(d Deployable) error {
	hd := &HerokuDeploy{d}
	fmt.Println(hd.AppName())
	//addons, err := h.Client.AddonList("r101-shipr", nil)
	//if err != nil {
	//panic(err)
	//}
	//fmt.Println(addons)
	return nil
}

func (hd *HerokuDeploy) AppName() string {
	return hd.RepoName()
}
