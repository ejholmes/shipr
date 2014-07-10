package shipr

import (
	"code.google.com/p/goauth2/oauth"
	"github.com/ejholmes/heroku-go/v3"
)

type HerokuClient struct {
	*heroku.Service
}

func NewHerokuClient(token string) *HerokuClient {
	t := &oauth.Transport{
		Token:     &oauth.Token{AccessToken: token},
		Transport: heroku.DefaultTransport,
	}

	return &HerokuClient{heroku.NewService(t.Client())}
}
