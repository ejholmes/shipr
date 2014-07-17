package slack

import (
	"bytes"
	"text/template"

	"github.com/remind101/shipr"
	"github.com/remind101/shipr/notifiers/slack/slack"
)

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

	attachment
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
	a, err := n.Attachment(n)
	if err != nil {
		return nil, err
	}

	return &slack.Payload{Attachments: []*slack.Attachment{a}}, nil
}

// attachment represents a color/message combination.
type attachment struct {
	color    string
	template string
}

// Attachment renders the template using the data provided and returns a slack.Attachment.
func (a *attachment) Attachment(data interface{}) (*slack.Attachment, error) {
	msg, err := a.message(data)
	if err != nil {
		return nil, err
	}

	return &slack.Attachment{
		Text:     msg,
		Fallback: msg,
		Color:    a.color,
	}, nil
}

// message returns rendered template.
func (a *attachment) message(data interface{}) (string, error) {
	var b bytes.Buffer

	t, err := template.New("message").Parse(a.template)
	if err != nil {
		return "", err
	}

	err = t.Execute(&b, data)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
