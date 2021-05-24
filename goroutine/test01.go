package main

import (
	"fmt"
	"time"
)

// 当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine
func main() {
	go spinner(100 * time.Millisecond) // 此处的动画与fib函数并行在2个协程中执行， 但主函数返回时，所有的goroutine都会被直接打断
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
