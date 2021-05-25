package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)

	v := reflect.ValueOf(3) // a reflect.Value
	fmt.Println(v)          // "3"
}
