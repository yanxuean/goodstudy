package main

import (
	"fmt"
	"time"
)

func resetOrReuseTimer(t *time.Timer, d time.Duration, sawTimeout bool) *time.Timer {
	if t == nil {
		return time.NewTimer(d)
	}
	if !t.Stop() && !sawTimeout {
		<-t.C
	}
	t.Reset(d)
	return t
}

func main() {
	var t1 *time.Timer
	var dur time.Duration = 30 * time.Second
	for _ = range []int{1, 2, 3, 4} {
		t1 = resetOrReuseTimer(t1, dur, true)

		fmt.Printf("start to wait,%#v\n", time.Now())

		select {
		case value := <-t1.C:
			fmt.Printf("receive info :%#v\n", value)

		}
		fmt.Printf("end to wait,%#v\n", time.Now())
	}
}
