package main

import (
	"fmt"
	// "math" // 不允许出现包引入而没有被使用
)

// 全局变量是允许声明但不使用。 同一类型的多个变量可以声明在同一行，如
// 变量声明
var out = 1 // 全局变量运行定义了，没有被使用

func main() {
	var a int = 1
	var b = 2
	c := 3 //这种方式只能声明局部变量
	// var d = 4 // go中不能有没有使用的局部变量
	fmt.Println(a + b + c)
}