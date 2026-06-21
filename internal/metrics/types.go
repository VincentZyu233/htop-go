package metrics

import "time"

type Snapshot struct {
	CollectedAt    time.Time
	CPUPercent     float64
	PerCorePercent []float64
	Memory         MemoryStats
	Swap           SwapStats
	Disk           DiskStats
	Network        NetworkStats
	ProcessCount   int
	ThreadCount    int
	RunningCount   int
	SleepingCount  int
	TopProcesses   []ProcessInfo
	ErrorSummaries []string
}

type MemoryStats struct {
	Used        uint64
	Total       uint64
	PercentUsed float64
}

type SwapStats struct {
	Used        uint64
	Total       uint64
	PercentUsed float64
}

type DiskStats struct {
	Used        uint64
	Total       uint64
	PercentUsed float64
}

type NetworkStats struct {
	BytesSent uint64
	BytesRecv uint64
}

type ProcessInfo struct {
	PID        int32
	Name       string
	User       string
	CPUPercent float64
	MemPercent float32
	RSS        uint64
	Threads    int32
	Status     string
	Command    string
}

type SortMode string

const (
	SortByCPU  SortMode = "cpu"
	SortByMem  SortMode = "mem"
	SortByPID  SortMode = "pid"
	SortByName SortMode = "name"
)
