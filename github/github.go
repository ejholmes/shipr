package github

import (
	"net/url"

	"code.google.com/p/goauth2/oauth"
	"github.com/ejholmes/go-github/github"
)

// Client is an interface that defines the operations we perform
// against the GitHub api.
type Client interface {
	GetArchiveLink(owner, repo, ref string) (*url.URL, error)
}

// client is an implementation of the GitHubClient interface.
type client struct {
	github *github.Client
}

// NewClient reeturns a Client that is configured to authenticate
// via an oauth token.
func NewClient(token string) Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}

	// Setup a client for talking to GitHub.
	return &client{github.NewClient(t.Client())}
}

// GetArchiveLink returns the url to download a tarball archive of the source for the specified ref
func (c *client) GetArchiveLink(owner, repo, ref string) (*url.URL, error) {
	url, _, err := c.github.Repositories.GetArchiveLink(owner, repo, github.Tarball, &github.RepositoryContentGetOptions{Ref: ref})
	return url, err
}
