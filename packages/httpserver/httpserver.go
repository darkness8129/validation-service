package httpserver

import "time"

// HTTPServer provides a logic for starting HTTP server.
// It reports any errors encountered while running the server through a channel,
// which can be accessed using the Notify method.
type HTTPServer interface {
	Start()
	Notify() <-chan error
	Router() interface{}
	Shutdown(timeout time.Duration) error
}
