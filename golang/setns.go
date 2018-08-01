package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func main() {
	if syscall.Geteuid() != 0 {
		fmt.Println("abort: you want to run this as root")
		os.Exit(1)
	}

	if len(os.Args) != 2 {
		fmt.Println("abort: you must provide a PID as the sole argument")
		os.Exit(2)
	}

	namespaces := []string{"ipc", "uts", "net", "pid", "mnt"}

	for i := range namespaces {
		fd, _ := syscall.Open(filepath.Join("/proc", os.Args[1], "ns", namespaces[i]), syscall.O_RDONLY, 0644)
		err, _, msg := syscall.RawSyscall(308, uintptr(fd), 0, 0) // 308 == setns

		if err != 0 {
			fmt.Println("setns on", namespaces[i], "namespace failed:", msg)
		} else {
			fmt.Println("setns on", namespaces[i], "namespace succeeded")
		}

	}
}
