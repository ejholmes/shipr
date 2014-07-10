package heroku

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/ejholmes/heroku-go/v3"
)

// HerokuClient is an interface that defines the heroku client that we need.
type Client interface {
	BuildCreate(app, url, version string) (*Build, error)
}

// Build wraps heroku.Build.
type Build struct {
	*heroku.Build
}

// herokuClient is an implementation of the HerokuClient interface.
type client struct {
	heroku *heroku.Service
}

// newHerokuClient returns a new HerokuClient that is configured to authenticate
// with heroku via an oauth token.
func NewClient(token string) Client {
	t := &oauth.Transport{
		Token:     &oauth.Token{AccessToken: token},
		Transport: heroku.DefaultTransport,
	}

	return &client{heroku.NewService(t.Client())}
}

// BuildCreate creates a build and returns it.
func (c *client) BuildCreate(app, url, version string) (*Build, error) {
	build, err := c.heroku.BuildCreate(app, heroku.BuildCreateOpts{
		SourceBlob: struct {
			URL     *string `json:"url,omitempty"`
			Version *string `json:"version,omitempty"`
		}{&url, &version},
	})

	return &Build{build}, err
}
