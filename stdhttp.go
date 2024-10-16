package routetest

import (
	"net/http"
	"strings"
)

// StdHttpReporter tracks visited routes of std http router.
// It is used to check if all routes have been tested.
type StdHttpReporter struct {
	t *tracker
}

// NewStdHttpReporter creates a new std http reporter.
func NewStdHttpReporter() *StdHttpReporter {
	return &StdHttpReporter{t: newTracker()}
}

// Register registers std http router for testing.
func (rep *StdHttpReporter) Register(pattern string) {
	method, path, _ := strings.Cut(pattern, " ")
	rep.t.track(method, path)
}

// Visitor is a middleware that tracks visited routes.
func (rep *StdHttpReporter) Visitor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, path, _ := strings.Cut(r.Pattern, " ")
		rep.t.visit(r.Method, path)
		next.ServeHTTP(w, r)
	})
}

// Report returns a report of visited routes.
func (rep *StdHttpReporter) Report() Report {
	return newReport(rep.t)
}
