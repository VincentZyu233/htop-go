package dashboard

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	labelStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#D1D5DB"))
	valueStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#D1D5DB"))
	greenStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#22C55E"))
	redStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#EF4444"))
	blueStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#60A5FA"))
	cyanStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#67E8F9"))
	mutedBar    = lipgloss.NewStyle().Foreground(lipgloss.Color("#374151"))
	columnStyle = lipgloss.NewStyle().Padding(0, 1)
	cardStyle   = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#334155")).
			Padding(0, 1)
)

func RenderCPUBar(width int, cpu CPUBar) string {
	label := cpu.Label
	if label == "" {
		label = "?"
	}

	suffix := fmt.Sprintf("%.1f%%]", cpu.Percent)
	prefix := fmt.Sprintf("%s[", label)
	barWidth := InnerBarWidth(width, prefix, suffix)
	used := ClampInt(int(cpu.Percent/100*float64(barWidth)), 0, barWidth)

	greenCount := used
	redCount := 0
	if cpu.Percent > 75 {
		redCount = ClampInt(int((cpu.Percent-75)/25*float64(used)), 0, used)
		greenCount = used - redCount
	}

	bar := greenStyle.Render(repeatBar(greenCount)) +
		redStyle.Render(repeatBar(redCount)) +
		mutedBar.Render(repeatBar(barWidth-used))

	return labelStyle.Render(prefix) + bar + valueStyle.Render(suffix)
}

func RenderMemoryBar(width int, mem MemoryBar) string {
	label := mem.Label
	if label == "" {
		label = "Mem"
	}
	display := mem.Display
	if display == "" {
		display = fmt.Sprintf("%s/%s", HumanBytes(mem.Used), HumanBytes(mem.Total))
	}

	suffix := fmt.Sprintf("%s]", display)
	prefix := fmt.Sprintf("%s[", label)
	barWidth := InnerBarWidth(width, prefix, suffix)

	parts := make([]string, 0, len(mem.Segments)+1)
	filled := 0
	for _, seg := range mem.Segments {
		count := ClampInt(int(seg.Percent/100*float64(barWidth)), 0, barWidth-filled)
		if count <= 0 {
			continue
		}
		parts = append(parts, lipgloss.NewStyle().Foreground(seg.Color).Render(repeatBar(count)))
		filled += count
		if filled >= barWidth {
			break
		}
	}
	if filled < barWidth {
		parts = append(parts, mutedBar.Render(repeatBar(barWidth-filled)))
	}

	return labelStyle.Render(prefix) + join(parts) + valueStyle.Render(suffix)
}

func RenderTasksLine(tasks TasksInfo) string {
	return labelStyle.Render("Tasks: ") +
		greenStyle.Render(fmt.Sprintf("%d", tasks.Tasks)) +
		labelStyle.Render(", ") +
		cyanStyle.Render(fmt.Sprintf("%d", tasks.Threads)) +
		labelStyle.Render(" thr; ") +
		greenStyle.Render(fmt.Sprintf("%d", tasks.Running)) +
		blueStyle.Render(" running")
}

func RenderLoadLine(load [3]float64) string {
	return labelStyle.Render("Load average: ") +
		cyanStyle.Render(fmt.Sprintf("%.2f %.2f %.2f", load[0], load[1], load[2]))
}

func RenderUptimeLine(uptime string) string {
	return labelStyle.Render("Uptime: ") + valueStyle.Render(uptime)
}

func LabelStyle() lipgloss.Style {
	return labelStyle
}

func ValueStyle() lipgloss.Style {
	return valueStyle
}

func GreenStyle() lipgloss.Style {
	return greenStyle
}

func CyanStyle() lipgloss.Style {
	return cyanStyle
}

func MutedBarStyle() lipgloss.Style {
	return mutedBar
}

func ColumnStyle() lipgloss.Style {
	return columnStyle
}

func CardStyle() lipgloss.Style {
	return cardStyle
}
