package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
        time := time.Now().UnixNano()
        rand.Seed(time)

	fmt.Printf("Hello, playground: time:=%v, %v", time, rand.Int63())
}
