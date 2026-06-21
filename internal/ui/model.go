package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/VincentZyu233/htop-go/internal/dashboard"
	"github.com/VincentZyu233/htop-go/internal/metrics"
)

type model struct {
	width     int
	height    int
	version   string
	style     dashboard.Style
	showTable bool
	showTask  bool
	showLoad  bool
	interval  time.Duration
	ready     bool
	spinner   int
	snapshot  metrics.Snapshot
	top       dashboard.Data
	cursor    int
}

func NewModel(version string, style dashboard.Style, showTable bool, showTask bool, showLoad bool, interval time.Duration) tea.Model {
	if interval <= 0 {
		interval = 2 * time.Second
	}
	return model{
		version:   version,
		style:     style,
		showTable: showTable,
		showTask:  showTask,
		showLoad:  showLoad,
		interval:  interval,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.refreshCmd(), tick(m.interval), spinTick())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case refreshMsg:
		m.top = msg.dashboard
		m.snapshot = msg.snapshot
		m.ready = true
		if m.cursor >= len(m.snapshot.TopProcesses) && len(m.snapshot.TopProcesses) > 0 {
			m.cursor = len(m.snapshot.TopProcesses) - 1
		}
	case spinnerMsg:
		if !m.ready {
			m.spinner = (m.spinner + 1) % len(spinnerFrames)
			return m, spinTick()
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.showTable && m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.showTable && m.cursor < len(m.snapshot.TopProcesses)-1 {
				m.cursor++
			}
		}
	case time.Time:
		return m, tea.Batch(m.refreshCmd(), tick(m.interval))
	}
	return m, nil
}

func (m model) refreshCmd() tea.Cmd {
	if m.showTable {
		return fetchData(m.showTask, m.showLoad)
	}
	return fetchDashboardOnly(m.showTask, m.showLoad)
}
