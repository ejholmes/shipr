package test

import (
	"net/url"

	"github.com/remind101/shipr"
)

// Notification is a fake implementation of the shipr.Notification interface.
type Notification struct {
	RawState string
}

func (n *Notification) Guid() int {
	return 1234
}

func (n *Notification) Sha() string {
	return "208db1b5d2bb89b0ff0b79cb7f702e21a750f3fc"
}

func (n *Notification) Ref() string {
	return "master"
}

func (n *Notification) RepoName() shipr.RepoName {
	return shipr.RepoName("remind101/r101-api")
}

func (n *Notification) Description() string {
	return "Deploying my repo"
}

func (n *Notification) Environment() string {
	return "staging"
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
