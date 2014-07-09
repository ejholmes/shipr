package github

import (
	"net/http"

	"github.com/google/go-github/github"
)

type Client struct {
	*github.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{github.NewClient(httpClient)}
}
