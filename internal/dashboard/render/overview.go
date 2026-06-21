package render

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/VincentZyu233/htop-go/internal/dashboard"
)

func Overview(width int, data dashboard.Data, showTask bool, showLoad bool) string {
	topCPU := dashboard.AverageCPU(data.CPUs)
	colWidth := overviewColumnWidth(width)

	topRowContent := []string{
		renderOverviewBar("CPU", topCPU, fmt.Sprintf("%.1f%%", topCPU)),
		renderOverviewMemory(data.Memory),
		renderOverviewMemory(data.Swap),
	}

	bottomRowContent := make([]string, 0, 2)

	topRow := renderOverviewRow(topRowContent, colWidth)
	if showTask {
		bottomRowContent = append(bottomRowContent, renderOverviewTasks(data))
	}
	if showLoad {
		bottomRowContent = append(bottomRowContent, renderOverviewLoad(data))
	}
	if len(bottomRowContent) == 0 {
		return topRow
	}
	bottomRow := renderOverviewRow(bottomRowContent, overviewBottomColumnWidth(width, len(bottomRowContent)))
	return dashboard.JoinRows(topRow, bottomRow)
}

func renderOverviewBar(label string, percent float64, value string) string {
	width := 18
	filled := dashboard.ClampInt(int(percent/100*float64(width)), 0, width)
	bar := dashboard.GreenStyle().Render(strings.Repeat("█", filled)) + dashboard.MutedBarStyle().Render(strings.Repeat("░", width-filled))
	return fmt.Sprintf("%s\n%s\n%s", dashboard.LabelStyle().Render(label), bar, dashboard.ValueStyle().Render(value))
}

func renderOverviewMemory(mem dashboard.MemoryBar) string {
	return renderOverviewBar(mem.Label, memoryPercent(mem), mem.Display)
}

func renderOverviewTasks(data dashboard.Data) string {
	return strings.Join([]string{
		dashboard.LabelStyle().Render("Tasks"),
		fmt.Sprintf("all  %s", dashboard.GreenStyle().Render(fmt.Sprintf("%d", data.Tasks.Tasks))),
		fmt.Sprintf("run  %s", dashboard.GreenStyle().Render(fmt.Sprintf("%d", data.Tasks.Running))),
		fmt.Sprintf("thr  %s", dashboard.CyanStyle().Render(fmt.Sprintf("%d", data.Tasks.Threads))),
	}, "\n")
}

func renderOverviewLoad(data dashboard.Data) string {
	return strings.Join([]string{
		dashboard.LabelStyle().Render("Load"),
		dashboard.CyanStyle().Render(fmt.Sprintf("%.2f", data.Load[0])),
		dashboard.CyanStyle().Render(fmt.Sprintf("%.2f", data.Load[1])),
		dashboard.CyanStyle().Render(fmt.Sprintf("%.2f", data.Load[2])),
	}, "\n")
}

func renderOverviewRow(contents []string, colWidth int) string {
	height := 0
	for _, content := range contents {
		cardHeight := lipgloss.Height(content)
		if cardHeight > height {
			height = cardHeight
		}
	}
	if height < 2 {
		height = 2
	}

	cards := make([]string, 0, len(contents))
	for _, content := range contents {
		cards = append(cards, dashboard.CardStyle().
			Width(colWidth).
			Height(height).
			Render(content))
	}
	return dashboard.JoinCards(cards...)
}

func overviewColumnWidth(width int) int {
	colWidth := (width - 2) / 3
	if colWidth < 18 {
		return 18
	}
	return colWidth
}

func overviewBottomColumnWidth(width int, columns int) int {
	if columns <= 0 {
		return width
	}
	colWidth := (width - 1) / columns
	if colWidth < 18 {
		return 18
	}
	return colWidth
}

func memoryPercent(mem dashboard.MemoryBar) float64 {
	total := 0.0
	for _, seg := range mem.Segments {
		total += seg.Percent
	}
	if total < 0 {
		return 0
	}
	if total > 100 {
		return 100
	}
	return total
}
