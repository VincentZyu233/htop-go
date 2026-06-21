package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/VincentZyu233/htop-go/internal/metrics"
)

const spinnerInterval = 120 * time.Millisecond

func fetchData(showTask bool, showLoad bool) tea.Cmd {
	return func() tea.Msg {
		snap := metrics.Collect(metrics.SortByCPU, "", 200, showTask)
		return refreshMsg{
			dashboard: metrics.DashboardDataFromSnapshot(snap, showTask, showLoad),
			snapshot:  snap,
		}
	}
}

func fetchDashboardOnly(showTask bool, showLoad bool) tea.Cmd {
	return func() tea.Msg {
		snap := metrics.CollectDashboardSnapshot(showTask)
		return refreshMsg{
			dashboard: metrics.DashboardDataFromSnapshot(snap, showTask, showLoad),
			snapshot:  snap,
		}
	}
}

func tick(interval time.Duration) tea.Cmd {
	return tea.Tick(interval, func(t time.Time) tea.Msg { return t })
}

func spinTick() tea.Cmd {
	return tea.Tick(spinnerInterval, func(time.Time) tea.Msg { return spinnerMsg{} })
}
