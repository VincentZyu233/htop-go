package dashboard

import "github.com/charmbracelet/lipgloss"

type Data struct {
	CPUs   []CPUBar
	Memory MemoryBar
	Swap   MemoryBar
	Tasks  TasksInfo
	Load   [3]float64
	Uptime string
}

type Style string

const (
	StyleClassic  Style = "classic"
	StyleOverview Style = "overview"
)

type CPUBar struct {
	Label   string
	Percent float64
}

type MemoryBar struct {
	Label    string
	Total    uint64
	Used     uint64
	Display  string
	Segments []BarSegment
}

type BarSegment struct {
	Percent float64
	Color   lipgloss.Color
}

type TasksInfo struct {
	Tasks   int
	Threads int
	Running int
}
