package main

import "fmt"

func main() {
	x := 100
	p := &x // 地址 *p 100
	x = 1
	fmt.Println(*p)

	var foo *int
	foo = new(int)
	fmt.Println(foo, *foo) // foo = 地址 *foo = 0
}
