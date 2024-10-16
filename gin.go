package routetest

import (
	"github.com/gin-gonic/gin"
)

// GinReporter tracks visited routes of gin router.
// It is used to check if all routes have been tested.
type GinReporter struct {
	t *tracker
}

// NewGinReporter creates a new gin reporter.
func NewGinReporter() *GinReporter {
	return &GinReporter{t: newTracker()}
}

// Register registers gin router for testing.
func (rep *GinReporter) Register(engine *gin.Engine) {
	for _, route := range engine.Routes() {
		rep.t.track(route.Method, route.Path)
	}
}

// Visitor is a middleware that tracks visited routes.
func (rep *GinReporter) Visitor(ctx *gin.Context) {
	rep.t.visit(ctx.Request.Method, ctx.FullPath())
	ctx.Next()
}

// Report returns a report of visited routes.
func (rep *GinReporter) Report() Report {
	return newReport(rep.t)
}
