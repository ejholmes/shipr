package slack

import "github.com/remind101/shipr"

// Notifier implements the shipr.Notifier interface for sending status update
// notifications to Slack.
type Notifier struct {
}

// NewNotifier returns a new Slack Notifier.
func NewNotifier(token string) *Notifier {
	return &Notifier{}
}

// Notify sends a message to Slack.
func (t *Notifier) Notify(n shipr.Notification) {
}
