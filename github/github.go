package github

import (
	"net/http"
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

// New returns a Client that is configured to authenticate
// via an oauth token.
func New(token string) Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}

	return NewClient(t.Client())
}

// NewClient returns a Client.
func NewClient(c *http.Client) Client {
	return &client{github.NewClient(c)}
}

// GetArchiveLink returns the url to download a tarball archive of the source for the specified ref
func (c *client) GetArchiveLink(owner, repo, ref string) (*url.URL, error) {
	url, _, err := c.github.Repositories.GetArchiveLink(owner, repo, github.Tarball, &github.RepositoryContentGetOptions{Ref: ref})
	return url, err
}
