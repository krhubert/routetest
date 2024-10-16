package routetest

import (
	"net/http"
	"net/http/httptest"
	"slices"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGinReporter(t *testing.T) {
	rep := NewGinReporter()
	gin.SetMode(gin.TestMode)
	engine := gin.New()

	handler := func(c *gin.Context) {}

	engine.Use(rep.Visitor)
	engine.GET("/tested/:id", handler)
	engine.GET("/not-tested", handler)
	rep.Register(engine)

	ts := httptest.NewServer(engine)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/tested/1")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status code got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	report := rep.Report()
	if report.Success() {
		t.Fatal("report should not be success")
	}
	if report.Total != 2 {
		t.Fatalf("total got %d, want 2", report.Total)
	}
	if report.Tested != 1 {
		t.Fatalf("tested got %d, want 1", report.Tested)
	}
	if report.Missed != 1 {
		t.Fatalf("missed got %d, want 1", report.Missed)
	}

	missedRoutes := []string{"GET /not-tested"}
	if !slices.Equal(report.MissedRoutes, missedRoutes) {
		t.Fatalf("missed routes got %v, want %v", report.MissedRoutes, missedRoutes)
	}
}
