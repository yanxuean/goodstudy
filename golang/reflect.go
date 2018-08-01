package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Arg1 int
	Arg2 string
}

func main() {
	str1 := MyStruct{1, "hello"}
	str2 := MyStruct{2, "everyone"}

	fmt.Printf("str1=%#v\n", str1)
	fmt.Printf("str2=%#v\n", str2)

	value := reflect.ValueOf(str1).FieldByName("Arg1")
	fmt.Printf("arg1=%v\n", value.String())
	reflect.ValueOf(&str2).Elem().FieldByName("Arg1").Set(value)

	fmt.Printf("str2=%#v\n", str2)

}
