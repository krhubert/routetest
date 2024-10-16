package routetest

import (
	"net/http"
	"sync"
)

type tracker struct {
	routes map[string]bool
	mx     sync.Mutex
}

func newTracker() *tracker {
	return &tracker{
		routes: make(map[string]bool),
	}
}

func (t *tracker) track(method, path string) {
	if !isMethodTested(method) {
		return
	}
	t.mx.Lock()
	t.routes[method+" "+path] = false
	t.mx.Unlock()
}

func (t *tracker) visit(method, path string) {
	t.mx.Lock()
	if _, ok := t.routes[method+" "+path]; ok {
		t.routes[method+" "+path] = true
	}
	t.mx.Unlock()
}

func isMethodTested(method string) bool {
	return method == http.MethodGet ||
		method == http.MethodPost ||
		method == http.MethodPut ||
		method == http.MethodPatch ||
		method == http.MethodDelete
}
