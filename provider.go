package shipr

import (
	"net/url"
	"time"

	"github.com/remind101/shipr/github"
	"github.com/remind101/shipr/heroku"
)

// Provider is an interface that can be implemented for deploying a Deployable to
// some platform.
type Provider interface {
	Deploy(Deployable) error
}

// herokuProvider is an implementation of the Provider interface for deploying
// to Heroku using the Platform API: https://devcenter.heroku.com/articles/platform-api-reference#build
type herokuProvider struct {
	github github.Client
	heroku heroku.Client
}

// newHerokuProvider builds a new herokuProvider and returns it.
func newHerokuProvider(g github.Client, h heroku.Client) Provider {
	return &herokuProvider{g, h}
}

// Provider implements the Provider interface. Builds a new herokuDeploy and runs it.
func (h *herokuProvider) Deploy(d Deployable) error {
	return newHerokuDeploy(h, d).run()
}

// herokuDeploy wraps a deployable for managing the Heroku build process.
type herokuDeploy struct {
	*herokuProvider
	Deployable
}

// newHerokuDeploy builds a new herokuDeploy and returns it.
func newHerokuDeploy(h *herokuProvider, d Deployable) *herokuDeploy {
	return &herokuDeploy{h, d}
}

// run runs the build process.
func (d *herokuDeploy) run() error {
	b, err := d.createBuild()
	if err != nil {
		return err
	}

	go d.poll(b)

	return nil
}

// polls polls the build output.
func (d *herokuDeploy) poll(b *heroku.Build) {
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
func (d *herokuDeploy) createBuild() (*heroku.Build, error) {
	source, err := d.sourceBlob()
	if err != nil {
		return nil, err
	}

	return d.heroku.BuildCreate(d.app(), source.String(), d.Sha())
}

// sourceBlob returns the archive link where the source can be downloaded by the heroku build.
func (d *herokuDeploy) sourceBlob() (*url.URL, error) {
	repoName := d.RepoName()
	return d.github.GetArchiveLink(repoName.Owner(), repoName.Repo(), d.Sha())
}

// app returns the name of the app.
func (d *herokuDeploy) app() string {
	return d.RepoName().Repo()
}

type herokuLogLine struct {
	Line   string
	Stream string
}

// newLines returns log lines after the provided index.
func newLines(b *heroku.BuildResult, idx int) []*herokuLogLine {
	raw := b.Lines[idx:len(b.Lines)]
	lines := make([]*herokuLogLine, len(raw))
	for i, l := range raw {
		lines[i] = &herokuLogLine{Line: l.Line, Stream: l.Stream}
	}
	return lines
}
