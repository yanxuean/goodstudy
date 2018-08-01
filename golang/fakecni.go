package main

import (
	"fmt"
	"os"
	"time"
)


func main() {
	
  os.Exit(1)
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
