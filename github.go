package shipr

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/ejholmes/go-github/github"
)

type GitHubClient struct {
	*github.Client
}

func NewGitHubClient(token string) *GitHubClient {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}

	// Setup a client for talking to GitHub.
	return &GitHubClient{github.NewClient(t.Client())}
}
