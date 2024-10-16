package routetest

import (
	"github.com/labstack/echo/v4"
)

// EchoV4Reporter tracks visited routes of echo router.
// It is used to check if all routes have been tested.
type EchoV4Reporter struct {
	t *tracker
}

// NewReporter creates a new echo reporter.
func NewEchoV4Reporter() *EchoV4Reporter {
	return &EchoV4Reporter{t: newTracker()}
}

// Register registers echo router for testing.
func (rep *EchoV4Reporter) Register(router *echo.Echo) {
	for _, route := range router.Routes() {
		rep.t.track(route.Method, route.Path)
	}
}

// Visitor is a middleware that tracks visited routes.
func (rep *EchoV4Reporter) Visitor(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		rep.t.visit(c.Request().Method, c.Path())
		return next(c)
	}
}

// Report returns a report of visited routes.
func (rep *EchoV4Reporter) Report() Report {
	return newReport(rep.t)
}
