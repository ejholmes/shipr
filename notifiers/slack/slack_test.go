package slack

import (
	"testing"

	"github.com/remind101/shipr/notifiers/slack/slack"
	"github.com/remind101/shipr/test"
)

// nullClient is an implementation of the Client interface that stores the payload for
// inspection.
type nullClient struct {
	payload *slack.Payload
}

func (c *nullClient) Notify(p *slack.Payload) error {
	c.payload = p
	return nil
}

func Test_Notify(t *testing.T) {
	client := &nullClient{}
	notifier := &Notifier{client: client}
	notification := &test.Notification{}

	tests := []struct {
		state, expected string
	}{
		{"pending", "ejholmes is <http://shipr.example.org/1234|deploying> r101-api@208db1b5d2bb89b0ff0b79cb7f702e21a750f3fc to staging"},
		{"success", "ejholmes <http://shipr.example.org/1234|deployed> r101-api@208db1b5d2bb89b0ff0b79cb7f702e21a750f3fc to staging"},
		{"failure", "ejholmes failed to <http://shipr.example.org/1234|deploy> r101-api@208db1b5d2bb89b0ff0b79cb7f702e21a750f3fc to staging"},
		{"error", "ejholmes failed to <http://shipr.example.org/1234|deploy> r101-api@208db1b5d2bb89b0ff0b79cb7f702e21a750f3fc to staging"},
	}

	for i, test := range tests {
		notification.RawState = test.state

		err := notifier.Notify(notification)
		if err != nil {
			t.Error(err)
		}

		text := client.payload.Attachments[0].Text
		if text != test.expected {
			t.Errorf("%v: Want %v; Got %v", i, test.expected, text)
		}
	}
}
