package ui

import (
	"github.com/VincentZyuApps/htop-go/internal/dashboard"
	"github.com/VincentZyuApps/htop-go/internal/metrics"
)

type refreshMsg struct {
	dashboard dashboard.Data
	snapshot  metrics.Snapshot
}

type spinnerMsg struct{}
