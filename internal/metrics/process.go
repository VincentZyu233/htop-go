package metrics

import (
	"math"
	"sort"
	"strings"

	"github.com/shirou/gopsutil/v4/process"
)

func collectProcessInfo(proc *process.Process) ProcessInfo {
	info := ProcessInfo{PID: proc.Pid}

	if name, err := proc.Name(); err == nil {
		info.Name = name
	}
	if user, err := proc.Username(); err == nil {
		info.User = shorten(user, 18)
	}
	if pct, err := proc.CPUPercent(); err == nil {
		info.CPUPercent = pct
	}
	if pct, err := proc.MemoryPercent(); err == nil {
		info.MemPercent = pct
	}
	if memInfo, err := proc.MemoryInfo(); err == nil && memInfo != nil {
		info.RSS = memInfo.RSS
	}
	if threads, err := proc.NumThreads(); err == nil {
		info.Threads = threads
	}
	if statuses, err := proc.Status(); err == nil {
		info.Status = strings.Join(statuses, ",")
	}
	if cmd, err := proc.Cmdline(); err == nil {
		info.Command = shorten(cmd, 80)
	}

	return info
}

func updateProcessTotals(snap *Snapshot, info ProcessInfo) {
	if strings.Contains(info.Status, "running") {
		snap.RunningCount++
	} else {
		snap.SleepingCount++
	}
	snap.ThreadCount += int(info.Threads)
}

func sortProcesses(items []ProcessInfo, mode SortMode) {
	sort.SliceStable(items, func(i, j int) bool {
		a, b := items[i], items[j]
		switch mode {
		case SortByMem:
			if a.MemPercent == b.MemPercent {
				return a.CPUPercent > b.CPUPercent
			}
			return a.MemPercent > b.MemPercent
		case SortByPID:
			return a.PID < b.PID
		case SortByName:
			if a.Name == b.Name {
				return a.PID < b.PID
			}
			return strings.ToLower(a.Name) < strings.ToLower(b.Name)
		default:
			if math.Abs(a.CPUPercent-b.CPUPercent) < 0.01 {
				return a.PID < b.PID
			}
			return a.CPUPercent > b.CPUPercent
		}
	})
}
