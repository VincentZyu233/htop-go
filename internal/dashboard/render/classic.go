package render

import (
	"strings"

	"github.com/VincentZyu233/htop-go/internal/dashboard"
)

func Classic(width int, data dashboard.Data, showTask bool, showLoad bool) string {
	leftWidth := width / 2
	rightWidth := width - leftWidth

	leftBars, rightBars := splitCPUColumns(data.CPUs, leftWidth, rightWidth)

	leftBars = append(leftBars,
		dashboard.RenderMemoryBar(leftWidth, data.Memory),
		dashboard.RenderMemoryBar(leftWidth, data.Swap),
	)

	rightBars = append(rightBars, dashboard.RenderUptimeLine(data.Uptime))
	if showTask {
		rightBars = append(rightBars[:len(rightBars)-1], append([]string{dashboard.RenderTasksLine(data.Tasks)}, rightBars[len(rightBars)-1:]...)...)
	}
	if showLoad {
		insertAt := len(rightBars) - 1
		rightBars = append(rightBars[:insertAt], append([]string{dashboard.RenderLoadLine(data.Load)}, rightBars[insertAt:]...)...)
	}
	if len(rightBars) > 0 {
		rightBars = rightBars[:len(rightBars)-1]
	}

	left := dashboard.ColumnStyle().Width(leftWidth).Render(strings.Join(leftBars, "\n"))
	right := dashboard.ColumnStyle().Width(rightWidth).Render(strings.Join(rightBars, "\n"))
	return dashboard.JoinColumns(left, right)
}

func splitCPUColumns(cpus []dashboard.CPUBar, leftWidth, rightWidth int) ([]string, []string) {
	if len(cpus) == 0 {
		return []string{dashboard.RenderCPUBar(leftWidth, dashboard.CPUBar{Label: "0"})}, []string{}
	}

	mid := (len(cpus) + 1) / 2
	leftBars := make([]string, 0, mid)
	rightBars := make([]string, 0, len(cpus)-mid)

	for i, cpu := range cpus {
		line := dashboard.RenderCPUBar(widthForColumn(i < mid, leftWidth, rightWidth), cpu)
		if i < mid {
			leftBars = append(leftBars, line)
		} else {
			rightBars = append(rightBars, line)
		}
	}

	return leftBars, rightBars
}

func widthForColumn(left bool, leftWidth, rightWidth int) int {
	if left {
		return leftWidth
	}
	return rightWidth
}
