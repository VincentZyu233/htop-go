//go:build !windows

package metrics

func rootDiskPath() string {
	return "/"
}
