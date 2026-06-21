package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"

	"github.com/VincentZyu233/htop-go/internal/metrics"
)

var (
	tableHeaderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#94A3B8"))
	selectedRowStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#020617")).
				Background(lipgloss.Color("#F59E0B"))
)

func renderProcessTable(items []metrics.ProcessInfo, cursor, width, height int) string {
	if width < 40 {
		width = 40
	}

	rows := []string{
		tableHeaderStyle.Render(fmt.Sprintf("%-6s %-24s %-8s %-8s %-10s %-8s %-10s %s", "PID", "NAME", "CPU%", "MEM%", "RSS", "THR", "STATE", "CMD")),
	}

	if height < 5 {
		height = 5
	}
	maxRows := height - 1
	start := 0
	if cursor >= maxRows {
		start = cursor - maxRows + 1
	}
	end := start + maxRows
	if end > len(items) {
		end = len(items)
	}

	for idx := start; idx < end; idx++ {
		p := items[idx]
		line := fmt.Sprintf(
			"%-6d %-24s %-8.1f %-8.1f %-10s %-8d %-10s %s",
			p.PID,
			clip(p.Name, 24),
			p.CPUPercent,
			p.MemPercent,
			metrics.HumanBytes(p.RSS),
			p.Threads,
			clip(p.Status, 10),
			clip(p.Command, max(10, width-84)),
		)
		if idx == cursor {
			line = selectedRowStyle.Render(line)
		}
		rows = append(rows, line)
	}

	out := strings.Join(rows, "\n")
	currentHeight := lipgloss.Height(out)
	if currentHeight < height {
		out += strings.Repeat("\n", height-currentHeight)
	}
	return out
}

func clip(v string, width int) string {
	r := []rune(v)
	if len(r) <= width {
		return v
	}
	if width <= 1 {
		return string(r[:width])
	}
	return string(r[:width-1]) + "…"
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
