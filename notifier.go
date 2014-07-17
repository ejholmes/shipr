package shipr

import "net/url"

// Notification is an interface that is provided to Notifiers when there is a status
// update on a deploy.
type Notification interface {
	Description

	URL() *url.URL
	User() string
	State() string
}

// Notifier is an interface that notifiers can implement to forward notifications
// to an external system, like HipChat or Slack.
type Notifier interface {
	Notify(Notification) error
}
