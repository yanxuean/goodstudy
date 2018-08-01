package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	var workdir string
	port := "8070"

	fd, err := os.Open("/proc/self/cmdline")
	if err != nil {
		fmt.Printf("Can't read /proc/self/cmdline")
		return
	}
	fd.Read([]byte(workdir))
	fd.Close()
	workdir = filepath.Dir(workdir)

	args := os.Args
	if len(args) > 1 {
		port = args[2]
	}
	fmt.Printf("Http workdir: %s, port: %s\n", workdir, port)

	http.Handle("/", http.FileServer(http.Dir(workdir)))
	e := http.ListenAndServe(":"+port, nil)
	fmt.Println(e)
}
