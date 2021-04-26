package main

import (
    "fmt"
    "math"
)
// var i, j int = 100, 60
// var i, j = 100, 60
var (
    i int = 100
    j int = 110
)

// 函数使用
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

/* 声明函数变量 */
var getSquareRoot = func(x float64) float64 {
  return math.Sqrt(x)
}

// 闭包
func getSequence() func() int {
    i := 0
    return func() int {
        i += 1
        return i  
    }
}


func main() {
	println(max(i, j))

    /* 使用函数 */
    fmt.Println(getSquareRoot(9))


    nextNumber := getSequence()

    fmt.Println(nextNumber())
    fmt.Println(nextNumber())
}
