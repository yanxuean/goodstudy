package main

import (
	"fmt"
	"os"
	"strconv"
	"os/signal"
	"time"
	//"sync"

	"golang.org/x/sys/unix"
)


func main() {
	if err := executeShim(); err != nil {
		fmt.Fprintf(os.Stderr, "containerd-shim: %s\n", err)
		os.Exit(1)
	}
}

func executeShim() error {
	args := os.Args
	delayNum := 10
	if len(args) > 1 {
		var err error
		delayNum, err = strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("invalid arg: %v",args)
		}
	}

	signals := make(chan os.Signal, 32)
	signal.Notify(signals)
	//signal.Notify(signals, unix.SIGTERM, unix.SIGINT, unix.SIGCHLD, unix.SIGPIPE)

	return handleSignals(signals, delayNum)
}


func handleSignals(signals chan os.Signal, delay int) error {
  var (
		//termOnce sync.Once
		done     = make(chan struct{})
		id       = 0
	)

	go func() {
		file := "./term-sys.log"
		for  {
			time.Sleep(1*time.Second)
			writeLog(file, "==")
		}
	}()

	for {
		file := "./term.log"
		select {
		case <-done:
			return nil
		case s := <-signals:
			id++
			switch s {
			case unix.SIGTERM, unix.SIGINT:
				//  go termOnce.Do(func() {
				go func(id int) {
					msg := fmt.Sprintf("==%dth: Receive signal: %v,delay %v s", id, s, delay)
					fmt.Println(msg)
					writeLog(file, msg)
					for i:= 0; i<delay; i++ {
						time.Sleep(1*time.Second)
						writeLog(file, "==%dth: has delayed %d s",id,i+1)
					}
					writeLog(file, "==%dth: delay over",id)
					//close(done)
				}(id)
			default:
				msg := fmt.Sprintf("==%dth: Receive unexpect signal: %v", id,s)
				fmt.Println(msg)
				writeLog(file, msg)
			}
		}
	}

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
