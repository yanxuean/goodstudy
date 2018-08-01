package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"time"
)

type Config struct {
	Age        int
	Cats       []string
	Pi         float64
	Perfection []int
	Address    string
	DOB        time.Time // requires `import time`
}

func main() {
	fmt.Println("vim-go", os.Args)
	fmt.Printf("first-arg:%s\n", os.Args[0])
	var struct1 Config = Config{Age: 30, Address: "Xian"}

	if _, err := toml.DecodeFile("test.toml", &struct1); err != nil {
		fmt.Printf("parse error,%s\n", err)
		return
	}

	fmt.Printf("the struct %#v\n", struct1)
}
