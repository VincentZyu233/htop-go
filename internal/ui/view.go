package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/VincentZyu233/htop-go/internal/dashboard"
	drender "github.com/VincentZyu233/htop-go/internal/dashboard/render"
)

var (
	pageStyle = lipgloss.NewStyle().Padding(1, 2)
	headStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F8FAFC")).
			Background(lipgloss.Color("#0F766E")).
			Bold(true).
			Padding(0, 1)
	metaStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#94A3B8"))
	loadingStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#67E8F9")).
			Bold(true)
)

var spinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

func (m model) View() string {
	if m.width == 0 {
		return "loading..."
	}
	if !m.ready {
		return renderLoading(m.width, m.height, m.spinner)
	}

	header := renderHeader(m.version, m.style, m.top.Uptime)
	dashboardView := renderDashboard(m.width-4, m.top, m.style, m.showTask, m.showLoad)
	content := []string{header, dashboardView}

	if m.showTable {
		usedHeight := lipgloss.Height(header) + lipgloss.Height(dashboardView) + 1
		tableHeight := m.height - usedHeight - 4
		if tableHeight < 6 {
			tableHeight = 6
		}
		content = append(content,
			renderProcessTable(m.snapshot.TopProcesses, m.cursor, m.width-4, tableHeight),
			metaStyle.Render("q quit  up/down scroll"),
		)
	} else {
		content = append(content, metaStyle.Render("q quit  --table enables the lower process table"))
	}

	return pageStyle.Render(strings.Join(content, "\n"))
}

func renderHeader(version string, style dashboard.Style, uptime string) string {
	return headStyle.Render(" htop-go ") + " " + metaStyle.Render(fmt.Sprintf("v%s  style=%s  uptime=%s", version, style, uptime))
}

func renderDashboard(width int, data dashboard.Data, style dashboard.Style, showTask bool, showLoad bool) string {
	if width < 40 {
		width = 40
	}

	switch style {
	case dashboard.StyleOverview:
		return drender.Overview(width, data, showTask, showLoad)
	default:
		return drender.Classic(width, data, showTask, showLoad)
	}
}

func renderLoading(width, height, spinner int) string {
	line := loadingStyle.Render(spinnerFrames[spinner] + " Preparing dashboard...")
	hint := metaStyle.Render("first sample can take a moment on Windows")
	block := lipgloss.JoinVertical(lipgloss.Center, line, hint)
	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, block)
}
