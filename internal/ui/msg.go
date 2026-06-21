package ui

import (
	"github.com/VincentZyu233/htop-go/internal/dashboard"
	"github.com/VincentZyu233/htop-go/internal/metrics"
)

type refreshMsg struct {
	dashboard dashboard.Data
	snapshot  metrics.Snapshot
}

type spinnerMsg struct{}
