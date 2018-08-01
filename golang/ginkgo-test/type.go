package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

type I interface {
	m()
}

func f(y I) {
	r := y.(io.Reader)
	fmt.Println("r type", reflect.TypeOf(r))
	r.Read([]byte{})
}

type I1 struct {
}

func (i I1) m() {
}

func (i I1) Read(p []byte) (n int, err error) {
	return 0, nil
}

func main() {
	f(I1{})
	var r io.Reader

	fmt.Println("Hello, playground")
	tty, err := os.OpenFile("yan.txt", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file fail", err)
		return
	}
	r = tty // 《--r contains, schematically, the (value, type) pair, (tty, *os.File).
	fmt.Println("r type:", reflect.TypeOf(r))
	fmt.Printf("r=%#v\n", r)
	var w io.Writer
	w = r.(io.Writer) //  《--- w will contain the pair (tty, *os.File)

	var empty interface{}
	empty = w // 《--empty will again contain that same pair, (tty, *os.File).

	fmt.Println("w type:", reflect.TypeOf(w))
	fmt.Printf("w=%#v\n", w)
	fmt.Printf("empty =%#v\n", empty)
	fmt.Println("empty type:", reflect.TypeOf(empty))

}
