package routetest

import (
	"bytes"
	"fmt"
)

// Report is a report of visited routes.
type Report struct {
	Total  int
	Tested int
	Missed int

	MissedRoutes []string
}

// newReport returns a report of visited routes.
func newReport(t *tracker) Report {
	report := Report{}

	for route, visited := range t.routes {
		report.Total++
		if visited {
			report.Tested++
		} else {
			report.Missed++
			report.MissedRoutes = append(report.MissedRoutes, route)
		}
	}

	return report
}

// Success returns true if all routes have been tested.
func (r Report) Success() bool {
	return r.Missed == 0
}

// String returns a string representation of the report.
func (r Report) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "\n\nRoutes test report:\n")
	fmt.Fprintf(&buf, "Total: %d, Tested: %d, Missed: %d\n",
		r.Total, r.Tested, r.Missed)

	if len(r.MissedRoutes) > 0 {
		buf.WriteString("\nMissed routes:\n")
		for _, route := range r.MissedRoutes {
			buf.WriteString("  " + route + "\n")
		}
	}

	return buf.String()
}
