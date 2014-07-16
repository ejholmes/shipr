package notifiers

import "github.com/remind101/shipr/notifiers/null"

// NewNullNotifier returns a new null.Notifier.
func NewNullNotifier() *null.Notifier {
	return &null.Notifier{}
}
