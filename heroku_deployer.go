package main

import (
	"fmt"

	"github.com/bgentry/heroku-go"
)

type HerokuDeployer struct {
	Client heroku.Client
}

func NewHerokuDeployer(username, password string) *HerokuDeployer {
	client := heroku.Client{Username: username, Password: password}
	return &HerokuDeployer{client}
}

func (h *HerokuDeployer) Deploy(j *Job) error {
	addons, err := h.Client.AddonList("r101-shipr", nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(addons)
	return nil
}
