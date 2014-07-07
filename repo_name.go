package main

import "strings"

// RepoName is value object that represents the <owner/repo> format.
type RepoName string

// Owner returns the owner part of the repo name.
func (n RepoName) Owner() string {
	return n.parts()[0]
}

// Repo returns the repo part of the repo name.
func (n RepoName) Repo() string {
	return n.parts()[1]
}

// Parts returns the owner/repo parts.
func (n RepoName) parts() []string {
	return strings.SplitN(string(n), "/", 2)
}
