package shipr

import (
	"net/url"

	"code.google.com/p/goauth2/oauth"
	"github.com/ejholmes/go-github/github"
)

// GitHubClient is an interface that defines the operations we perform
// against the GitHub api.
type GitHubClient interface {
	GetArchiveLink(owner, repo, ref string) (*url.URL, error)
}

// gitHubClient is an implementation of the GitHubClient interface.
type gitHubClient struct {
	github *github.Client
}

// newGitHubClient reeturns a GitHubClient that is configured to authenticate
// via an oauth token.
func newGitHubClient(token string) GitHubClient {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}

	// Setup a client for talking to GitHub.
	return &gitHubClient{github.NewClient(t.Client())}
}

// GetArchiveLink returns the url to download a tarball archive of the source for the specified ref
func (c *gitHubClient) GetArchiveLink(owner, repo, ref string) (*url.URL, error) {
	url, _, err := c.github.Repositories.GetArchiveLink(owner, repo, github.Tarball, &github.RepositoryContentGetOptions{Ref: ref})
	return url, err
}
