package heroku

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/ejholmes/heroku-go/v3"
)

// HerokuClient is an interface that defines the heroku client that we need.
type Client interface {
	BuildCreate(appId, url, version string) (*Build, error)
	BuildOutputStream(appId, buildId string) chan *BuildResultLine
}

// Build wraps heroku.Build.
type Build struct {
	*heroku.Build
}

// BuildResultLine represnts a log line from the build result.
type BuildResultLine struct {
	Line   string
	Stream string
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
func (c *client) BuildCreate(appId, url, version string) (*Build, error) {
	build, err := c.heroku.BuildCreate(appId, heroku.BuildCreateOpts{
		SourceBlob: struct {
			URL     *string `json:"url,omitempty"`
			Version *string `json:"version,omitempty"`
		}{&url, &version},
	})

	return &Build{build}, err
}

// BuildOutputStream returns a channel that streams the build output.
func (c *client) BuildOutputStream(appId, buildId string) chan *BuildResultLine {
	ch := make(chan *BuildResultLine)
	return ch
}
