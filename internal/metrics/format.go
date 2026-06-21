package metrics

import "fmt"

func HumanBytes(v uint64) string {
	const unit = 1024
	if v < unit {
		return fmt.Sprintf("%d B", v)
	}
	div, exp := unit, 0
	for n := v / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(v)/float64(div), "KMGTPE"[exp])
}

func humanBytesCompact(v uint64) string {
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

func shorten(v string, max int) string {
	if max <= 0 || len(v) <= max {
		return v
	}
	if max <= 1 {
		return v[:max]
	}
	return v[:max-1] + "…"
}
