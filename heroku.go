package shipr

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/ejholmes/heroku-go/v3"
)

// HerokuClient is an interface that defines the heroku client that we need.
type HerokuClient interface {
	BuildCreate(app, url, version string) (*heroku.Build, error)
}

// herokuClient is an implementation of the HerokuClient interface.
type herokuClient struct {
	heroku *heroku.Service
}

// newHerokuClient returns a new HerokuClient that is configured to authenticate
// with heroku via an oauth token.
func newHerokuClient(token string) HerokuClient {
	t := &oauth.Transport{
		Token:     &oauth.Token{AccessToken: token},
		Transport: heroku.DefaultTransport,
	}

	return &herokuClient{heroku.NewService(t.Client())}
}

// BuildCreate creates a build and returns it.
func (c *herokuClient) BuildCreate(app, url, version string) (*heroku.Build, error) {
	return c.heroku.BuildCreate(app, heroku.BuildCreateOpts{
		SourceBlob: struct {
			URL     *string `json:"url,omitempty"`
			Version *string `json:"version,omitempty"`
		}{&url, &version},
	})
}
