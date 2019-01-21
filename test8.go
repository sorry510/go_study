package main

import "fmt"

// 指针
func main() {
	a := 2.9
	var fp *float64    /* 指向浮点型 */
	fp = &a

	fmt.Println(&a)
	fmt.Println(fp)
	fmt.Println(*fp)
}