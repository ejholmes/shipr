package slack

import (
	"bytes"
	"text/template"

	"github.com/remind101/shipr"
	"github.com/remind101/shipr/notifiers/slack/slack"
)

// attachment represents a color/message combination.
type attachment struct {
	color    string
	template string
}

var templates = map[string]attachment{
	"pending": attachment{"#ff0", "{{.User}} is <{{.URL}}|deploying> {{.Repo}}@{{.Sha}} to {{.Environment}}"},
	"success": attachment{"#0f0", "{{.User}} <{{.URL}}|deployed> {{.Repo}}@{{.Sha}} to {{.Environment}}"},
	"failure": attachment{"#f00", "{{.User}} failed to <{{.URL}}|deploy> {{.Repo}}@{{.Sha}} to {{.Environment}}"},
	"error":   attachment{"#f00", "{{.User}} failed to <{{.URL}}|deploy> {{.Repo}}@{{.Sha}} to {{.Environment}}"},
}

// Notifier implements the shipr.Notifier interface for sending status update
// notifications to Slack.
type Notifier struct {
	client slack.Client
}

// NewNotifier returns a new Slack Notifier.
func NewNotifier(org, token string) *Notifier {
	c := slack.New(org, token, nil)
	return &Notifier{c}
}

// Notify sends a message to Slack.
func (t *Notifier) Notify(n shipr.Notification) error {
	p, err := newNotification(n).Payload()
	if err != nil {
		return err
	}
	return t.client.Notify(p)
}

type notification struct {
	shipr.Notification

	User        string
	URL         string
	Repo        string
	Sha         string
	Environment string

	attachment attachment
}

func newNotification(n shipr.Notification) *notification {
	return &notification{
		User:        n.User(),
		Repo:        n.RepoName().Repo(),
		URL:         n.URL().String(),
		Sha:         n.Sha(),
		Environment: n.Environment(),
		attachment:  templates[n.State()],
	}
}

// Payload returns a slack.Payload for this notification.
func (n *notification) Payload() (*slack.Payload, error) {
	msg, err := n.message()
	if err != nil {
		return nil, err
	}

	color := n.attachment.color
	a := &slack.Attachment{
		Text:     msg,
		Fallback: msg,
		Color:    color,
	}

	return &slack.Payload{Attachments: []*slack.Attachment{a}}, nil
}

// message returns the rendered message.
func (n *notification) message() (string, error) {
	var b bytes.Buffer

	t, err := template.New("message").Parse(n.attachment.template)
	if err != nil {
		return "", err
	}

	err = t.Execute(&b, n)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
