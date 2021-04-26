package main

import "fmt"

// 递归
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)

}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fib(i))
	}
}
