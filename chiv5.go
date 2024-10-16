package routetest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// ChiV5Reporter tracks visited routes of chi router.
// It is used to check if all routes have been tested.
type ChiV5Reporter struct {
	t *tracker
}

// NewChiV5Reporter creates a new chi reporter.
func NewChiV5Reporter() *ChiV5Reporter {
	return &ChiV5Reporter{t: newTracker()}
}

// Register registers chi router for testing.
func (rep *ChiV5Reporter) Register(router chi.Router) {
	for _, route := range router.Routes() {
		for method := range route.Handlers {
			rep.t.track(method, route.Pattern)
		}
	}
}

// Visitor is a middleware that tracks visited routes.
func (rep *ChiV5Reporter) Visitor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := chi.RouteContext(r.Context())
		rep.t.visit(ctx.RouteMethod, ctx.RoutePattern())
		next.ServeHTTP(w, r)
	})
}

// Report returns a report of visited routes.
func (rep *ChiV5Reporter) Report() Report {
	return newReport(rep.t)
}
