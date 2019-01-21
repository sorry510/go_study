package main

var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)
var c = 3

//变量声明方式 

func main() {
	println(&c)
}