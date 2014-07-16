package null

import (
	"log"
	"os"

	"github.com/remind101/shipr"
)

// Provider is an implementation of the shipr.Provider interface that
// simply logs to stdout.
type Provider struct {
	*log.Logger
}

// Deploy logs the shipr.Deployment to Stdout.
func (p *Provider) Deploy(d shipr.Deployment) error {
	p.logger().Printf(`guid=%v repo="%s"`, d.Guid(), d.RepoName())
	return nil
}

func (p *Provider) logger() *log.Logger {
	if p.Logger == nil {
		p.Logger = log.New(os.Stdout, "[provider] ", 0)
	}
	return p.Logger
}
