package metrics

import (
	"strings"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	gnet "github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
)

func CollectDashboardSnapshot(showTask bool) Snapshot {
	snap := Snapshot{
		CollectedAt: time.Now(),
	}

	collectCPU(&snap)
	collectMemory(&snap)
	collectDisk(&snap)
	collectNetwork(&snap)
	if showTask {
		collectProcessSummary(&snap)
	}
	return snap
}

func Collect(sortMode SortMode, filter string, limit int, showTask bool) Snapshot {
	snap := CollectDashboardSnapshot(showTask)

	processes, err := process.Processes()
	if err != nil {
		snap.ErrorSummaries = append(snap.ErrorSummaries, "proc:"+err.Error())
		return snap
	}

	filter = strings.ToLower(strings.TrimSpace(filter))
	items := make([]ProcessInfo, 0, len(processes))
	for _, proc := range processes {
		info := collectProcessInfo(proc)
		if showTask {
			updateProcessTotals(&snap, info)
		}

		if filter != "" {
			haystack := strings.ToLower(info.Name + " " + info.Command + " " + info.User)
			if !strings.Contains(haystack, filter) {
				continue
			}
		}
		items = append(items, info)
	}

	if showTask {
		snap.ProcessCount = len(processes)
	}
	sortProcesses(items, sortMode)
	if limit > 0 && len(items) > limit {
		items = items[:limit]
	}
	snap.TopProcesses = items
	return snap
}

func collectCPU(snap *Snapshot) {
	if vals, err := cpu.Percent(0, false); err == nil && len(vals) > 0 {
		snap.CPUPercent = vals[0]
	} else if err != nil {
		snap.ErrorSummaries = append(snap.ErrorSummaries, "cpu:"+err.Error())
	}

	if vals, err := cpu.Percent(0, true); err == nil {
		snap.PerCorePercent = vals
	} else {
		snap.ErrorSummaries = append(snap.ErrorSummaries, "cpu/core:"+err.Error())
	}
}

func collectMemory(snap *Snapshot) {
	if vm, err := mem.VirtualMemory(); err == nil {
		snap.Memory = MemoryStats{Used: vm.Used, Total: vm.Total, PercentUsed: vm.UsedPercent}
	} else {
		snap.ErrorSummaries = append(snap.ErrorSummaries, "mem:"+err.Error())
	}

	if sm, err := mem.SwapMemory(); err == nil {
		snap.Swap = SwapStats{Used: sm.Used, Total: sm.Total, PercentUsed: sm.UsedPercent}
	}
}

func collectDisk(snap *Snapshot) {
	if du, err := disk.Usage(rootDiskPath()); err == nil {
		snap.Disk = DiskStats{Used: du.Used, Total: du.Total, PercentUsed: du.UsedPercent}
	} else {
		snap.ErrorSummaries = append(snap.ErrorSummaries, "disk:"+err.Error())
	}
}

func collectNetwork(snap *Snapshot) {
	if counters, err := gnet.IOCounters(false); err == nil && len(counters) > 0 {
		snap.Network = NetworkStats{
			BytesSent: counters[0].BytesSent,
			BytesRecv: counters[0].BytesRecv,
		}
	} else if err != nil {
		snap.ErrorSummaries = append(snap.ErrorSummaries, "net:"+err.Error())
	}
}

func collectProcessSummary(snap *Snapshot) {
	processes, err := process.Processes()
	if err != nil {
		snap.ErrorSummaries = append(snap.ErrorSummaries, "proc:"+err.Error())
		return
	}

	snap.ProcessCount = len(processes)
	for _, proc := range processes {
		if threads, err := proc.NumThreads(); err == nil {
			snap.ThreadCount += int(threads)
		}
		if statuses, err := proc.Status(); err == nil {
			status := strings.Join(statuses, ",")
			if strings.Contains(status, "running") {
				snap.RunningCount++
			} else {
				snap.SleepingCount++
			}
		}
	}
}
