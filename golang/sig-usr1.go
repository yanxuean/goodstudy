package main

import (
	"fmt"
	"os"
//	"strconv"
	"os/signal"
	"time"
	"syscall"
  "runtime"

//	"golang.org/x/sys/unix"
)


func main() {
  dump := make(chan os.Signal, 32)
  signal.Notify(dump, syscall.SIGUSR1)

  go func() {
    for range dump {
      dumpStacks()
    }
  }()
  time.Sleep(time.Second * 300)	
}

func dumpStacks() {
	var (
    buf []byte
    stackSize int
  )
  bufferLen := 16384
  for stackSize == len(buf) {
		buf = make([]byte, bufferLen)
		stackSize = runtime.Stack(buf, true)
		bufferLen *= 2
	}
	buf = buf[:stackSize]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}



func writeLog(file string, format string, args ...interface{}) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	args = append([]interface{}{time.Now()}, args...)
	msg := fmt.Sprintf("%v"+format+"\n", args...)
	f.Write([]byte(msg))
}
