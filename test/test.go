package test

import (
	"net/url"

	"github.com/remind101/shipr"
)

// description is a fake implementation of the shipr.Description interface.
type description struct {
}

func (d *description) Guid() int {
	return 1234
}

func (d *description) Sha() string {
	return "208db1b5d2bb89b0ff0b79cb7f702e21a750f3fc"
}

func (d *description) Ref() string {
	return "master"
}

func (d *description) RepoName() shipr.RepoName {
	return shipr.RepoName("remind101/r101-api")
}

func (d *description) Description() string {
	return "Deploying my repo"
}

func (d *description) Environment() string {
	return "staging"
}

// Notification is a fake implementation of the shipr.Notification interface.
type Notification struct {
	description
	RawState string
}

func (n *Notification) URL() *url.URL {
	u, _ := url.Parse("http://shipr.example.org/1234")
	return u
}

func (n *Notification) User() string {
	return "ejholmes"
}

func (n *Notification) State() string {
	if n.RawState != "" {
		return n.RawState
	}
	return "pending"
}
