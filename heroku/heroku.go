package heroku

import (
	"net/http"

	"code.google.com/p/goauth2/oauth"
	"github.com/ejholmes/heroku-go/v3"
)

// Client is an interface that defines the heroku client that we need.
type Client interface {
	BuildCreate(appId, url, version string) (*Build, error)
	BuildResultInfo(appId, buildId string) (*BuildResult, error)
}

// Build wraps heroku.Build.
type Build struct {
	*heroku.Build
}

// BuildResult wraps heroku.BuildResult.
type BuildResult struct {
	*heroku.BuildResult
}

// client is an implementation of the HerokuClient interface.
type client struct {
	heroku *heroku.Service
}

// newClient builds a new client.
func NewClient(c *http.Client) Client {
	return &client{heroku.NewService(c)}
}

// New returns a new Client that is configured to authenticate
// with heroku via an oauth token.
func New(token string) Client {
	t := &oauth.Transport{
		Token:     &oauth.Token{AccessToken: token},
		Transport: heroku.DefaultTransport,
	}

	return NewClient(t.Client())
}

// BuildCreate creates a build and returns it.
func (c *client) BuildCreate(appId, url, version string) (*Build, error) {
	b, err := c.heroku.BuildCreate(appId, heroku.BuildCreateOpts{
		SourceBlob: struct {
			URL     *string `json:"url,omitempty"`
			Version *string `json:"version,omitempty"`
		}{&url, &version},
	})

	return &Build{b}, err
}

// BuildResultInfo gets info about a a Build.
func (c *client) BuildResultInfo(appId, buildId string) (*BuildResult, error) {
	b, err := c.heroku.BuildResultInfo(appId, buildId)
	return &BuildResult{b}, err
}
