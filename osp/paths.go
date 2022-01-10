package osp

import "runtime"

var (
	winPath   = "%SystemRoot%\\System32\\drivers\\etc\\hosts"
	linuxPath = "/etc/hosts"
)

func GetPath() string {
	switch runtime.GOOS {
	case "windows":
		return winPath
	case "linux":
		return linuxPath
	}
	return winPath
}
