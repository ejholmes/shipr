package providers

import (
	"github.com/remind101/shipr"
	"github.com/remind101/shipr/github"
	"github.com/remind101/shipr/providers/heroku"
	"github.com/remind101/shipr/providers/null"
)

// NewNullProvider returns a new null.Provider.
func NewNullProvider() shipr.Provider {
	return &null.Provider{}
}

// NewHerokuProvider returns a new heroku.Provider.
func NewHerokuProvider(g github.Client, token string) shipr.Provider {
	return heroku.New(g, token)
}
