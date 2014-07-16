package null

import (
	"log"
	"os"

	"github.com/remind101/shipr"
)

// Notifier is an implementation of the shipr.Notifier interface that logs the
// notification to Stdout.
type Notifier struct {
	*log.Logger
}

// Notify logs the notification to Stdout.
func (notifier *Notifier) Notify(n shipr.Notification) error {
	notifier.logger().Println(n)
	return nil
}

func (notifier *Notifier) logger() *log.Logger {
	if notifier.Logger == nil {
		notifier.Logger = log.New(os.Stdout, "[notifier]", 0)
	}
	return notifier.Logger
}
