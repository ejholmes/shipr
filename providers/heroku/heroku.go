package heroku

import (
	"net/url"
	"time"

	"github.com/remind101/shipr"
	"github.com/remind101/shipr/github"
	"github.com/remind101/shipr/providers/heroku/heroku"
)

// Provider is an implementation of the shipr.Provider interface for deploying
// to Heroku using the Platform API: https://devcenter.heroku.com/articles/platform-api-reference#build
type Provider struct {
	github github.Client
	heroku heroku.Client
}

// newHerokuProvider builds a new herokuProvider and returns it.
func New(g github.Client, token string) shipr.Provider {
	h := heroku.New(token)
	return &Provider{g, h}
}

// Provider implements the Provider interface. Builds a new herokuDeploy and runs it.
func (h *Provider) Deploy(d shipr.Deployment) error {
	return newDeploy(h, d).run()
}

// deploy wraps a shipr.Deployment for managing the Heroku build process.
type deploy struct {
	*Provider
	shipr.Deployment
}

// newHerokuDeploy builds a new herokuDeploy and returns it.
func newDeploy(h *Provider, d shipr.Deployment) *deploy {
	return &deploy{h, d}
}

// run runs the build process.
func (d *deploy) run() error {
	b, err := d.createBuild()
	if err != nil {
		return err
	}

	go d.poll(b)

	return nil
}

// poll polls the build output.
func (d *deploy) poll(b *heroku.Build) {
	idx := 0
	throttle := time.Tick(1 * time.Second)

	for {
		<-throttle

		b, err := d.heroku.BuildResultInfo(d.app(), b.ID)
		if err != nil {
			continue
		}

		lines := newLines(b, idx)
		for _, l := range lines {
			d.AddLine(l.Line, time.Now())
		}
		idx += len(lines)

		if b.Build.Status == "succeeded" || b.Build.Status == "failed" {
			d.SetExitCode(int(b.ExitCode))
			break
		}
	}
}

// createBuild creates the Heroku build.
func (d *deploy) createBuild() (*heroku.Build, error) {
	source, err := d.sourceBlob()
	if err != nil {
		return nil, err
	}

	return d.heroku.BuildCreate(d.app(), source.String(), d.Sha())
}

// sourceBlob returns the archive link where the source can be downloaded by the heroku build.
func (d *deploy) sourceBlob() (*url.URL, error) {
	repoName := d.RepoName()
	return d.github.GetArchiveLink(repoName.Owner(), repoName.Repo(), d.Sha())
}

// app returns the name of the app.
func (d *deploy) app() string {
	return d.RepoName().Repo()
}

type logLine struct {
	Line   string
	Stream string
}

// newLines returns log lines after the provided index.
func newLines(b *heroku.BuildResult, idx int) []*logLine {
	raw := b.Lines[idx:len(b.Lines)]
	lines := make([]*logLine, len(raw))
	for i, l := range raw {
		lines[i] = &logLine{Line: l.Line, Stream: l.Stream}
	}
	return lines
}
