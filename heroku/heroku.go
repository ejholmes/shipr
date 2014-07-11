package heroku

import (
	"net/http"
	"time"

	"code.google.com/p/goauth2/oauth"
	"github.com/ejholmes/heroku-go/v3"
)

// Client is an interface that defines the heroku client that we need.
type Client interface {
	BuildCreate(appId, url, version string) (*Build, error)
	BuildOutputStream(appId, buildId string, lines chan *BuildResultLine, status chan string)
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
	build, err := c.heroku.BuildCreate(appId, heroku.BuildCreateOpts{
		SourceBlob: struct {
			URL     *string `json:"url,omitempty"`
			Version *string `json:"version,omitempty"`
		}{&url, &version},
	})

	return &Build{build}, err
}

// BuildOutputStream returns a channel that streams the build output.
func (c *client) BuildOutputStream(appId, buildId string, lch chan *BuildResultLine, sch chan string) {
	idx, status := 0, ""
	throttle := time.Tick(1 * time.Second)

	go func() {
		for {
			<-throttle

			b, err := c.heroku.BuildResultInfo(appId, buildId)
			if err != nil {
				continue
			}

			lines := newBuildResultLines(b, idx)
			for _, l := range lines {
				lch <- l
			}
			idx += len(lines)

			if b.Build.Status != status {
				status = b.Build.Status
				sch <- status
			}
		}
	}()
}

// newBuildResultLines returns log lines after the provided index.
func newBuildResultLines(b *heroku.BuildResult, idx int) []*BuildResultLine {
	raw := b.Lines[idx:len(b.Lines)]
	lines := make([]*BuildResultLine, len(raw))
	for i, l := range raw {
		lines[i] = &BuildResultLine{Line: l.Line, Stream: l.Stream}
	}
	return lines
}
