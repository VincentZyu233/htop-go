package metrics

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"

	"github.com/VincentZyu233/htop-go/internal/dashboard"
)

func CollectDashboardData() dashboard.Data {
	return DashboardDataFromSnapshot(Collect(SortByCPU, "", 0, false), false, false)
}

func DashboardDataFromSnapshot(snap Snapshot, showTask bool, showLoad bool) dashboard.Data {
	coreCount := len(snap.PerCorePercent)
	if coreCount == 0 {
		coreCount = 8
	}

	data := dashboard.Data{
		CPUs:   make([]dashboard.CPUBar, 0, coreCount),
		Memory: memoryBarFromStats("Mem", snap.Memory),
		Swap:   memoryBarFromSwap("Swp", snap.Swap),
		Uptime: formatUptime(),
	}

	if showTask {
		data.Tasks = dashboard.TasksInfo{
			Tasks:   snap.ProcessCount,
			Threads: snap.ThreadCount,
			Running: snap.RunningCount,
		}
	}

	for i := 0; i < coreCount; i++ {
		pct := 0.0
		if i < len(snap.PerCorePercent) {
			pct = snap.PerCorePercent[i]
		}
		data.CPUs = append(data.CPUs, dashboard.CPUBar{
			Label:   fmt.Sprintf("%d", i),
			Percent: pct,
		})
	}

	if showLoad {
		if avg, err := load.Avg(); err == nil {
			data.Load = [3]float64{avg.Load1, avg.Load5, avg.Load15}
		}
	}

	return data
}

func memoryBarFromStats(label string, stats MemoryStats) dashboard.MemoryBar {
	usedPct := stats.PercentUsed * 0.72
	cachePct := stats.PercentUsed * 0.18
	bufferPct := stats.PercentUsed * 0.10
	return dashboard.MemoryBar{
		Label:   label,
		Total:   stats.Total,
		Used:    stats.Used,
		Display: fmt.Sprintf("%s/%s", humanBytesCompact(stats.Used), humanBytesCompact(stats.Total)),
		Segments: []dashboard.BarSegment{
			{Percent: usedPct, Color: lipgloss.Color("#22C55E")},
			{Percent: bufferPct, Color: lipgloss.Color("#60A5FA")},
			{Percent: cachePct, Color: lipgloss.Color("#FACC15")},
		},
	}
}

func memoryBarFromSwap(label string, stats SwapStats) dashboard.MemoryBar {
	usedPct := stats.PercentUsed * 0.85
	cachePct := stats.PercentUsed * 0.15
	return dashboard.MemoryBar{
		Label:   label,
		Total:   stats.Total,
		Used:    stats.Used,
		Display: fmt.Sprintf("%s/%s", humanBytesCompact(stats.Used), humanBytesCompact(stats.Total)),
		Segments: []dashboard.BarSegment{
			{Percent: usedPct, Color: lipgloss.Color("#22C55E")},
			{Percent: cachePct, Color: lipgloss.Color("#EF4444")},
		},
	}
}

func formatUptime() string {
	seconds, err := host.Uptime()
	if err != nil {
		return "unknown"
	}

	days := seconds / 86400
	hours := (seconds % 86400) / 3600
	minutes := (seconds % 3600) / 60
	secs := seconds % 60

	if days > 0 {
		return fmt.Sprintf("%d days, %02d:%02d:%02d", days, hours, minutes, secs)
	}
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, secs)
}
