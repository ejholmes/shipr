package main

import (
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/binding"
)

func NewApi() *martini.ClassicMartini {
	m := martini.Classic()
	m.Post("/v1/deploys", binding.Bind(DeployForm{}), DeployCreate)
	m.Get("/v1/deploys", DeployList)
	m.Get("/v1/deploys/:id", DeployInfo)
	return m
}

type DeployForm struct {
	Name      string `json:"name"`
	Ref       string `json:"ref"`
	Force     bool   `json:"force"`
	AutoMerge bool   `json:"auto_merge"`
	Payload   map[string]interface{}
}

func DeployCreate(form DeployForm) {
}

func DeployList() {
}

func DeployInfo() {
}
