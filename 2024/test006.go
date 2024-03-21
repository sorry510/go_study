package main

import "fmt"

func main() {
	protect()
}

func protect() {
	// recover 必须用在 defer 函数中，类似于其它语言的 catch 异常
	defer func() {
		fmt.Println("done")
		// Println executes normally even if there is a panic
		if err := recover(); err != nil {
			fmt.Printf("catch 了这个错误, 错误信息是: %v", err)
		}
	}()
	fmt.Println("start")
	panic("手动抛出一个错误")
}