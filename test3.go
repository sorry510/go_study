package main

// import (
//     "reflect"
// )

var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)
var c = 3
var d = c
var e = 4
var ptr *int

//变量声明方式 

func main() {
    // &地址
	println(&c, &d)
    ptr = &e
    // println(reflect.TypeOf(f))
    println(*ptr)
}