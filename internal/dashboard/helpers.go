package dashboard

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func CPUAt(cpus []CPUBar, idx int) CPUBar {
	if idx >= 0 && idx < len(cpus) {
		return cpus[idx]
	}
	return CPUBar{Label: fmt.Sprintf("%d", idx)}
}

func InnerBarWidth(totalWidth int, prefix, suffix string) int {
	width := totalWidth - lipgloss.Width(prefix) - lipgloss.Width(suffix) - 2
	if width < 4 {
		return 4
	}
	return width
}

func ClampInt(v, minV, maxV int) int {
	if v < minV {
		return minV
	}
	if v > maxV {
		return maxV
	}
	return v
}

func HumanBytes(v uint64) string {
	const unit = 1024
	if v < unit {
		return fmt.Sprintf("%dB", v)
	}
	div, exp := uint64(unit), 0
	for n := v / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f%c", float64(v)/float64(div), "KMGTPE"[exp])
}

func AverageCPU(cpus []CPUBar) float64 {
	if len(cpus) == 0 {
		return 0
	}
	total := 0.0
	for _, cpu := range cpus {
		total += cpu.Percent
	}
	return total / float64(len(cpus))
}

func JoinColumns(left, right string) string {
	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

func JoinCards(cards ...string) string {
	return lipgloss.JoinHorizontal(lipgloss.Top, cards...)
}

func JoinRows(rows ...string) string {
	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

func repeatBar(n int) string {
	return strings.Repeat("|", n)
}

func join(parts []string) string {
	return strings.Join(parts, "")
}
