package routetest

import (
	"net/http"
	"testing"
)

func TestTracker(t *testing.T) {
	tr := newTracker()

	for _, method := range []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodHead,
		http.MethodConnect,
	} {
		tr.track(method, "/")
		tr.visit(method, "/")
	}

	if len(tr.routes) != 5 {
		t.Errorf("expected 8 routes, got %d", len(tr.routes))
	}

	for _, method := range []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
	} {
		if !tr.routes[method+" /"] {
			t.Errorf("route %q not visited", method+" /")
		}
	}
}
