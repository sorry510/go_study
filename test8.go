package main

import "fmt"

// 指针
func main() {
	a := 2.9
	var ptr * float64    /* 指向浮点型 */
    var pptr ** float64 //指针的指针
	ptr = &a
    pptr = &ptr

    if(ptr != nil) { //nil 空指针
        fmt.Println(&ptr) // 指针的地址
        fmt.Println(ptr) // 地址
        fmt.Println(*ptr) // 值
        fmt.Println(*pptr)
        fmt.Println(**pptr)
    }
}