package routetest

import (
	"testing"
)

func TestReport(t *testing.T) {
	rep := newReport(newTracker())
	rep.MissedRoutes = []string{"GET /"}
	want := `

Routes test report:
Total: 0, Tested: 0, Missed: 0

Missed routes:
  GET /
`
	if get := rep.String(); get != want {
		t.Errorf("Report.String() = %q, want %q", get, want)
	}
}
