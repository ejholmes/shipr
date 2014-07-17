package slack

import "net/http"

// Transport implements the http.RoundTripper interface.
type Transport struct {
	Transport http.RoundTripper
}

// RoundTrip round trips the http.Request.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.Transport == nil {
		t.Transport = http.DefaultTransport
	}

	return t.Transport.RoundTrip(req)
}
