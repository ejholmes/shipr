package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	base            = "https://%s.slack.com"
	incomingWebhook = "/services/hooks/incoming-webhook?token=%s"
)

var (
	// DefaultTransport is a default Transport instance.
	DefaultTransport = &Transport{}

	// DefaultClient is the default http client to use.
	DefaultClient = &http.Client{
		Transport: DefaultTransport,
	}
)

// Client is an interface for our Slack client.
type Client interface {
	Notify(*Payload) error
}

// Field represents a field inside an Attachment.
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

// Attachment represents an incoming webhook attachment.
type Attachment struct {
	Fallback string   `json:"fallback"`
	Text     string   `json:"text"`
	PreText  string   `json:"pretext"`
	Color    string   `json:"color"`
	Fields   []*Field `json:"fields,omitempty"`
}

// Payload represents an incoming webhook payload.
type Payload struct {
	Username    string        `json:"username"`
	Text        string        `json:"text"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

// client is an implementation of the Client interface.
type client struct {
	client              *http.Client
	organization, token string
}

// New takes an incoming-webhook token and generates a new Client.
func New(org, token string, c *http.Client) Client {
	if c == nil {
		c = DefaultClient
	}
	return &client{c, org, token}
}

// Notify sends a slack notification.
func (c *client) Notify(p *Payload) error {
	j, err := json.Marshal(p)
	if err != nil {
		return err
	}

	req, err := c.NewRequest("POST", fmt.Sprintf(incomingWebhook, c.token), j)
	if err != nil {
		return err
	}

	_, err = c.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *client) NewRequest(method, path string, body []byte) (*http.Request, error) {
	fullPath := fmt.Sprintf(base, c.organization) + path
	return http.NewRequest(method, fullPath, bytes.NewReader(body))
}
