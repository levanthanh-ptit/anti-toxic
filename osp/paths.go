package osp

import (
	"fmt"
	"runtime"
)

var (
	winPath   = "%SystemRoot%\\System32\\drivers\\etc\\hosts"
	linuxPath = "/etc/hosts"
)

func GetPath() string {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("OS: Windows")
		return winPath
	case "linux":
		fmt.Println("OS: Linux")
		return linuxPath
	}
	return winPath
}
